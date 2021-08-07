package config

import (
  "fmt"
  "testing"
  "os"
)

func TestFieldInt32(t *testing.T){
  var empty int32
  empty = 0

  var non_empty int32
  non_empty = 1

  type config_format struct {
    Empty int32
    Default int32 `default:"1"`
    Int32 int32
    Set int32 
  }

  var envar string
  envar = "Int32"

  var temp_envar string
  temp_envar = os.Getenv(envar)
  os.Setenv(envar, fmt.Sprint(non_empty))

  var config *config_format
  config = &config_format{}
  if err := Init(config); err != nil {
    t.Fatal(err.Error())
  }
  config.Set = non_empty

  if err := LoadEnvar(config); err != nil {
    t.Fatalf(err.Error())
  }

  if config.Empty != empty {
    t.Errorf(empty_err)
  }

  if config.Default != non_empty {
    t.Errorf(default_err)
  }

  if config.Int32 != non_empty {
    t.Errorf(envar_err)
  }

  if config.Set != non_empty {
    t.Errorf(set_err)
  }

  os.Setenv(envar, temp_envar)
}

func TestInvalidFieldInt32(t *testing.T){
  var err error

  type config_format struct{
    Int32Invalid int32 `default:"a"`
  }
  var config *config_format
  config = &config_format{}

  err = Init(config)
  
  if err == nil {
    t.Fatalf(invalid_default_err)
  }
  if err != ErrFieldNotNumber {
    t.Fatalf(invalid_default_err)
  }

  var envar string
  envar = "Int32Invalid"

  var temp_envar string
  temp_envar = os.Getenv(envar)

  os.Setenv(envar, "a")

  err = LoadEnvar(config)

  if err == nil {
    t.Fatalf(invalid_envar_err)
  }
  if err != ErrFieldNotNumber {
    t.Fatal(invalid_envar_err)
  }

  os.Setenv(envar, temp_envar)
}
