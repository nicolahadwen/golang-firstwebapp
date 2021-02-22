package wiki

const (
	UNKNOWN = iota
	INVALID_URL = iota
)
type Error struct {
	Code int
	Err error
}

func (e Error) Error() string {
	panic("implement me")
}

