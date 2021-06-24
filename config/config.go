package config

import (
	"os"
	"reflect"
	"strconv"

	"github.com/kuli-app/libs/errors"
)

var Errors = errors.Errors{
  "PTR" : "'config' is not a pointer.",
  "STRUCT" : "'config' is not a structure.",
  "UNSUPPORT" : "a field of 'config' has unsupported kind.",
  "NOTNUMBER" : "value is not a number.",
}

func LoadConfig(config interface{}) error {
  var defaultTag string
  defaultTag = "default"
  
  var config_type reflect.Type
  config_type = reflect.TypeOf(config)

  if config_type.Kind() != reflect.Ptr {
    return Errors.Code("PTR")
  }
  config_type = config_type.Elem()
  
  if config_type.Kind() != reflect.Struct {
    return Errors.Code("STRUCT")
  }

  var config_value reflect.Value
  config_value = reflect.ValueOf(config).Elem()

  var i int
  for i=0; i<config_value.NumField(); i++ {
    var field_value reflect.Value
    field_value = config_value.Field(i)
    var field_type reflect.StructField
    field_type = config_type.Field(i) 

    var value string
    var exists bool
    if value, exists = os.LookupEnv(field_type.Name); exists{
      setValue(field_value, value)
      continue
    }

    if !field_value.IsZero() {
      continue
    }

    if value, exists = field_type.Tag.Lookup(defaultTag); exists{
      setValue(field_value, value)
      continue
    }

    setValue(field_value, "")
  }
  return nil
}

func setValue(var_value reflect.Value, value string) error{
  var kind reflect.Kind
  kind = var_value.Type().Kind()
  switch kind {

  case reflect.String:
    var_value.SetString(value)

  case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
    var value_int64 = int64(0)
    var err error

    if value == "" {
      var_value.SetInt(value_int64)
      return nil
    }

    value_int64, err = strconv.ParseInt(value, 10, (int(kind) - 2) * 8)
    if err != nil {
      return Errors.Code("NOTNUMBER")
    }

    var_value.SetInt(value_int64)

  default:
    return Errors.Code("UNSUPPORT")

  }

  return nil
}
