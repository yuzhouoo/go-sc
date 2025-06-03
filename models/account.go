package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/plugin/soft_delete"
)

type AccountModel struct {
	ID        int64                 `gorm:"not null;primaryKey;autoIncrement;comment:自增id;"`
	UUID      uuid.UUID             `gorm:"not null;comment:账户uuid;"`
	Name      string                `gorm:"not null;varchar:256;comment:名字;"`
	Account   string                `gorm:"not null;varchar:100;comment:账号;"`
	CreatedAt int64                 `gorm:"autoCreateTime;comment:创建时间;"`
	UpdatedAt int64                 `gorm:"autoUpdateTime;comment:更新时间;"`
	DeletedAt soft_delete.DeletedAt `gorm:"comment:删除时间;"`
}

func (AM *AccountModel) TableName() string {
	return "account"
}
