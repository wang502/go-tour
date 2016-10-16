package main

import "fmt"
import "time"

func worker(done chan bool){
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  done<- true
}

func main(){
  done := make(chan bool, 1)
  go worker(done)

  // receiver blocks until we recieve a notification from worker on channel
  // if remove <-done, the program would exit before the worker started
  <-done
}

/* ------------------------------- */

func receiver(done chan bool){
  fmt.Print("receiving...")
  time.Sleep(time.Second)
  fmt.Println("received")

  <-done
}
