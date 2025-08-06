package kernel

import "net/http"

// IAccount 账号
type IAccount interface {
	Type() uint8 // 渠道
	Uid() int64  // fixme: 可能需要做转换
	Sid() string // session id
}

// cookie 管理
type ICookie interface {
	Load() []*http.Cookie           // load cookie
	Storage(cookies []*http.Cookie) // storage cookie
}

// IDefault 默认参数
type IDefault interface {
	Headers() map[string]string     // default headers
	QueryParams() map[string]string // default query
}

// IManager 管理器
type IManager interface {
	Account() IAccount
	Cookie() ICookie
	DefaultValue() IDefault
}
