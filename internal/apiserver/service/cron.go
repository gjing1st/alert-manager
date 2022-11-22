// Path: internal/apiserver/service
// FileName: cron.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/10/31$ 22:59$

package service

import (
	"alert-manager/internal/apiserver/config"
	"alert-manager/internal/apiserver/store"
	"alert-manager/internal/apiserver/store/mysql"
	"alert-manager/pkg/utils/crontab"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CrontabService struct {
}

var crontabStore mysql.CrontabStore

// CrontabPublish
// @description: 定时发布
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 10:53
// @success:
func CrontabPublish() {
	now := time.Now()
	cronList, err := crontabStore.SearchWaitingStatus(now)
	if err != nil {
		return
	}
	_ = store.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range cronList {

			if err != nil {
				return err
			}
			//删除待执行的定时任务
			crontabId := v.ID
			err = crontabStore.Delete(tx, crontabId)
		}

		return err
	})

}

// AddCron
// @description: 添加定时计划
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/11/1 13:52
// @success:
func AddCron() {
	seconds := config.Config.CrontabTime
	fmt.Println("定时任务设置时间seconds=", seconds)
	_, _ = crontab.GetCron().AddFunc("00 * * * * *", CrontabPublish)
}
