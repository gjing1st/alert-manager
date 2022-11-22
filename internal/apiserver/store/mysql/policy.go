// Path: internal/apiserver/store/mysql
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 17:33$

package mysql

import (
	"alert-manager/internal/apiserver/model/entity"
	"alert-manager/internal/apiserver/model/request"
	"alert-manager/internal/apiserver/store"
	"alert-manager/internal/pkg/functions"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PolicyStore struct {
}

// Create
// @description: 新增策略
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/17 17:34
// @success:
func (ps *PolicyStore) Create(tx *gorm.DB, policy *entity.Policy) (id uint, err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Create(policy).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql创建策略失败"})
	}
	return policy.ID, err
}

// List
// @description: mysql中查询策略列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 13:36
// @success:
func (ps PolicyStore) List(req *request.PolicyList) (policies []entity.Policy, total int64, err error) {
	db := store.DB.Model(&entity.Policy{})
	if req.Keyword != "" {
		db.Where("name like ?", "%"+req.Keyword+"%").Or("1=1")
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	err = db.Count(&total).Error
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&policies).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql查询策略列表失败"})
	}
	return policies, total, err

}

// GetInfoById
// @description: 通过策略id查询策略信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:19
// @success:
func (ps PolicyStore) GetInfoById(id uint) (policy entity.Policy, err error) {
	err = store.DB.Where("id = ?", id).First(&policy).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql查询策略信息失败", "id": id})

	}
	return

}

// Edit
// @description: 策略编辑
// @param: tx 开启事务的db实例
// @param: policy 策略实体
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:43
// @success:
func (ps PolicyStore) Edit(tx *gorm.DB, policy *entity.Policy) (err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Updates(policy).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql策略编辑失败", "policy": policy})
	}
	return
}

// DeleteByIds
// @description: 通过id批量删除策略
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 15:08
// @success:
func (ps PolicyStore) DeleteByIds(tx *gorm.DB, ids []int) (err error) {
	if tx == nil {
		tx = store.DB
	}
	err = tx.Where("id in ?", ids).Delete(&entity.Policy{}).Error
	if err != nil {
		functions.AddLog(log.Fields{"err": err, "msg": "mysql策略删除失败", "ids": ids})
	}
	return

}
