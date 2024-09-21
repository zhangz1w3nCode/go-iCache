package etcd

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

// EtcdHotUpdate 服务发现
type EtcdHotUpdate struct {
	cli  *clientv3.Client // etcd连接
	lock sync.RWMutex     // 读写互斥锁
	Ctx  context.Context
}

func NewConfigHotUpdate(endpoints []string, watchPathPrefix string) (*EtcdHotUpdate, error) {
	//客户端连接etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	//viper连接etcd
	if err := viper.AddRemoteProvider("etcd3", endpoints[0], watchPathPrefix); err != nil {
		log.Fatalf("Get etcd remote config error:%v", err)
	}
	viper.SetConfigType("yaml") //设置配置文件格式
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatalf("Read etcd remote config error:%v", err)
	}

	return &EtcdHotUpdate{
		cli: cli,
		Ctx: context.Background(),
	}, nil
}

// ConfigDiscovery 读取etcd的服务并开启协程监听
func (e *EtcdHotUpdate) ConfigDiscovery(watchPathPrefix string, rootNode string) error {
	// 根据服务名称的前缀，获取所有的注册服务
	_, err := e.cli.Get(e.Ctx, watchPathPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	// 开启监听协程，监听prefix的变化
	go func() {
		watchRespChan := e.cli.Watch(e.Ctx, watchPathPrefix, clientv3.WithPrefix())
		log.Printf("watching prefix:%s now...", watchPathPrefix)
		for watchResp := range watchRespChan {
			for _, event := range watchResp.Events {
				switch event.Type {
				//如果监听到的事件是新增或者修改就重新刷新viper中的配置
				case mvccpb.PUT:
					e.hotUpdateConfig(rootNode)
				}
			}
		}
	}()

	return nil
}

// hotUpdateConfig 新增或修改本地服务
func (s *EtcdHotUpdate) hotUpdateConfig(rootNode string) {
	s.lock.Lock()
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Println("Viper Add Remote Provider ERROR", err)
	}
	value := viper.GetStringMap(rootNode)
	fmt.Println("监听到值变化, value = ", value)
	s.lock.Unlock()
}

// GetConfig 获取本地服务的value
func (s *EtcdHotUpdate) GetConfig(configPath string) string {
	s.lock.RLock()
	value := viper.GetString(configPath)
	s.lock.RUnlock()
	return value
}

// Close 关闭服务
func (e *EtcdHotUpdate) Close() error {
	return e.cli.Close()
}
