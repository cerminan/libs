package errors

import (
  "errors"
  "fmt"
)

type ErrorCode string
type Errors struct {
  errs map[ErrorCode]error
}

func New(list map[ErrorCode]string) Errors {
  var errs map[ErrorCode]error
  errs = make(map[ErrorCode]error)

  for code, msg := range list {
    errs[code] = errors.New(msg)
  }

  return Errors{
    errs: errs,
  }
}

var uknownFormat = "Error code %s is uknown"

func (e Errors) Code(code ErrorCode) error{
  var err error
  var exists bool
  if err, exists = e.errs[code]; exists{
    return err
  }

  return errors.New(fmt.Sprintf(uknownFormat, code))
}
