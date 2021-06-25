package config

import (
  "testing"
  "os"
)

func TestNil(t *testing.T){
  type config_format struct{}
  var config_ptr *config_format

  if err := LoadConfig(config_ptr); err.Error() != Errors.Code("NIL").Error() {
    t.Fatal("Unable to detect nill config variable")
  }
}

func TestPointer(t *testing.T){
  type config_format struct{}

  if err := LoadConfig(&config_format{}); err != nil {
    t.Fatal("Unable to parse pointer config variable")
  }

  if err := LoadConfig(config_format{}); err.Error() != Errors.Code("PTR").Error() {
    t.Fatal("Unable to detect non pointer config variable")
  }
}

func TestStruct(t *testing.T){
  type config_format struct{}

  var Int int
  Int = 0

  if err := LoadConfig(&config_format{}); err != nil {
    t.Fatal("Unable to parse pointer struct type")
  }
  
  if err := LoadConfig(&Int); err.Error() != Errors.Code("STRUCT").Error() {
    t.Fatal("Unable to detect non pointer struct type")
  }
}


func TestPriority(t *testing.T){
  type config_format struct{
    PriorityEnvar string `default:"default"`
    PrioritySet string `default:"default"`
  }
  var config *config_format
  config = &config_format{PriorityEnvar: "set", PrioritySet:"set"}

  var temp_envar_PriorityEnvar string
  temp_envar_PriorityEnvar = os.Getenv("PriorityEnvar")
  os.Setenv("PriorityEnvar", "envar")

  
  if err := LoadConfig(config); err != nil {
    t.Fatalf(err.Error())
  }

  if config.PriorityEnvar != "envar" {
    t.Error("Unable to set `PriorityEnvar` field")
  }

  if config.PrioritySet != "set" {
    t.Error("Unable to set `PrioritySet` field")
  }

  os.Setenv("PriorityEnvar", temp_envar_PriorityEnvar)
}
