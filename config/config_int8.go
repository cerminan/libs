package config

import (
  "reflect"
  "strconv"
)

var configInt8 = configKind{
  Kind: reflect.Int8,
  SetValue: func(reflectValue reflect.Value, value string) error {
    var err error
    var value_int64 = int64(0)

    if value == "" {
      reflectValue.SetInt(value_int64)
      return nil
    }

    value_int64, err = strconv.ParseInt(value, 10, 8)
    if err != nil {
      return ErrFieldNotNumber
    }

    reflectValue.SetInt(value_int64)

    return nil
  },
}
