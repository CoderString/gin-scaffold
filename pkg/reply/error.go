package reply

import "errors"

var (
	ErrSQLCondition = errors.New("must include where condition")
	ErrNotFound     = errors.New("data not found")
	ErrDirtyMsg     = errors.New("invalid message")
)

const (
	RecordNotFound = 50000 + iota
)
