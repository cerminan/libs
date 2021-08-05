package errors

import (
	"fmt"
	"testing"
  "os"
)

var e *Errors
var err error

func TestMain(m *testing.M){
  const path = "errors_test.json"

  e, _ = New(path)

  os.Exit(m.Run())
}

func TestKnownCode(t *testing.T){
  err = e.Code("err")
  if err.Error() != "msg" {
    t.Error("Unable to extract known error code")
  }
}

func TestUknownCode(t *testing.T){
  var code string
  code = "eer"
  err = e.Code(code)
  if err.Error() != fmt.Sprintf(uknownFormat, code){
    t.Error("Unable to react on uknown error code")
  }
}

func TestCompare(t *testing.T){
  var code string
  code = "err"
  if e.Code(code) != e.Code(code) {
    t.Error("Unable to detect similar error as identical")
  }
}
