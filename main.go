package main

import (
	"SLS/fmlog"
	"fmt"
	"os"
	"path/filepath"
)
type LogError struct{
	err error
	str string
}
func (lr *LogError)Error()string{
	return lr.str
}
//生成NewlogError 结构体
func NewLogError(str string,err error)*LogError{
	lr := LogError{}
	lr.err = err
	lr.str = str
	return &lr
}

type LogMessage struct{
	Level int
	MessagePATH string
	MessageNAME	string
	Filed	*os.File
}
const(
	ERROR = iota
	WARN
	INFO
	DEBUG
)
//生成LogMessage结构体
func NewLogMessage(level int,messagepath string,messagename string)(*LogMessage,error){
	filepathstring := filepath.Join(messagepath, messagename)
	f, err := os.OpenFile(filepathstring, os.O_APPEND|os.O_CREATE, 755)
	if err != nil {
		return nil,err
	}

	return &LogMessage{
		Level:level,
		MessageNAME: messagename,
		MessagePATH: messagepath,
		Filed:       f,
	},nil
}
//关闭文件
func (l *LogMessage)Close(){
	l.Filed.Close()
}
//日志记录函数
func (l *LogMessage)WriterLog(str string,level int,flag int)error{
	//生成格式化string

	//如果日志结构等级大于level则将本条log输出到终端
	if l.Level > level{
		println(str)
		return nil
	}
	_,err:=fmt.Fprintln(l.Filed,fmlog.New(str,flag))
	if err!=nil{
		return NewLogError("写入Log失败",err)
	}
	return nil
}

