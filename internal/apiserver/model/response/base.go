// Path: internal/apiserver/model/response
// FileName: base.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/28$ 17:24$

package response

import (
	"alert-manager/internal/pkg/functions"
	"alert-manager/internal/pkg/global"
	"alert-manager/pkg/utils/errcode"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// Response http.code=200时统一返回格式
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func OkWithMessage(message string, c *gin.Context) {
	Result(errcode.SuccessCode, map[string]interface{}{}, message, c)
}
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errcode.SuccessCode, data, message, c)
}

// Ok
// @description: 请求成功
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:03
// @success: http_code = 200
func Ok(message string, c *gin.Context) {
	Result(errcode.SuccessCode, gin.H{}, message, c)
}

// OkWithData
// @description: 请求成功并返回data数据
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:37
// @success:
func OkWithData(data interface{}, message string, c *gin.Context) {
	Result(errcode.SuccessCode, data, message, c)
}

// Failed
// @description: 返回错误
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/17 18:42
// @success: http_code = 500
func Failed(code int, message string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		code,
		gin.H{},
		message,
	})
}

// FailedNotFound
// @description: 数据没找到
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:31
// @success: http_code = 404
func FailedNotFound(code int, c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		code,
		gin.H{},
		global.DataNotFound,
	})
}

// ParamErr
// @description: 参数请求错误
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/18 14:11
// @success: http_code = 400
func ParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, global.RequestParamErr)
}

func FailWithCodeMessage(code int, message string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		code,
		gin.H{},
		message,
	})
}

// FailWithErr
// @description: 服务端错误，记录日志并返回错误信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/6 21:18
// @success:
func FailWithErr(errMsg log.Fields, code int, message string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{
		errcode.ServerPublishCode + code,
		gin.H{},
		message,
	})
	functions.AddLog(errMsg)
}
