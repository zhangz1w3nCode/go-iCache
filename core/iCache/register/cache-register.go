package register

import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/zhangz1w3nCode/go-iCache/config"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/etcd"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	monitorsvc "github.com/zhangz1w3nCode/go-iCache/internal/service/monitor"
	monitorpb "github.com/zhangz1w3nCode/go-iCache/pb/generate/cache-monitor"
	"google.golang.org/grpc"
	"log"
	"time"
)

func RegisterZookeeper(zookeeperServers []string, serviceName string, ip string, info map[string]grpc.ServiceInfo) error {
	zkConn, _, err := zk.Connect(zookeeperServers, time.Second*10)
	if err != nil {
		return err
	}
	defer func() {
		zkConn.Close()
	}()

	p := "/"
	if _, err := zkConn.Create(p, nil, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}

	path0 := "/services"
	if _, err := zkConn.Create(path0, nil, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}

	path := "/services/" + serviceName
	if _, err := zkConn.Create(path, nil, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}
	path2 := "/services/" + serviceName + "/" + ip
	infoStr, err := json.Marshal(info)
	data := infoStr
	if _, err := zkConn.Create(path2, data, int32(0), zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNoNode {
			return err
		}
	}

	return nil
}

func RegisterCacheServcie(s *grpc.Server, serviceName, bizAppIp string, etcdAddress string, managerCache *manager.CacheManager) *etcd.EtcdRegister {
	//读取配置文件地址设置
	config.Init("./config/config.yaml")

	monitorService := monitorsvc.NewMonitorService(managerCache)
	monitorpb.RegisterCacheMonitorServiceServer(s, monitorService)

	info := s.GetServiceInfo()
	jsonStr, err := json.Marshal(info)

	//errZk := RegisterZookeeper([]string{etcdAddress}, serviceName, bizAppIp, info)
	etcdRegister, err := etcd.NewEtcdRegister(etcdAddress)
	if err != nil {
		log.Fatalf("failed to get etcd connect: %v", err)
	}
	err = etcdRegister.ServiceRegister("/services/"+serviceName+"/"+bizAppIp, string(jsonStr), 60)
	if err != nil {
		log.Fatalf("failed to register etcd: %v", err)
	}

	return etcdRegister
}

// ServerPostHandler 关闭资源
func ServerPostHandler(serverName string, serverAddress string, zkAddress string) {
	//注销之前注册到zk的服务的grpc信息

	// 连接zookeeper
	zkConn, _, err := zk.Connect([]string{zkAddress}, time.Second*10)
	if err != nil {
		log.Fatalf("Connect Remote Server Error! %v", err)
	}

	// 检查services节点是否存在
	exists, stat, err := zkConn.Exists("/services")

	if err != nil {
		log.Fatalf("Get path resource from zookeeper stat error! %v", err)
	}
	if stat == nil {
		log.Fatalf("Get path resource from zookeeper stat error! %v", err)
	}
	if !exists {
		log.Printf("services has been deleted!")
		return
	}

	//services节点下的serverName节点是否存在
	exists, stat, err = zkConn.Exists("/services/" + serverName)

	if err != nil {
		log.Fatalf("Get path resource from zookeeper stat error! %v", err)
	}
	if stat == nil {
		log.Printf("Get path resource from zookeeper stat error! %v", err)
	}
	if !exists {
		log.Printf("services.%s has been deleted!", serverName)
		return
	}

	//services节点下的serverName节点的具体机器注册的节点是否存在
	exists, stat, err = zkConn.Exists("/services/" + serverName + "/" + serverAddress)

	if err != nil {
		log.Fatalf("Get path resource from zookeeper stat error! %v", err)
	}
	if stat == nil {
		log.Printf("Get path resource from zookeeper stat error! %v", err)
	}
	if !exists {
		log.Printf("services.%s.%s has been deleted!", serverName, serverAddress)
		return
	}

	//删除services节点下的serverName节点
	path := "/services/" + serverName + "/" + serverAddress
	if err = zkConn.Delete(path, int32(0)); err != nil {
		log.Printf("Delete services.%s.%s error!", serverName, serverAddress)
	}

	exists, stat, err = zkConn.Exists("/services/" + serverName)

	if err != nil {
		log.Printf("Get path resource from zookeeper stat error! %v", err)
	}
	if stat == nil {
		log.Printf("Get path resource from zookeeper stat error! %v", err)
	}
	if !exists {
		log.Printf("services.%s has been deleted!", serverName)
		return
	}

	children, stat, err := zkConn.Children("/services/" + serverName)
	if err != nil {
		log.Printf("Get path resource from zookeeper stat error! %v", err)
	}
	if stat == nil {
		log.Printf("Get path resource from zookeeper stat error! %v", err)
	}

	if len(children) == 0 {
		//删除services节点下的serverName节点
		path = "/services/" + serverName
		if err = zkConn.Delete(path, int32(0)); err != nil {
			log.Printf("Delete services.%s error!", serverName)
		}
	}
	log.Printf("Delete %sservices.%s.%s%s successful!", "[", serverName, serverAddress, "]")
}
