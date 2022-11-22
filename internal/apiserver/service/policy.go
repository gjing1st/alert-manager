// Path: internal/apiserver/service
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 17:25$

package service

import (
	"alert-manager/internal/apiserver/model/entity"
	"alert-manager/internal/apiserver/model/request"
	"alert-manager/internal/apiserver/store/mysql"
	"alert-manager/internal/pkg/global"
	log "github.com/sirupsen/logrus"
)

type PolicyService struct {
}

var policyStore mysql.PolicyStore

// Create
// @description: 创建策略
// @param: req 新增策略请求体
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/17 17:38
// @success:
func (ps *PolicyService) Create(req *request.PolicyCreate) (id uint, err error) {
	//TODO 去ks创建告警策略
	//密码机添加监控：(1 - avg(rate(node_cpu_seconds_total{mode="idle"}[2m])) by (instance))*100>11

	var data entity.Policy
	data.Name = req.Name
	data.AlertType = req.AlertType
	data.Resources = req.Resources
	data.Content = req.Content
	data.ContinuousNumber = req.ContinuousNumber
	data.Threshold = req.Threshold
	data.Duration = req.Duration
	data.SendDuration = req.SendDuration
	data.SendType = req.SendType
	data.Severity = req.Severity
	data.Video = req.Video
	data.SendStartTime = req.SendStartTime
	data.SendEndTime = req.SendEndTime
	data.Status = global.PolicyStatusEnable
	id, err = policyStore.Create(nil, &data)
	return
}

// List
// @description: 策略列表逻辑
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 11:23
// @success:
func (ps PolicyService) List(req *request.PolicyList) (list interface{}, total int64, err error) {
	list, total, err = policyStore.List(req)
	if err != nil {
		log.WithField("获取文章列表错误", err).Error(global.LogMsg)
	}
	return
}

// Info
// @description: 策略信息
// @param: id 策略id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:15
// @success:
func (ps PolicyService) Info(id uint) (list interface{}, err error) {
	list, err = policyStore.GetInfoById(id)
	return

}

// Edit
// @description: 策略编辑
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:41
// @success:
func (ps PolicyService) Edit(req *request.PolicyEdit) (err error) {
	//TODO 去ks创建告警策略

	var data entity.Policy
	data.Name = req.Name
	data.AlertType = req.AlertType
	data.Resources = req.Resources
	data.Content = req.Content
	data.ContinuousNumber = req.ContinuousNumber
	data.Threshold = req.Threshold
	data.Duration = req.Duration
	data.SendDuration = req.SendDuration
	data.SendType = req.SendType
	data.Severity = req.Severity
	data.Video = req.Video
	data.SendStartTime = req.SendStartTime
	data.SendEndTime = req.SendEndTime
	err = policyStore.Edit(nil, &data)
	return

}

// Delete
// @description: 批量删除策略
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 15:01
// @success:
func (ps PolicyService) Delete(ids []int) (err error) {
	//TODO 去ks删除策略
	err = policyStore.DeleteByIds(nil, ids)
	return
}
