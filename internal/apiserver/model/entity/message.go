// Path: internal/apiserver/model/entity
// FileName: message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/18$ 15:30$

package entity

import "time"

type Message struct {
	BaseModel
	AlertName      string    `json:"alert_name" gorm:"column:alert_name;comment:告警策略名称;type:varchar(64);NOT NULL;"`
	AliasName      string    `json:"alias_name" gorm:"column:alias_name;comment:告警策略别名;type:varchar(64)"`
	Description    string    `json:"description" gorm:"column:description;comment:描述;type:varchar(255)"`
	Message        string    `json:"message" gorm:"column:message;comment:警告详情"`
	HostIp         string    `json:"host_ip" gorm:"column:host_ip;comment:节点ip地址;type:varchar(32);NOT NULL"`
	Node           string    `json:"node" gorm:"column:node;comment:节点名称;type:varchar(32);NOT NULL"`
	Instance       string    `json:"instance" gorm:"column:instance;comment:密码机IP地址;type:varchar(32);NOT NULL"`
	Resources      string    `json:"resources" gorm:"column:resources;comment:监控对象;NOT NULL"`
	Rules          string    `json:"rules" gorm:"column:alert_name;comment:规则;NOT NULL"`
	Summary        string    `json:"summary" gorm:"column:summary;comment:警告概括;type:varchar(255)"`
	Status         string    `json:"status" form:"status" gorm:"column:status;comment:状态;type:varchar(16)"`
	StartsAt       time.Time `json:"starts_at" gorm:"column:starts_at;comment:激活时间"`
	GroupAlertName string    `json:"group_alert_name" gorm:"column:group_alert_name;comment:告警策略名称;type:varchar(64);NOT NULL"`
	Policy         Policy    `gorm:"foreignKey:Name;references:alert_name"`
}
