package model

type User struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserName string `gorm:"column:user_name" json:"user_name"`
}

func (o *User) ToPB() {

}

// TableName sets the insert table name for this struct type
func (o *User) TableName() string {
	return "ocean_user"
}
