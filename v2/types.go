package v2

type (
	TPid        uint64 // pid type alise
	TArtId      uint64 // artId type alise
	TIllustType uint8
)

// IPixiver pixiver
type IPixiver interface {
	Pid() TPid
	SessionID() string
	State() int

	SetPid(TPid)
	UpdateState(state int)
}
