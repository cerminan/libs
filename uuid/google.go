package uuid

import (
  googleUUID "github.com/google/uuid"
)

func Google() (string, error){
  var uuid googleUUID.UUID
  var err error

  uuid, err = googleUUID.NewRandom()

  if err != nil {
    return "", err
  }

  var uuid_str string
  uuid_str = uuid.String()

  return uuid_str, nil
}
