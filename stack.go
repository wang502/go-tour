package main

import "fmt"

type stack []interface{}

func (s stack) Push(i interface{}) (stack){
  return append(s, i)
}

func (s stack) Pop() (stack, interface{}){
  l := len(s)
  return s[:l-1], s[l-1]
}

func main(){
  s := make(stack, 0)
  s = s.Push(1)
  s = s.Push(2)

  s, v := s.Pop()
  fmt.Println(v)
  s, v2 := s.Pop()
  fmt.Println(v2)

  //s2 := make(stack, 0)
  var s2 stack
  s2 = s2.Push("go")
  s2 = s2.Push("lang")

  s2, last_s := s2.Pop()
  fmt.Println(last_s)
}
