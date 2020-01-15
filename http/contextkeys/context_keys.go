package contextkeys

type contextKey int

const (
	ContentType = contextKey(iota)
	AcceptHeader
	QueryValues
)
