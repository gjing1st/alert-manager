package global

import "errors"

const (
	TimeFormat = "2006-01-02 15:04:05"
	MaxLimit   = 1000
)

const PageSizeDefault = 10 //默认每页显示数量

const LogMsg = "alert" //日志收集关键信息标识

// 附件限制
const (
	MaxUploadFileSize = 1024 * 1024 * 50 //文件大小最大限制
	MaxFileCount      = 5                //一个文章最大附件数量
)

var (
	MaxUploadFileSizeErr = errors.New("添加附件失败，单个文件不超过50MB")
	MaxFileCountErr      = errors.New("添加附件失败，最多添加5个附件")
)

// 告警策略类型
const (
	AlertTypeCritical  = 1 //故障告警
	AlertTypeThreshold = 2 //阈值告警
)
const (
	PolicyStatusDisabled = 0 //禁用
	PolicyStatusEnable   = 1 //启用

)
