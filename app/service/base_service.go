package service

import "gorm.io/gorm"

// service基类，其他Service都继承它
type baseService struct {
	db *gorm.DB
}
