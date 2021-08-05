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
}

const uknownFormat = "Error code %s is uknown"

func New(path string) (*Errors, error) {
  var err error
  var raw []byte
  raw, err = ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }

  var dict map[string]string
  err = json.Unmarshal(raw, &dict)

  var errDict errorDictionary
  errDict = make(errorDictionary)

  for code, msg := range dict {
    errDict[code] = errors.New(msg)
  }

  return &Errors{
    errDict: errDict,
  }, nil
}

func (e Errors) Code(code string) error{
  var err error
  var exists bool
  if err, exists = e.errDict[code]; exists{
    return err
  }

  return errors.New(fmt.Sprintf(uknownFormat, code))
}
