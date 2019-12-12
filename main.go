package main

import (
	"SLS/fmlog"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
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
//日志信息结构
type LogMessage struct{
	Level int         //日志等级
	MessagePATH string	//日志文件地址
	MessageNAME	string	//日志名
	mux	sync.Mutex
	cnx  context.Context
	cnxdel context.CancelFunc
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
	cnx,concel:=context.WithCancel(context.Background())
	var mu sync.Mutex
	return &LogMessage{
		Level:level,
		MessageNAME: messagename,
		MessagePATH: messagepath,
		Filed:f,
		cnx:cnx,
		cnxdel:concel,
		mux:mu,
	},nil
}
//关闭文件
func (l *LogMessage)Close(){
	l.cnxdel()
	l.Filed.Close()
}
//日志记录函数
//如果写入的日志等级低于LogMessage的等级 就将日志直接输出到终端
//通过fmlog包 格式化
func (l *LogMessage)WriterLog(str string,level int,flag int)error{
	//如果日志结构等级大于level则将本条log输出到终端
	if l.Level > level{
		println(str)
		return nil
	}
	l.mux.Lock()
	_,err:=fmt.Fprintln(l.Filed,fmlog.New(str,flag))
	l.mux.Unlock()
	if err!=nil{
		return NewLogError("写入Log失败",err)
	}
	return nil
}
//异步日志存储信息的结构
type AsynLog struct{
	l	*LogMessage
	level 	int
	flag	int
	str string
}
//异步通道
var asynChan chan *AsynLog
var initOnceChan sync.Once

//给全局变量赋值
//开始监听异步通道
func (l *LogMessage)AsynMessage(str string,level int,flag int){
	//初始化异步通道
	initOnceChan.Do(
			func(){
				asynChan = make(chan *AsynLog,20)
				go gAsyncWriteLog(asynChan,l.cnx)
			})

	//初始化log
	asl := AsynLog{
		l:     l,
		level: level,
		flag:  flag,
		str:str,
	}
	asynChan<-&asl
}
//异步通道监听函数
//利用context防止goroutinue泄露
func gAsyncWriteLog(c <-chan *AsynLog,cnx context.Context){
	var log *AsynLog
	for{
		select {
		case <-cnx.Done():
			return
		case log=<-c:
			if log.l.Level > log.level{
				println(log.str)
				return
			}

			_,err:=fmt.Fprintln(log.l.Filed,fmlog.New(log.str,log.flag))
			if err!=nil{
				fmt.Printf("日志写入失败:%s",err)
			}
		}
	}
}