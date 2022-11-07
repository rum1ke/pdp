package pdp

import "time"

func prepareMessage(message *string) {
	getDatetime()
	restrictFractDigits()
	customizeTime()
	customizeDate()
	addOtherElems(message)
}

func getDatetime() {
	p.data = time.Now().String()
}

func restrictFractDigits() {

	if p.data[20] == cwspc {
		// add 6 zeros
		p.data = p.data[:20] + zeros6 + p.data[20:]
	} else if p.data[21] == cwspc {
		// add 5 zeros
		p.data = p.data[:21] + zeros5 + p.data[21:]
	} else if p.data[22] == cwspc {
		// add 4 zeros
		p.data = p.data[:22] + zeros4 + p.data[22:]
	} else if p.data[23] == cwspc {
		// add 3 zeros
		p.data = p.data[:23] + zeros3 + p.data[23:]
	} else if p.data[24] == cwspc {
		// add 2 zeros
		p.data = p.data[:24] + zeros2 + p.data[24:]
	} else if p.data[25] == cwspc {
		// add 1 zero
		p.data = p.data[:25] + zeros1 + p.data[25:]
	} else if p.data[26] == cwspc {
		// exit if exact 6 digits is presents
		return
	} else if p.data[27] == cwspc {
		// reduce 7 digits
		p.data = p.data[:26] + p.data[27:]
	} else if p.data[28] == cwspc {
		// reduce 8 digits
		p.data = p.data[:26] + p.data[28:]
	} else if p.data[29] == cwspc {
		// reduce 9 digits (max of time precision)
		p.data = p.data[:26] + p.data[29:]
	}
}

func customizeTime() {
	if p.TimePrec == PrecToTimezonePlace {
		// reduce time length to timezone place
		p.data = p.data[:36]
	} else if p.TimePrec == PrecToTimezoneShift {
		// reduce time length to timezone shift
		p.data = p.data[:32]
	} else if p.TimePrec == PrecToTimezoneShiftShort {
		// reduce time length to timezone shift short
		p.data = p.data[:30]
	} else if p.TimePrec == PrecToMicroseconds {
		// reduce time length to microseconds
		p.data = p.data[:26]
	} else if p.TimePrec == PrecToMilliseconds {
		// reduce time length to milliseconds
		p.data = p.data[:23]
	} else if p.TimePrec == PrecToSeconds {
		// reduce time length to seconds
		p.data = p.data[:19]
	}
}

func customizeDate() {
	if p.DateFormat == DateFormatShortYear {
		// reduce 2 first chars of date
		p.data = p.data[2:]
	} else if p.DateFormat == DateFormatNoYear {
		// reduce 11 first chars of date
		p.data = p.data[11:]
	}
}

func addOtherElems(message *string) {
	if p.DatetimePos == DatetimePosSecond {
		p.data = getErrorDesc() + tab + p.data + tab + *message + newLine
	} else if p.DatetimePos == DatetimePosLast {
		p.data = getErrorDesc() + tab + *message + tab + p.data + newLine
	} else {
		p.data = p.data + tab + getErrorDesc() + tab + *message + newLine
	}
}

func getErrorDesc() string {
	if p.errType == errTypeInfo {
		return errDescInfo
	} else if p.errType == errTypeDebug {
		return errDescDebug
	} else if p.errType == errTypeWarning {
		return errDescWarning
	} else if p.errType == errTypeError {
		return errDescError
	}
	return errDescFatal
}
