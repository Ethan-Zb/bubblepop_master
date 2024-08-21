package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UserBaseInfo struct {
	ID              uint      `gorm:"primary_key;autoIncrement"`
	LoginType       uint      `gorm:"type:int(11);default:0"`
	UserName        string    `gorm:"column:username;type:varchar(255);default:''"`
	Gender          uint      `gorm:"type:int(11);default:0"`
	Age             uint      `gorm:"type:int(11);default:0"`
	Avatar          string    `gorm:"type:varchar(255);default:''"`
	RegisterVersion string    `gorm:"type:varchar(20);default:''"`
	RegisterIp      string    `gorm:"type:varchar(36);default:''"`
	RegisterChannel string    `gorm:"type:varchar(50);default:''"`
	NetworkName     string    `gorm:"type:varchar(50);default:''"`
	CountryCode     string    `gorm:"type:varchar(20);default:''"`
	LoginIp         string    `gorm:"type:varchar(36);default:''"`
	AndroidId       string    `gorm:"type:varchar(64);default:''"`
	Gaid            string    `gorm:"type:varchar(64);default:''"`
	Status          uint      `gorm:"type:int(11);default:0"`
	LoginTime       time.Time `gorm:"column:login_time;type:datetime;default:CURRENT_TIMESTAMP"`
	LoginDays       uint      `gorm:"type:int(11);default:0"`
	CreateAt        time.Time `gorm:"column:create_time;autoCreateTime;index" json:"create_time"`
	UpdateAt        time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (UserBaseInfo) TableName() string {
	return "user_base_info"
}

func InsertBaseInfo(db *gorm.DB, baseInfo map[string]interface{}) (uint, error) {
	if baseInfo == nil {
		return 0, nil
	}

	user := UserBaseInfo{}
	result := db.Table("user_base_info").Clauses(clause.OnConflict{
		UpdateAll: true, // 可选,如果插入冲突则更新数据
	}).Create(&baseInfo)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func GetBaseInfo(db *gorm.DB, userID uint) (map[string]interface{}, error) {

	var user UserBaseInfo

	err := db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	userInfo := map[string]interface{}{
		"user_id":   user.ID,
		"user_name": user.UserName,
	}
	return userInfo, nil
}
