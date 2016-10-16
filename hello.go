package main

import "fmt"
//import s "strings"

func main(){
  fmt.Printf("hello world\n")

  /* declaring integers */
  var i int = 4
  fmt.Printf("integer i: %d\n", i)
  j := 5
  fmt.Printf("integer j: %d\n", j)
  var m = 6
  fmt.Printf("integer m: %d\n", m)

  /* string */
  s := "abcd"
  fmt.Printf("length of s is: %d\n", len(s))
  var reversed_s string = reverse(s)
  fmt.Printf("s is reversed: %s\n", reversed_s)
  /* extract rune value of a character */
  var r rune = rune(s[0])
  fmt.Println(r)

  /* integer array */
  arr := []int{1,2,3}
  fmt.Printf("length of arr is : %d\n", len(arr))
  arr = append(arr, 4)
  fmt.Println(arr)
  test_pass_values(arr)
  fmt.Println(arr)

  /* range */
  for i := range arr{
    fmt.Printf("%d: %d\n", i, arr[i])
  }

  for index, value := range arr{
    fmt.Printf("%d: %d\n", index, value)
  }
  /* pointer */
  var test_i int = 10
  test_alter_pointer(&test_i)
  fmt.Println(test_i)
}

func reverse(s string) string{
  r := []rune(s)
  var i, j = 0, len(s)-1
  for i < j{
    r[i], r[j] = r[j], r[i]
    i += 1
    j -= 1
  }
  return string(r)
}

func test_pass_values(arr []int) {
  var i, j = 0, len(arr)-1
  arr[i], arr[j] = arr[j], arr[i]
}

func test_alter_pointer(i *int){
  *i = 100
}
