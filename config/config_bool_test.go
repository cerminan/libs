package config

import (
  "fmt"
  "os"
  "testing"
)

func TestFieldBool(t *testing.T){
  var empty bool
  empty = false

  var non_empty bool
  non_empty = true

  type config_format struct {
    Empty bool
    Default bool `default:"1"`
    Bool bool
    Set bool 
  }

  var envar string
  envar = "Bool"

  var temp_envar string
  temp_envar = os.Getenv(envar)
  os.Setenv(envar, fmt.Sprint(non_empty))

  var config *config_format
  config = &config_format{Set: non_empty}

  if err := LoadConfig(config); err != nil {
    t.Fatalf(err.Error())
  }

  if config.Empty != empty {
    t.Errorf(empty_err)
  }

  if config.Default != non_empty {
    t.Errorf(default_err)
  }

  if config.Bool != non_empty {
    t.Errorf(envar_err)
  }

  if config.Set != non_empty {
    t.Errorf(set_err)
  }

  os.Setenv(envar, temp_envar)
}

func TestInvalidFieldBool(t *testing.T){
  var err error

  type config_default struct{
    Default bool `default:"a"`
  }

  err = LoadConfig(&config_default{})
  
  if err == nil {
    t.Fatalf(invalid_default_err)
  }
  if err.Error() != Errors.Code("NOTBOOL").Error() {
    t.Fatalf(invalid_default_err)
  }

  type config_envar struct {
    BoolInvalid bool
  }

  var envar string
  envar = "BoolInvalid"

  var temp_envar string
  temp_envar = os.Getenv(envar)

  os.Setenv(envar, "a")

  err = LoadConfig(&config_envar{})

  if err == nil {
    t.Fatalf(invalid_envar_err)
  }
  if err.Error() != Errors.Code("NOTBOOL").Error() {
    t.Fatal(invalid_envar_err)
  }

  os.Setenv(envar, temp_envar)
}
