package v2

type TPid int64 // pid type(alise)

// IPixiver pixiver
type IPixiver interface {
	Pid() TPid
	SessID() string
	IsEnable() bool
}

// IPool pixiver pool
type IPool interface {
	Push(pixiver IPixiver) error
	Pop() IPixiver
	List() []IPixiver
}
