package service

import (
	"yunfei/app/model"
)

type adminService struct {
	baseService
}

func (a *adminService) Profile() (model model.Admin) {
	a.db.First(&model)
	//global.DB.First(&model)
	return model
}
