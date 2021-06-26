package config

import (
  "fmt"
  "testing"
  "os"
)

func TestFieldInt16(t *testing.T){
  var empty int16
  empty = 0

  var non_empty int16
  non_empty = 1

  type config_format struct {
    Empty int16
    Default int16 `default:"1"`
    Int16 int16
    Set int16 
  }

  var envar string
  envar = "Int16"

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

  if config.Int16 != non_empty {
    t.Errorf(envar_err)
  }

  if config.Set != non_empty {
    t.Errorf(set_err)
  }

  os.Setenv(envar, temp_envar)
}

func TestInvalidFieldInt16(t *testing.T){
  var err error

  type config_format struct{
    Int16Invalid int16 `default:"a"`
  }
  var config *config_format
  config = &config_format{}

  err = Init(config)
  
  if err == nil {
    t.Fatalf(invalid_default_err)
  }
  if err.Error() != Errors.Code("NOTNUMBER").Error() {
    t.Fatalf(invalid_default_err)
  }

  var envar string
  envar = "Int16Invalid"

  var temp_envar string
  temp_envar = os.Getenv(envar)

  os.Setenv(envar, "a")

  err = LoadEnvar(config)

  if err == nil {
    t.Fatalf(invalid_envar_err)
  }
  if err.Error() != Errors.Code("NOTNUMBER").Error() {
    t.Fatal(invalid_envar_err)
  }

  os.Setenv(envar, temp_envar)
}
