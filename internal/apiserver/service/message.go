// Path: internal/apiserver/service
// FileName: message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/21$ 14:59$

package service

import (
	"alert-manager/internal/apiserver/model/entity"
	"alert-manager/internal/apiserver/model/request"
	"alert-manager/internal/apiserver/store/mysql"
)

type MessageService struct {
}

var messageStore mysql.MessageStore

// Create
// @description: 告警消息创建
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/21 15:22
// @success:
func (ms *MessageService) Create(req *request.KSMessageCreate) (err error) {
	dataArr := make([]entity.Message, len(req.Alerts))
	//组装ks的数据
	for i, v := range req.Alerts {
		dataArr[i].AlertName = v.Labels.AlertName
		dataArr[i].AliasName = v.Annotations.AliasName
		dataArr[i].HostIp = v.Labels.HostIp
		dataArr[i].Node = v.Labels.Node
		dataArr[i].Resources = v.Annotations.Resources
		dataArr[i].Description = v.Annotations.Description
		dataArr[i].Message = v.Annotations.Message
		dataArr[i].Rules = v.Annotations.Rules
		dataArr[i].Summary = v.Annotations.Summary
		dataArr[i].Status = v.Status
		dataArr[i].StartsAt = v.StartsAt
		dataArr[i].GroupAlertName = req.GroupLabels.AlertName
		dataArr[i].Instance = v.Labels.Instance
	}
	err = messageStore.CreateInBatches(nil, &dataArr, 0)
	return
}

func (ms MessageService) List(req *request.MessageList) (list interface{}, total int64, err error) {

	return
}
