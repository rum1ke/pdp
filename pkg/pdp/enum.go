package pdp

// Datetime Position.

const (
	DatetimePosSecond = iota
	DatetimePosLast
	DatetimePosFirst
)

// Date Format.

const (
	DateFormatFullYear = iota
	DateFormatShortYear
	DateFormatNoYear
)

// Precision.

const (
	PrecToTimezonePlace = iota
	PrecToTimezoneShift
	PrecToTimezoneShiftShort
	PrecToMicroseconds
	PrecToMilliseconds
	PrecToSeconds
)

// Defines error type.

const (
	errTypeInfo = iota
	errTypeDebug
	errTypeWarning
	errTypeError
	errTypeFatal
)
