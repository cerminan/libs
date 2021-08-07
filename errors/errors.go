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
  var raw []byte
  raw, err = ioutil.ReadFile(path)
  if err != nil {
    return Errors{}
  }

  var dict map[string]string
  err = json.Unmarshal(raw, &dict)
  if err != nil {
    return Errors{}
  }

  var errDict errorDictionary
  errDict = make(errorDictionary)

  for code, msg := range dict {
    errDict[code] = errors.New(msg)
  }

  return Errors{
    errDict: errDict,
    debug: nil,
  } 
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

  e.debug(errors.New(fmt.Sprintf(uknownFormat, code)))

  return errUknown
}

func (e Errors) Debug(v ...interface{}) {
  if e.debug != nil {
    e.debug(v)
  }
}
