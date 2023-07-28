package v2

// IPixiver pixiver
type IPixiver interface {

	// Pid pixiv id
	Pid() int64

	// Status 账户状态
	Status() uint8

	// IsEnable 可用状态
	IsEnable() bool
}
