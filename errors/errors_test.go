package errors

import (
	"fmt"
	"testing"
)

var e = Errors{
  "err" : "msg",
}
var err error

func TestKnownCode(t *testing.T){
  err = e.Code("err")
  if err.Error() != "msg" {
    t.Error("Unable to extract known error code")
  }
}

func TestUknownCode(t *testing.T){
  var code ErrorCode
  code = "eer"
  err = e.Code(code)
  if err.Error() != fmt.Sprintf(uknownFormat, code){
    t.Error("Unable to react on uknown error code")
  }
}
