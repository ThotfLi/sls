package fmlog

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

//格式化字符串
//接受自定义格式化字符串

const(
	DEFAULT = 0
	TIME = 1 << iota  //时间
	STACK			  //堆栈内容
	LINE			  //行
	FILE			  //文件名
)
//存放对应标识符的map
type Fmflages struct{
	flags int   // 位掩码标志位
	buf []byte  //要格式化的字节
}

//初始化 Fmflages结构
func New(str string,flags int)string{
	n := new(Fmflages)
	n.flags = flags
	n.buf = []byte(str)
	return n.format()
}
//根据标志位提供不同格式内容
func (f *Fmflages)format()string{
	pcname,line := getFileAndLine()
	pcname = filepath.Base(pcname)
	bpcname := []byte(pcname)
	bline := []byte(strconv.FormatInt(int64(line),10))
	quq := []byte(" ")
	//默认，文件名 行数 时间  消息内容
	if f.flags|DEFAULT == 0 {
		/*
		pcname  文件名
		line  行
		str_t 格式化后的时间
		*/

		t := time.Now()
		str_t := t.Format("2006/01/02 15:04")

		//join 文件名 行 时间 消息内容 用“ ”分隔
		f.buf = bytes.Join([][]byte{
							bpcname,
			 					bline,
			 						[]byte(str_t),f.buf},
			 							quq)

		return string(f.buf)

	}
	//存在stack flag 格式  文件内容 stackflag
	if f.flags&STACK != 0{
		buf := make([]byte,300)
		runtime.Stack(buf,false)
		f.buf = bytes.Join([][]byte{f.buf,buf},[]byte("\n"))
	}
	if f.flags&TIME !=0{
		t := time.Now()
		str_t :=t.Format("2006/01/02 15:04")
		f.buf = bytes.Join([][]byte{[]byte(str_t),f.buf},quq)
	}
	if f.flags&LINE != 0{
		f.buf = bytes.Join([][]byte{bline,f.buf},quq)
	}
	if f.flags&FILE !=0{
		f.buf = bytes.Join([][]byte{bpcname,f.buf},quq)
	}
	return string(f.buf)
}
//返回运行文件名和行数
func getFileAndLine()(string,int){
	pc,_,_,_:= runtime.Caller(1)
	pcName,line := runtime.FuncForPC(pc).FileLine(pc)
	return pcName,line
}