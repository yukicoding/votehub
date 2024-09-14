package response

const (
	// 系统级错误码 (1000-1999)
	ErrCodeSystemError   = 1000
	ErrCodeDatabaseError = 1001
	ErrCodeCacheError    = 1002

	// 用户相关错误码 (2000-2999)
	ErrCodeInvalidInput     = 2000
	ErrCodeUserNotFound     = 2001
	ErrCodePasswordMismatch = 2002

	// 权限相关错误码 (3000-3999)
	ErrCodeUnauthorized = 3000
	ErrCodeForbidden    = 3001

	// 业务逻辑错误码 (4000-4999)
	ErrCodeResourceNotFound = 4000
	ErrCodeDuplicateEntry   = 4001

	// ... 可以继续添加更多自定义错误码 ...
)

// ErrorMessage 映射错误码到错误消息
var ErrorMessage = map[int]string{
	ErrCodeSystemError:      "系统内部错误",
	ErrCodeDatabaseError:    "数据库错误",
	ErrCodeCacheError:       "缓存错误",
	ErrCodeInvalidInput:     "无效的输入",
	ErrCodeUserNotFound:     "用户不存在",
	ErrCodePasswordMismatch: "密码不匹配",
	ErrCodeUnauthorized:     "未经授权的访问",
	ErrCodeForbidden:        "禁止访问",
	ErrCodeResourceNotFound: "资源不存在",
	ErrCodeDuplicateEntry:   "重复的条目",
}
