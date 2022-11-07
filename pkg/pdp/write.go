package pdp

import (
	"errors"
	"os"
	"strings"
)

// For log common information.
func Info(message string) {
	p.errType = errTypeInfo
	prepareMessage(&message)
	writeMessage()
}

// For log debug data.
func Debug(message string) {
	p.errType = errTypeDebug
	prepareMessage(&message)
	writeMessage()
}

// For warn user about happening in the system.
func Warning(message string) {
	p.errType = errTypeWarning
	prepareMessage(&message)
	writeMessage()
}

// For error handling and describing.
func Error(message string) {
	p.errType = errTypeError
	prepareMessage(&message)
	writeMessage()
}

// For critical warning corresponding.
func Fatal(message string) {
	p.errType = errTypeFatal
	prepareMessage(&message)
	writeMessage()
}

func writeMessage() {
	// Write to console.
	os.Stdout.Write([]byte(p.data))

	// Write to log file.
	if p.LogfilePath != empStr {
		if !isLogDirExists() {
			createLogDir()
		}
		if !isLogFileExists() {
			createLogFile()
		}
		writeToLogFile()
	}
}

func isLogDirExists() bool {
	_, err := os.Stat(getLogDirPath())
	return !errors.Is(err, os.ErrNotExist)
}

func createLogDir() {
	os.MkdirAll(getLogDirPath(), fullPerm)
}

func getLogDirPath() string {
	i := strings.LastIndex(p.LogfilePath, slash)
	return p.LogfilePath[:i]
}

func isLogFileExists() bool {
	_, err := os.Stat(p.LogfilePath)
	return !errors.Is(err, os.ErrNotExist)
}

func createLogFile() {
	file, err := os.Create(p.LogfilePath)
	if err != nil {
		Warning("Cannot create log file. Error: " + err.Error())
	}
	defer file.Close()
}

func writeToLogFile() {
	file, err := os.OpenFile(p.LogfilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, fullPerm)
	if err != nil {
		panic("Cannot open log file for writing. Error: " + err.Error())
	}
	defer file.Close()

	if _, err2 := file.WriteString(p.data); err2 != nil {
		Warning("Cannot write to log file. Error: " + err2.Error())
	}
}
