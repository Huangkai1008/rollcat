package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"rollcat/pkg/setting"
	"time"
)

var db *gorm.DB

type JsonTime struct {
	time.Time
}

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", jsonTime.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (jsonTime JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jsonTime.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jsonTime.Time, nil
}

// Scan valueOf time.Time
func (jsonTime *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jsonTime = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	// 默认model结构体
	ID        uint     `gorm:"primary_key" json:"id"`
	CreatedAt JsonTime `gorm:"type:datetime;column:create_time" json:"create_time"`
	UpdatedAt JsonTime `gorm:"type:datetime;column:update_time" json:"update_time"`
}

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = setting.Type
	dbName = setting.Name
	user = setting.User
	password = setting.Password
	host = setting.Host
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		panic("failed to connect database")
	}

	db.SingularTable(true) // 禁用复数表
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
