package exit

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func ExitSignal() chan error {
  var err chan error
  err = make(chan error, 1)

  go func() {
      c := make(chan os.Signal)
      signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
      err <- fmt.Errorf("%s", <-c)
  }()

  return err
}
