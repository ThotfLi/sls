package main
import(
	"testing"
)

func TestNewLogMessage(t *testing.T) {
	_,err:=NewLogMessage(1,"./","log.txt")
	if err!=nil{
		t.Fail()
	}
}