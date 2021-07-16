package config

import (
	"os"
	"reflect"

	"github.com/kuli-app/libs/errors"
)

var Errors = errors.New(errors.Dictionary{
  "PTR" : "'config' is not a pointer.",
  "STRUCT" : "'config' is not a structure.",
  "UNSUPPORT" : "a field of 'config' has unsupported kind.",
  "NOTNUMBER" : "value is not a number.",
  "NOTBOOL" : "value is not a valid boolean representation.",
  "NIL" : "'config' is nil.",
})

type configKind struct{
  Kind reflect.Kind
  SetValue func(reflectValue reflect.Value, value string) error
}

var configkinds = []configKind{
  configString,
  configInt,
  configInt8,
  configInt16,
  configInt32,
  configInt64,
  configBool,
}

func validateConfig(config interface{}) error {
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
  config_value = reflect.ValueOf(config)

  if config_value.IsZero() {
    return Errors.Code("NIL")
  }

  return nil
}

func Init(config interface{}) error {
  var err error

  err = validateConfig(config)
  if err != nil {
    return err
  }

  var config_type reflect.Type
  config_type = reflect.TypeOf(config).Elem()

  var config_value reflect.Value
  config_value = reflect.ValueOf(config).Elem()

  var defaultTag string
  defaultTag = "default"
  
  var i int
  for i=0; i<config_value.NumField(); i++ {
    var field_value reflect.Value
    field_value = config_value.Field(i)
    var field_type reflect.StructField
    field_type = config_type.Field(i) 

    var value string
    var exists bool

    if value, exists = field_type.Tag.Lookup(defaultTag); exists{
      if err := setValue(field_value, value); err != nil {
        return err
      }
      continue
    }
  }
  return nil
}

func LoadEnvar(config interface{}) error {
  var err error

  err = validateConfig(config)
  if err != nil {
    return err
  }

  var config_type reflect.Type
  config_type = reflect.TypeOf(config).Elem()

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
      if err := setValue(field_value, value); err != nil {
        return err
      }
      continue
    }
  }
  return nil
}

func setValue(var_value reflect.Value, value string) error{
  var err error
  var kind reflect.Kind
  kind = var_value.Type().Kind()

  for _, configkind := range configkinds {
    if kind == configkind.Kind {
      err = configkind.SetValue(var_value, value)
      if err != nil {
        return err
      }

      return nil
    }
  }

  return Errors.Code("UNSUPPORT")
}
