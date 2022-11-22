// Path: internal/apiserver/model/request
// FileName: message.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/18$ 15:53$

package request

import "time"

type MessageCreate struct {
}

// KSMessageCreate KubeSphere推送过来的告警消息
type KSMessageCreate struct {
	Alerts            []KSAlter `json:"alerts"`
	CommonAnnotations struct {
		Message string `json:"message"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		AlertName string `json:"alertname"`
		Alerttype string `json:"alerttype"`
		HostIp    string `json:"host_ip"`
		Node      string `json:"node"`
		Role      string `json:"role"`
		RuleId    string `json:"rule_id"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupLabels struct {
		AlertName string `json:"alertname"`
		RuleId    string `json:"rule_id"`
	}
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
}
type KSAlter struct {
	Annotations struct {
		Message        string    `json:"message"`     //详细描述
		AliasName      string    `json:"aliasName"`   //别名
		Description    string    `json:"description"` //描述
		Kind           string    `json:"kind"`
		Resources      string    `json:"resources" ` //监控对象数组
		RuleUpdateTime time.Time `json:"rule_update_time"`
		Summary        string    `json:"summary"` //规则概要
		Rules          string    `json:"rules"`
	} `json:"annotations"`
	EndsAt       time.Time `json:"endsAt"`
	Fingerprint  string    `json:"fingerprint"`
	GeneratorURL string    `json:"generatorURL"`
	Labels       struct {
		AlertName string `json:"alertname"`
		Alerttype string `json:"alerttype"`
		HostIp    string `json:"host_ip"`
		Node      string `json:"node"`
		Role      string `json:"role"`
		RuleId    string `json:"rule_id"`
		Cluster   string `json:"cluster"`
		Instance  string `json:"instance"`
	} `json:"labels"`
	StartsAt time.Time `json:"startsAt"`
	Status   string    `json:"status"`
}

type MessageList struct {
	PageInfo
	Status   int `json:"status" form:"status"`
	TenantId int `json:"tenant_id" form:"tenant_id"`
}
