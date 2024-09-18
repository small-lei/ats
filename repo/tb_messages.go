package repo

import "time"

type Messages struct {
	Id         int32     `json:"id" gorm:"id"`
	ActivityId int32     `json:"activity_id" gorm:"activity_id"`
	Phone      string    `json:"phone" gorm:"phone"`
	Message    string    `json:"message" gorm:"message"`
	Status     string    `json:"status" gorm:"status"`
	SendTime   time.Time `json:"send_time" gorm:"send_time"`
}

// TableName 表名称
func (Messages) TableName() string {
	return "Messages"
}
