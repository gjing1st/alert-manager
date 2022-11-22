// Path: internal/apiserver/router/v1
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 15:46$

package v1

import (
	"alert-manager/internal/apiserver/controller"
	"github.com/gin-gonic/gin"
)

// 策略模块
func initPolicyRouter(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("policy")
	articleController := controller.PolicyController{}
	api.GET("list", articleController.List)        //列表
	api.POST("create", articleController.Create)   //创建
	api.GET("info", articleController.Info)        //详情
	api.PUT("edit", articleController.Edit)        //编辑
	api.DELETE("delete", articleController.Delete) //删除

}
