package errors

import (
  "errors"
  "fmt"
)

type ErrorCode string
type Errors map[ErrorCode]string

var uknownFormat = "Error code %s is uknown"

func (e Errors) Code(code ErrorCode) error{
  var message string
  var exists bool
  if message, exists = e[code]; exists{
    return errors.New(message)
  }

  return errors.New(fmt.Sprintf(uknownFormat, code))
}
