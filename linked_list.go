package main

import "fmt"
import "strconv"

type Node struct{
  val int
  prev *Node
  next *Node
}

type LinkedList struct{
  head *Node
  tail *Node
  length int
}


func (list *LinkedList) Add(val int) *LinkedList{
  newNode := &Node{val, list.tail, nil}
  if list.tail != nil{
    list.tail.next = newNode
  }else{
    list.head = newNode
  }
  list.tail = newNode
  list.length += 1
  return list
}

func (list *LinkedList) AddNode(node *Node) *LinkedList{
  if list.tail != nil{
    node.prev = list.tail
    list.tail.next = node
  }else{
    list.head = node
  }
  list.tail = node
  list.length += 1
  return list
}

func (list *LinkedList) Delete(node *Node) *LinkedList{
  if list.head == list.tail && list.head == node{
    list.head = nil
    list.tail = nil
  }else if list.head == node{
    list.head = list.head.next
    list.head.prev = nil
  } else if list.tail == node{
    list.tail = list.tail.prev
    list.tail.next = nil
  }else{
    node.next.prev = node.prev
    node.prev.next = node.next
  }
  list.length -= 1
  return list
}

func (list *LinkedList) Print(){
  res := ""
  if list == nil{
    fmt.Println(res)
  }
  cur := list.head
  for cur != nil{
    res += strconv.Itoa(cur.val) + "->"
    cur = cur.next
  }
  res += "None\n"
  fmt.Println(res)
}

func (list *LinkedList) Search(val int) bool{
  cur := list.head
  for cur != nil{
    if cur.val == val{
      return true
    }
    cur = cur.next
  }
  return false
}

func (list *LinkedList) Reverse() *LinkedList{
  prev_head := list.head
  cur := list.head
  var reversed *Node
  for cur != nil{
    tmp := cur.next
    cur.next = reversed
    reversed = cur
    cur = tmp
  }
  list.head = reversed
  list.tail = prev_head
  return list
}

func (list *LinkedList) Size() int{
  return list.length
}

func main(){
  ll := &LinkedList{nil, nil, 0}
  ll.Add(1)
  ll.Add(2)
  ll.Add(3)
  ll.Print()
  ll.Delete(ll.head.next)
  ll.Print()

  ll.Reverse()
  ll.Print()

  ll.Reverse()

  ll2 := &LinkedList{nil, nil, 0}
  ll2.Add(2)
  ll2.Add(6)

}
