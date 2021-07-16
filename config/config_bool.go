package config

import (
  "reflect"
  "strconv"
)

var configBool = configKind {
  Kind: reflect.Bool,
  SetValue: func(reflectValue reflect.Value, value string) error {
    var err error
    var value_bool bool
    
    if value == "" {
      reflectValue.SetBool(false)
      return nil
    }

    value_bool, err = strconv.ParseBool(value)
    if err != nil {
      return Errors.Code("NOTBOOL")
    }

    reflectValue.SetBool(value_bool)

    return nil
  },
}
