// Path: internal/apiserver/store/database/system
// FileName: base.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 18:20$

package initdata

import (
	"alert-manager/internal/apiserver/model/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// RegisterTables
// @description: 注册数据库表专用。每次启动程序以entity实体为准，比对数据库表进行创建新增修改字段。
// @param: db *gorm.DB
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/11 17:00
// @success:
func RegisterTables(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		//系统模块表
		entity.Category{},
		entity.Crontab{},
		entity.Policy{},
		entity.Message{},

		//自动化模块表

	)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic("注册数据库表失败")
	}

}

// InitData
// @description: 对数据库表和表数据进行初始化
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 23:45
// @success:
func InitData(db *gorm.DB) {
	// 初始化表
	RegisterTables(db)
	//初始化表数据
	err := InitTableData(db)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic("初始表数据失败")
	}
}

type InitTable interface {
	DataInserted(*gorm.DB) bool          //表数据是否已经插入
	InitializeData(*gorm.DB) (err error) //初始化表数据
}

// InitTableData
// @description: 初始化表数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/10/31 23:21
// @success:
func InitTableData(db *gorm.DB) (err error) {

	var initTables []InitTable
	initTables = append(initTables,
		//此处添加需要初始化的实现表
		InitCategory{},
	)

	for _, init := range initTables {
		if init.DataInserted(db) {
			//表数据是否已存在
			continue
		}
		//初始化表数据
		if err = init.InitializeData(db); err != nil {
			return err
		}
	}
	return
}
