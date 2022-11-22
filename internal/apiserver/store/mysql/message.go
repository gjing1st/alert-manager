// Path: internal/apiserver/store/mysql
// FileName: message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/21$ 15:00$

package mysql

import (
	"alert-manager/internal/apiserver/model/entity"
	"alert-manager/internal/apiserver/model/request"
	"alert-manager/internal/apiserver/store"
	"alert-manager/internal/pkg/functions"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageStore struct {
}

// Create
// @description: 创建告警消息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/21 15:02
// @success:
func (ms MessageStore) Create(tx *gorm.DB, message *entity.Message) (id uint, err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Create(message).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql创建告警消息失败"})
	}
	return message.ID, err
}

// CreateInBatches
// @description: 批量插入告警消息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/21 15:56
// @success:
func (ms MessageStore) CreateInBatches(tx *gorm.DB, message *[]entity.Message, num uint) (err error) {
	if tx == nil {
		tx = store.DB
	}
	if num == 0 {
		err = tx.Create(message).Error

	} else {
		err = tx.CreateInBatches(message, int(num)).Error

	}
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql批量创建告警消息失败"})
	}
	return
}

// List
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/21 18:37
// @success:
func (ms MessageStore) List(req *request.MessageList) (messages []entity.Message, total int64, err error) {
	db := store.DB.Model(&entity.Message{})
	if req.Keyword != "" {
		db.Where("description like ?", "%"+req.Keyword+"%")
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	err = db.Count(&total).Error
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&messages).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
	}
	return

}
