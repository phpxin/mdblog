package log

import(
	"fmt"
	"github.com/phpxin/mdblog/tools/logger"
	"runtime"
	"strings"
	"time"
)

func PrintPanicStackError() {
	if x := recover(); x != nil {
		PrintPanicStack()
	}
}

func PrintPanicStack() {
	stack := make([]string, 0)
	for i := 0; i < 10; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			funcName := runtime.FuncForPC(funcName).Name()
			//WriteLog(LogPanic,"frame %d:[func:%s, file: %s, line:%d]", i, funcName, file, line)
			stack = append(stack, fmt.Sprintln("frame %d:[func:%s, file: %s, line:%d]", i, funcName, file, line))
		}
	}

	Error("errstack", strings.Join(stack, "\n")+"---END---\n")
}

// Log an error msg
// prefix: if null string, it will be system
// level: the degree of the exception
// f: the format of Sprintf
// args: the args of Sprintf
func writelog(prefix,level ,f string, args ... interface{}) {
	if prefix=="" {
		prefix = "system"
	}
	msg := fmt.Sprintf(f, args ...)
	now := time.Now().Format("2006-01-02.15:04:05")
	msg = fmt.Sprintf("%s\t%s\t%s", now, level, msg )
	err := logger.WLog(prefix, msg)
	if err!=nil {
		fmt.Println("write log failed!!!")
	}
}

func Debug(p , f string, args ... interface{}){
	writelog(p,"D", f, args ...)
}

func Info(p , f string, args ... interface{}){
	writelog(p,"I", f, args ...)
}

func Warning(p , f string, args ... interface{}){
	writelog(p,"W", f, args ...)
}

func Error(p , f string, args ... interface{}){
	writelog(p,"E", f, args ...)
}

func Alert(p , f string, args ... interface{}){
	writelog(p,"A", f, args ...)
}