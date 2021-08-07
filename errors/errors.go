package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type errorDictionary map[string]error
type Errors struct {
  errDict errorDictionary
  debug func(v ...interface{})
}

const uknownFormat = "Error code %s is uknown"
var errUknown = errors.New("Something wrong!")

func New(path string) (Errors) {
  var err error

  var errs Errors
  errs = Errors{
    errDict: make(errorDictionary, 0),
    debug: nil,
  }
  
  var raw []byte
  raw, err = ioutil.ReadFile(path)
  if err != nil {
    return errs
  }

  var dict map[string]string
  err = json.Unmarshal(raw, &dict)
  if err != nil {
    return errs
  }

  for code, msg := range dict {
    errs.errDict[code] = errors.New(msg)
  }

  return errs
}

func (e *Errors) SetDebug(fn func(v ...interface{})) {
  e.debug = fn
}

func (e Errors) Code(code string) error{
  var err error
  var exists bool
  if err, exists = e.errDict[code]; exists{
    return err
  }

  e.Debug(fmt.Sprintf(uknownFormat, code))

  return errUknown
}

func (e Errors) Debug(v ...interface{}) {
  if e.debug != nil {
    e.debug(v)
  }
}
