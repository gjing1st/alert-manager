// Path: internal/apiserver/model/entity
// FileName: policy.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 11:38$

package entity

import "time"

// Policy 策略表
type Policy struct {
	BaseModel
	Name             string    `json:"name" form:"name" gorm:"column:name;comment:策略名称;type:varchar(64);NOT NULL;index:alert_name,unique"`
	AliasName        string    `json:"alias_name" form:"alias_name" gorm:"column:alias_name;comment:策略别名;type:varchar(128);"`
	AlertType        int       `json:"alert_type"form:"alert_type" gorm:"column:alert_type;comment:告警类型;type:tinyint(3);" `
	Resources        string    `json:"resources" form:"resources" gorm:"column:resources;comment:监控对象;type:varchar(255);"`
	Content          int       `json:"content" form:"content" gorm:"column:content;comment:监控内容;type:tinyint(3);"`
	ContinuousNumber int       `json:"continuous_number" form:"continuous_number" gorm:"column:continuous_number;comment:连续次数,单位分钟;type:tinyint(3);"`
	Threshold        int       `json:"threshold" form:"threshold" gorm:"column:threshold;comment:阈值;type:tinyint(3);"`
	Duration         int       `json:"duration" form:"duration" gorm:"column:duration;comment:间隔时间,单位分钟;type:tinyint(3);"`
	SendDuration     int       `json:"send_duration" form:"send_duration" gorm:"column:send_duration;comment:发送间隔,该时间段内只触发一次通知;type:tinyint(3);"`
	Query            string    `json:"query" form:"query" gorm:"column:query;comment:查询条件;type:varchar(255);"`
	Severity         int       `json:"severity" form:"severity" gorm:"column:severity;comment:告警级别;type:tinyint(3);"`
	Video            int       `json:"video" form:"video" gorm:"column:video;comment:声音告警;type:tinyint(3);"`
	SendType         string    `json:"send_type" form:"send_type" gorm:"column:send_type;comment:通知方式;type:varchar(255);"`
	Summary          string    `json:"summary" form:"summary" gorm:"column:summary;comment:概括;type:varchar(255);"`
	Message          string    `json:"message" form:"message" gorm:"column:message;comment:告警通知消息;"`
	Description      string    `json:"description" form:"description" gorm:"column:description;comment:描述;type:varchar(255);"`
	Rules            string    `json:"rules" form:"rules" gorm:"column:rules;comment:规则;type:varchar(255);"`
	Status           int       `json:"status" form:"status" gorm:"column:status;comment:状态;type:tinyint(3);"`
	TenantId         int       `json:"tenant_id" form:"tenant_id" gorm:"column:tenant_id;comment:租户id;type:int(11);"`
	SendAt           time.Time `json:"send_at" gorm:"column:send_at;comment:发送时间;"`
	SendStartTime    time.Time `json:"send_start_time" gorm:"column:send_start_time;comment:发送开始时间该时间段内触发告警通知;"`
	SendEndTime      time.Time `json:"send_end_time" gorm:"column:send_end_time;comment:发送时间;"`
}

func (Policy) TableName() string {
	return "policy"
}
