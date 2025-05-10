package gorm

import (
	"github.com/ppd324/go-ddd-layout/conf"
	"gorm.io/gorm"
)

func New(cfg *conf.MysqlConfig) (*gorm.DB, error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local)",
	//	cfg.User, cfg.Passwd, cfg.Host, cfg.Port, cfg.TableName, cfg.Charset)
	//return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return nil, nil
}
