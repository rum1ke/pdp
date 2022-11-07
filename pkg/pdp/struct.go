package pdp

var p = Preset{}

// Personal Debugging Print - a package for fast technical messages printing.
// Version: 1.27  22.08.2022
type Preset struct {
	DatetimePos int
	DateFormat  int
	TimePrec    int
	LogfilePath string

	data    string
	errType int
}

func SetDatetimePosition(position int) {
	p.DatetimePos = position
}

func SetDateFormat(dateFormat int) {
	p.DateFormat = dateFormat
}

func SetTimePrecision(timePrecision int) {
	p.TimePrec = timePrecision
}

func SetLogfilePath(logfilePath string) {
	p.LogfilePath = logfilePath
}

func SetLogPreset(preset Preset) {
	p = preset
}
