package config

import (
  "testing"
  "os"
)

func TestFieldString(t *testing.T){
  var empty string
  empty = ""

  var non_empty string
  non_empty = "a"

  type config_format struct{
    Empty string
    Default string `default:"a"`
    String string
    Set string
  }

  var envar string
  envar = "String"

  var temp_envar string
  temp_envar = os.Getenv(envar)
  os.Setenv(envar, non_empty)

  var config *config_format
  config = &config_format{
    Set: non_empty, 
  }

  if err := LoadConfig(config); err != nil {
    t.Fatalf(err.Error())
  }

  if config.Empty != empty {
    t.Errorf(empty_err)
  }

  if config.Default != non_empty {
    t.Errorf(default_err)
  }

  if config.String != non_empty {
    t.Errorf(envar_err)
  }

  if config.Set != non_empty {
    t.Errorf(set_err)
  }

  os.Setenv(envar, temp_envar)
}

