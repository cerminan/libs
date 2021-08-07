package config

import (
	"os"
	"reflect"

	"github.com/cerminan/libs/errors"
)

const errorpath = "error.json"
var (
  ErrNotPointer error
  ErrNotStruct error
  ErrIsNil error
  ErrFieldNotSupported error
  ErrFieldNotNumber error
  ErrFieldNotBoolean error
)

func init() {
  var errs errors.Errors
  errs = errors.New(errorpath)

  ErrNotPointer = errs.Code("ConfigNotPointer")
  ErrNotStruct = errs.Code("ConfigNotStruct")
  ErrIsNil = errs.Code("ConfigIsNil")
  ErrFieldNotSupported = errs.Code("ConfigFieldNotSupported")
  ErrFieldNotNumber = errs.Code("ConfigFieldNotNumber")
  ErrFieldNotBoolean = errs.Code("ConfigFieldNotBoolean")
}

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
    return ErrNotPointer
  }
  config_type = config_type.Elem()
  
  if config_type.Kind() != reflect.Struct {
    return ErrNotStruct
  }

  var config_value reflect.Value
  config_value = reflect.ValueOf(config)

  if config_value.IsZero() {
    return ErrIsNil
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

  return ErrFieldNotSupported
}
