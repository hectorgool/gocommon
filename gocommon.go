/*
go install
  gocommon.a
x, err := x()
CheckError(err)
*/
package gocommon

import (
  "fmt"
  "os"
)

func CheckError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
