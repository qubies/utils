package utils

/*
Generic utility functions
*/
import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

const logLevel log.Level = log.DebugLevel
const logFolder string = "logs"

func init() {
	init_logger()
}

/*
 * LOGGING FUNCTIONS
 */

func init_logger() {
	os.MkdirAll(logFolder, os.ModePerm)
	log.SetFormatter(&log.JSONFormatter{})
	log_file, err := os.OpenFile(logFolder+"/main.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("unable to create log file...")
		return
	}
	log.SetOutput(log_file)
	log.SetLevel(logLevel)
	fmt.Println("Logger Created")
}

func buildFields(fieldsIn ...string) log.Fields {
	if len(fieldsIn)%2 != 0 {
		panic("uneven number of arguments given to LogInfo")
	}
	fields := log.Fields{}
	for x := 0; x < len(fieldsIn); x += 2 {
		fields[fieldsIn[x]] = fieldsIn[x+1]
	}
	return fields
}

// LogInfo logs a JSON structured error message using fields in a key value pair relationship.
// fieldsIn requires an even number of string arguments in the form:
// 		k,v,k,v....
func LogInfo(mainMessage string, fieldsIn ...string) {
	fields := buildFields(fieldsIn...)
	log.WithFields(fields).Info(mainMessage)
}

// LogError logs a JSON structured error message using fields in a key value pair relationship.
// fieldsIn requires an even number of string arguments in the form:
// 		k,v,k,v....
func LogError(mainMessage string, err error, fieldsIn ...string) {
	fields := buildFields(fieldsIn...)
	fields["ErrorMessage"] = err
	log.WithFields(fields).Error(mainMessage)
}

/*
 *  TESTING FUNCTIONS
 */

//DidPanic is a testing function that expects code to panic and handles around it
func DidPanic(f func()) bool {
	success := true
	func() {
		defer func() {
			if r := recover(); r == nil {
				success = false
			}
		}()
		f()
	}()
	return success
}

/*
 * OUTPUT FUNCTIONS
 */

//VerbPrint is a verbose level printing system, 0 will always print, higher levels are set through the program.
func VerbPrint(outputTo io.Writer, level, verb int, iface ...interface{}) {
	if level == 0 || level <= verb {
		fmt.Fprintln(outputTo, iface...)
	}
}

//BreakError panics, halting execution and reports an error
func BreakError(service string, err error) {
	if err != nil {
		fmt.Println(service, err.Error())
		LogError(service, err)
		panic(err)
	}
}

//PrintError does not halt execution, it just reports to std.err
func PrintError(service string, err error) {
	if err != nil {
		fmt.Println(service, err.Error())
		LogError(service, err)
	}
}

/*
 * FILE FUNCTIONS
 */

// FindDirs recursively searches the path for directories and returns a slice containing them
func FindDirs(path string) []string {
	var dirs []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return dirs
}

//FileExists checks if the given file path exists.
func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
