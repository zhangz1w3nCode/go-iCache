package entity

type OceanUser struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserName string `gorm:"column:user_name" json:"user_name"`
}

func (o *OceanUser) ToPB() {

}

// TableName sets the insert table name for this struct type
func (o *OceanUser) TableName() string {
	return "ocean_user"
}
