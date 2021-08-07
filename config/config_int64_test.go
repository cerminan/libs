package config

import (
  "fmt"
  "testing"
  "os"
)

func TestFieldInt64(t *testing.T){
  var empty int64
  empty = 0

  var non_empty int64
  non_empty = 1

  type config_format struct {
    Empty int64
    Default int64 `default:"1"`
    Int64 int64
    Set int64 
  }

  var envar string
  envar = "Int64"

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

  if config.Int64 != non_empty {
    t.Errorf(envar_err)
  }

  if config.Set != non_empty {
    t.Errorf(set_err)
  }

  os.Setenv(envar, temp_envar)
}

func TestInvalidFieldInt64(t *testing.T){
  var err error

  type config_format struct{
    Int64Invalid int64 `default:"a"`
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
  envar = "Int64Invalid"

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
