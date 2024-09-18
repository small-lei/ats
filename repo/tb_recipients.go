package repo

type Recipients struct {
	Id    int32  `json:"id" gorm:"id"`
	Phone string `json:"phone" gorm:"phone"`
	Name  string `json:"name" gorm:"name"`
}

// TableName 表名称
func (Recipients) TableName() string {
	return "recipients"
}
