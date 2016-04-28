package main
/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import(
  "log"
  "bufio"
  "os"
  "os/exec"
)

func main() {
  log.Println("Start")
  s := "test"
  log.Println(s)

  cmd := exec.Command("wget", "-r", "-l 10", "http://ja.stackoverflow.com/")
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  cmd.Start()
  scanner := bufio.NewScanner(stdout)
  for scanner.Scan() {
    log.Println(scanner.Text())
  }
  cmd.Wait()
}
