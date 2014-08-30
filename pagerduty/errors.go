package pagerduty

// ErrorCode represents potential errors returned by API calls
type ErrorCode uint

// ErrorCode constants
const (
	ErrInternalError ErrorCode = iota + 2001
	ErrInvalidInput
	ErrArgumentsCausedError
	ErrMissingArguments
	ErrInvalidSinceOrUntilParameterValues
	ErrInvalidQueryDateRange
	ErrAuthenticationFailed
	ErrAccountNotFound
	ErrAccountLocked
	ErrOnlyHTTPSAllowed
	ErrAccessDenied
	ErrRequesterIDRequired
	ErrAccountExpired
)
