// Path: internal/apiserver/controller
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 15:47$

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

type PolicyController struct {
	BaseController
}

var policyService service.PolicyService

// Create
// @description: 新增策略规则
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/17 16:14
// @success:
func (pc *PolicyController) Create(c *gin.Context) {
	var req request.PolicyCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		functions.AddLog(log.Fields{"err": err})
		return
	}
	_, err := policyService.Create(&req)
	if err != nil {
		response.Failed(errcode.AlertPolicyCreateFailed, global.CreatedFailed, c)
	} else {
		response.Ok(global.CreatedSuccess, c)
	}

}

// List
// @description: 策略列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 11:18
// @success:
func (pc PolicyController) List(c *gin.Context) {
	var req request.PolicyList
	_ = c.ShouldBindQuery(&req)
	if req.PageSize == 0 {
		req.PageSize = global.PageSizeDefault
	}
	list, total, err := policyService.List(&req)
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

// Info
// @description: 获取策略详情
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:14
// @success:
func (pc PolicyController) Info(c *gin.Context) {
	var req request.PolicyInfo
	_ = c.ShouldBindQuery(&req)
	if req.Id < 1 {
		response.ParamErr(c)
		return
	}

	list, err := policyService.Info(req.Id)
	if err != nil {
		if err == errcode.ErrRecordNotFound {
			response.FailedNotFound(errcode.AlertPolicyNotFound, c)
			return
		}
		response.Failed(errcode.AlertPolicyInfoFailed, global.QueryFailed, c)

	} else {
		response.OkWithData(list, global.QuerySuccess, c)
	}

}

// Edit
// @description: 策略编辑
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:35
// @success:
func (pc PolicyController) Edit(c *gin.Context) {
	var req request.PolicyEdit
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamErr(c)
		return
	}

	err := policyService.Edit(&req)
	if err != nil {
		response.Failed(errcode.AlertPolicySaveFailed, global.SaveFailed, c)
	} else {
		response.Ok(global.SaveSuccess, c)
	}
}

// Delete
// @description: 删除策略
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:59
// @success:
func (pc PolicyController) Delete(c *gin.Context) {
	var req request.PolicyDelete
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Ids) == 0 {
		response.ParamErr(c)
		return
	}
	err := policyService.Delete(req.Ids)
	if err != nil {
		response.Failed(errcode.AlertPolicyDeleteFailed, global.DeleteFailed, c)
	} else {
		response.Ok(global.DeleteSuccess, c)
	}
}
