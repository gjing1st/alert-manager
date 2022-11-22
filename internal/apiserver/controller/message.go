// Path: internal/apiserver/controller
// FileName: message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/18$ 15:57$

package controller

import (
	"alert-manager/internal/apiserver/model/request"
	"alert-manager/internal/apiserver/model/response"
	"alert-manager/internal/apiserver/service"
	"alert-manager/internal/pkg/functions"
	"alert-manager/internal/pkg/global"
	"alert-manager/pkg/utils/errcode"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type MessageController struct {
}

var messageService service.MessageService

// Create
// @description: 创建告警消息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 17:59
// @success:
func (mc MessageController) Create(c *gin.Context) {
	var req request.KSMessageCreate
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Alerts) == 0 {
		response.ParamErr(c)
		functions.AddLog(log.Fields{"err": err})
		return
	}
	err := messageService.Create(&req)
	if err != nil {
		response.Failed(errcode.AlertMessageKSHookFailed, global.CreatedFailed, c)

	} else {
		response.Ok(global.CreatedSuccess, c)
	}

}

// List
// @description: 告警消息列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/21 18:33
// @success:
func (mc MessageController) List(c *gin.Context) {
	var req request.MessageList
	_ = c.ShouldBindQuery(&req)
	if req.PageSize == 0 {
		req.PageSize = global.PageSizeDefault
	}
	list, total, err := messageService.List(&req)
	if err != nil {
		response.Failed(errcode.AlertPolicyListFailed, global.QueryFailed, c)

	} else {
		response.OkWithData(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, global.QuerySuccess, c)
	}
}
