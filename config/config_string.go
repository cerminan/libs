package config

import "reflect"

var configString = configKind{
  Kind: reflect.String,
  SetValue: func(reflectValue reflect.Value, value string) error {
    reflectValue.SetString(value)

    return nil
  },
}
