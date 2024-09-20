package service

import (
	"ats/repo"
	"errors"
	"gorm.io/gorm"
)

func CheckActUserSender(actId int32, phone string) (*repo.Messages, error) {
	row := repo.Messages{}
	if err := dbConn.Table(repo.Messages{}.TableName()).Where(" activity_id = ? and phone = ?", actId, phone).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &row, nil
}

func InsertMessage(msg repo.Messages) error {
	err := dbConn.Table(repo.Messages{}.TableName()).Create(&msg).Error
	if err != nil {
		return err
	}
	return nil
}
