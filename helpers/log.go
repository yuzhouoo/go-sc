package helpers

import (
	"fmt"
	"runtime"
)

type PrintLog struct{}

func (l *PrintLog) getFunName() string {
	pc, _, _, _ := runtime.Caller(2)
	funName := runtime.FuncForPC(pc).Name()
	return funName
}

func (l *PrintLog) Print(str string) {
	// rootPath, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("获取项目目录失败")
	// 	return
	// }
	//
	// logPath := rootPath + "/logs/logs.log"

	fmt.Println("log.Print - ", str)
	fmt.Println("log-Print", l.getFunName())

}
