package config

import (
  "reflect"
  "strconv"
)

var configInt32 = configKind{
  Kind: reflect.Int32,
  SetValue: func(reflectValue reflect.Value, value string) error {
    var err error
    var value_int64 = int64(0)

    if value == "" {
      reflectValue.SetInt(value_int64)
      return nil
    }

    value_int64, err = strconv.ParseInt(value, 10, 32)
    if err != nil {
      return Errors.Code("NOTNUMBER")
    }

    reflectValue.SetInt(value_int64)

    return nil
  },
}
