// Path: internal/pkg
// FileName: log.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/7$ 16:38$

package functions

import (
	"alert-manager/internal/pkg/global"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strconv"
)

// AddLog
// @description: 记录日志
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/7 16:40
// @success:
func AddLog(errMsg log.Fields) {
	//记录上一步调用者文件行号
	_, file, lineNo, _ := runtime.Caller(1)
	errMsg["log_file"] = file + ":" + strconv.Itoa(lineNo)
	log.WithFields(errMsg).Error(global.LogMsg)
}
