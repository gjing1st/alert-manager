// Path: pkg/utils/errcode
// FileName: alert.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/17$ 15:24$

package errcode

const (
	AlertPolicyCode  = 1 * ModuleCode //策略模块
	AlertMessageCode = 2 * ModuleCode //告警消息模块
)

// 告警模块策略相关
const (
	AlertPolicyKSCreateFailed = ServerAlertCode + AlertPolicyCode + 1 + iota //KubeSphere策略创建失败
	AlertPolicyCreateFailed                                                  //策略创建失败
	AlertPolicySaveFailed                                                    //策略保存(修改)失败
	AlertPolicyListFailed                                                    //查询策略列表失败
	AlertPolicyInfoFailed                                                    //查询策略详情失败
	AlertPolicyDeleteFailed                                                  //策略删除失败
	AlertPolicyEnableFailed                                                  //策略启用禁用失败
	AlertPolicyNotFound                                                      //策略已删除或不存在
)

const (
	AlertMessageKSHookFailed = ServerAlertCode + AlertMessageCode + 1 + iota //ks_hook告警消息创建失败
)
