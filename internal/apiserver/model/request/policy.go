// Path: internal/apiserver/model/request
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 16:14$

package request

import "time"

// PolicyCreate 策略告警
type PolicyCreate struct {
	Name             string    `json:"name"  binding:"required"`
	AlertType        int       `json:"alert_type"`
	Resources        string    `json:"resources"`                    //监控对象
	Content          int       `json:"content"`                      //监控内容 CPU，内存，硬盘使用率
	ContinuousNumber int       `json:"continuous_number"`            //连续次数
	Threshold        int       `json:"threshold"`                    //阈值
	Duration         int       `json:"duration" `                    //时间段范围。告警规则中设置的情形持续时间达到次数后，告警触发
	SendDuration     int       `json:"send_duration"`                //发送间隔
	SendType         string    `json:"send_type" binding:"required"` //通知方式
	Severity         int       `json:"severity" `                    //告警等级
	Video            int       `json:"video"`                        //声音告警
	SendStartTime    time.Time `json:"send_start_time"`              //该时间段内触发告警通知
	SendEndTime      time.Time `json:"send_end_time"`
}

// PolicyList 策略列表
type PolicyList struct {
	PageInfo
	Status   int `json:"status" form:"status"`
	TenantId int `json:"tenant_id" form:"tenant_id"`
}

// PolicyInfo 策略详情
type PolicyInfo struct {
	Id uint `json:"id" form:"id"`
}

// PolicyEdit 策略编辑
type PolicyEdit struct {
	PolicyCreate
	Id uint `json:"id" form:"id"`
}

// PolicyDelete 批量删除文章
type PolicyDelete struct {
	IdsReq
}
