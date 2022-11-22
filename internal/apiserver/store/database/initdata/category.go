// Path: internal/apiserver/store/database/system
// FileName: category.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 18:15$

package initdata

import (
	"alert-manager/internal/apiserver/model/entity"
	"alert-manager/internal/pkg/global"
	"errors"
	"gorm.io/gorm"
)

type InitCategory struct {
}

// DataInserted
// @description: 数据是否已插入
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 18:19
// @success:
func (i InitCategory) DataInserted(db *gorm.DB) bool {
	if errors.Is(db.Where("name = ?", "商用密码标准规范").First(&entity.Category{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// InitializeData
// @description: 初始化数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 17:26
// @success:
func (i InitCategory) InitializeData(db *gorm.DB) (err error) {
	if db == nil {
		return global.DBNullErr
	}
	entities := []entity.Category{
		{Name: "系统通知", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategoryNotification}}, //ID:1
		{Name: "帮助中心", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategoryHelp}},
		{Name: "安全资讯", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategorySafeInformation}},
		{Name: "法律法规", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategoryLawsRegulations}},
		{Name: "商用密码标准规范", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategoryCipherStandard}},
		{Name: "等保政策文件", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategoryGradePolicy}},
		{Name: "密评政策文件", Status: global.StatusEnable, BaseModel: entity.BaseModel{ID: global.CategoryEstimatePolicy}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return global.InitDataErr
	}
	return
}
