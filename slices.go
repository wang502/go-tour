package main

import "fmt"

func reverse(nums []int) {
  i := 0
  j := len(nums)-1
  for i < j{
    tmp := nums[i]
    nums[i] = nums[j]
    nums[j] = tmp
    i += 1
    j -= 1
  }
}

// accepts an array of specified length
func isPalindrome_i(nums [5]int) bool{
  i := 0
  j := len(nums)-1
  for i < j{
    if nums[i] != nums[j]{
      return false
    }
    i += 1
    j -= 1
  }
  return true
}

func isPalindrome_ii(nums []int) bool{
  i := 0
  j := len(nums)-1
  for i < j{
    if nums[i] != nums[j]{
      return false
    }
    i += 1
    j -= 1
  }
  return true
}

func print_slice(nums []int){
  for i:=0; i<len(nums); i++{
    fmt.Printf("%d ", nums[i])
  }
  fmt.Printf("\n")
}

func print_array(nums [5]int){
  for i:=0; i<5; i++{
    fmt.Printf("%d ", nums[i])
  }
  fmt.Printf("\n")
}

func testSlicePassBy_i(nums []int){
  if len(nums) == 0{
    return
  }
  nums[0] = 100
}

func testArrayPassBy_i(nums [5]int){
  nums[0] = 100
}

func main(){
  nums1 := make([]int, 10)
  for i:=0; i<10; i++{
    nums1[i] = i
  }
  reverse(nums1)
  print_slice(nums1)

  // array of fixed length
  nums2 := [5]int{1,2,3,2,1}
  fmt.Println(isPalindrome_i(nums2))

  nums3 := make([]int, 5)
  for i := 0; i< 5; i++{
    nums3[i] = i
  }
  fmt.Println(isPalindrome_ii(nums3))

  fmt.Println("Before calling testSlicePassBy_i:")
  print_slice(nums3)
  fmt.Println("After calling testSlicePassBy_i:")
  testSlicePassBy_i(nums3)
  print_slice(nums3)

  fmt.Println("Before calling testArrayPassBy_i:")
  print_array(nums2)
  testArrayPassBy_i(nums2)
  fmt.Println("After calling testArrayPassBy_i:")
  print_array(nums2)
}
