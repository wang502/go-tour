package main

import "golang.org/x/tour/tree"
import "fmt"

type stack []interface{}

func (s stack) Push(v interface{}) stack{
  return append(s, v)
}

func (s stack) Pop() (stack, interface{}){
  l := len(s)
  return s[:l-1], s[l-1]
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
  s := make(stack, 0)
  cur := t
  // declare a variable of interface{} type
  var last interface{}
  for true{
    if cur != nil{
      s = s.Push(cur)
      cur = cur.Left
    }else{
      if len(s) == 0{
        break
      }else{
        s, last = s.Pop()

        // cast interface to *tree.Tree
        last_t := last.(*tree.Tree)
        ch<-last_t.Value
        cur = last_t.Right
      }
    }
  }
  // if remove close(), a deadlock will occur
  close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
  ch1 := make(chan int)
  ch2 := make(chan int)

  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for v := range ch1{
    if v != <-ch2 {
      return false
    }
  }
  return true
}

func main(){
  tree1 := tree.New(1)
  tree2 := tree.New(1)
  fmt.Println(Same(tree1, tree2))

  tree3 := tree.New(1)
  tree4 := tree.New(2)
  fmt.Println(Same(tree3, tree4))

  tree5 := tree.New(3)
  tree6 := tree.New(3)
  fmt.Println(Same(tree5, tree6))

}
