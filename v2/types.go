package v2

type TPid uint64 // pid type(alise)

type TArtId uint64

type TIllustType uint8

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
