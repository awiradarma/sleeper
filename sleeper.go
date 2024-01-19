package main

import (
 "fmt"
 "time"
)

func main() {
   for {
      fmt.Println("Still here..")
      time.Sleep(1 * time.Minute)
   }
}
