package repo

type Activities struct {
	Id            int32  `json:"id" gorm:"id"`
	CsvPath       int32  `json:"csv_path" gorm:"csv_path"`
	Template      string `json:"template" gorm:"template"`
	ScheduledTime string `json:"scheduled_time" gorm:"scheduled_time"`
}

// TableName 表名称
func (Activities) TableName() string {
	return "activities"
}
