package errors

var (
	ErrWrongQuery           = NewHTTPError(403, "Wrong query")
	ErrWrongPaginationQuery = NewHTTPError(402, "Wrong pagination query")
	ErrInvlaidSessionUser   = NewHTTPError(401, "Invalid session user")
)
