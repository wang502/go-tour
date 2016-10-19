package main

// internet cafe problem from
// http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/

import (
  "fmt"
  "time"
  "math/rand"
)
// blocked goroutinew will be unblocked in FIFO order

func tourist(id int, ch chan int, done_ch chan int){
  select{
  case ch<-id:
    //turn = true
  default:
    fmt.Printf("Tourist %d is waiting for the turn\n", id)
    ch<-id
  }
  fmt.Printf("Tourist %d is online\n", id)
  time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
  <-ch
  fmt.Printf("Tourist %d is done\n", id)
  done_ch<-id
}

func tourist_ii(id int, ch chan int, done_ch chan int){
  //fmt.Printf("Tourist %d is waiting for the turn\n", id)
  //turn := false
  ch<-id
  fmt.Printf("Tourist %d is online\n", id)
  time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
  <-ch
  fmt.Printf("Tourist %d is done\n", id)
  done_ch<-id
}

func main(){
  numPlayers := 15
  ch := make(chan int, 8)
  done_ch := make(chan int, numPlayers)
  for i:= 1; i<=numPlayers; i++{
    go tourist_ii(i, ch, done_ch)
  }
  for i:= 1; i<= numPlayers; i++{
    <-done_ch
  }
  fmt.Println("The place is empty, let's close up and go to the beach!")
}
