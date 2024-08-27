package service

import (
	"sync"
	"yunfei/internal/global"
)

// var (
//
//	BaseService = &baseService{db: global.DB}
//
//	AdminService = &adminService{baseService: *BaseService}
//	//AdminService = &adminService{baseService: BaseService}
//
// )
var (
	once         sync.Once
	BaseService  *baseService
	AdminService *adminService
)

func InitServices() {
	once.Do(func() {
		BaseService = &baseService{db: global.DB}
		AdminService = &adminService{baseService: *BaseService}
	})
}
