package config

import (
  "fmt"
  "testing"
  "os"
)

func TestFieldInt(t *testing.T){
  var empty int
  empty = 0

  var non_empty int
  non_empty = 1

  type config_format struct {
    Empty int
    Default int `default:"1"`
    IntValid int
    Set int 
  }

  var envar string
  envar = "IntValid"
  
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

  if config.IntValid != non_empty {
    t.Errorf(envar_err)
  }

  if config.Set != non_empty {
    t.Errorf(set_err)
  }

  os.Setenv(envar, temp_envar)
}

func TestInvalidFieldInt(t *testing.T){
  var err error

  type config_default struct{
    Default int `default:"a"`
  }

  err = LoadConfig(&config_default{})
  
  if err == nil {
    t.Fatalf(invalid_default_err)
  }
  if err.Error() != Errors.Code("NOTNUMBER").Error() {
    t.Fatalf(invalid_default_err)
  }

  type config_envar struct {
    IntInvalid int
  }

  var envar string
  envar = "IntInvalid"

  var temp_envar string
  temp_envar = os.Getenv(envar)

  os.Setenv(envar, "a")

  err = LoadConfig(&config_envar{})

  if err == nil {
    t.Fatalf(invalid_envar_err)
  }
  if err.Error() != Errors.Code("NOTNUMBER").Error() {
    t.Fatal(invalid_envar_err)
  }

  os.Setenv(envar, temp_envar)
}
