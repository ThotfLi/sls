package main

import (
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
	Flag int
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
func NewLogMessage(flag int,messagepath string,messagename string)(*LogMessage,error){
	filepathstring := filepath.Join(messagepath, messagename)
	f, err := os.OpenFile(filepathstring, os.O_SYNC|os.O_CREATE, 755)
	if err != nil {
		return nil,err
	}

	return &LogMessage{
		Flag:flag,
		MessageNAME: messagename,
		MessagePATH: messagepath,
		Filed:       f,
	},nil
}
//关闭文件
func (l *LogMessage)Close(){
	l.Filed.Close()
}
//debug日志记录函数
func (l *LogMessage)Error(str string)error{
	_,err:=fmt.Fprintln(l.Filed,str)
	if err!=nil{
		return NewLogError("写入Log失败",err)
	}
	return err
}

