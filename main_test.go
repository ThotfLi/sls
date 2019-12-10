package main
import(
	"SLS/fmlog"
	"testing"
)

func TestNewLogMessage(t *testing.T) {
	log,err:=NewLogMessage(2,"./","log.txt")
	if err!=nil{
		t.Fail()
	}
	log.WriterLog("这是第五条记录",2,fmlog.FILE|fmlog.LINE|fmlog.STACK)
	log.WriterLog("这是第六条记录",2,fmlog.DEFAULT)

	log.Close()
}


