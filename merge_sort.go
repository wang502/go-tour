package main

import "fmt"
import "math/rand"

func MergeSort(nums []int) []int{
  if len(nums) == 0 || len(nums) == 1{
    return nums
  }
  mid := (0+len(nums))/2
  return Merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}

func Merge(nums1 []int, nums2 []int) []int{
  i,j := 0, 0
  cur := 0
  res := make([]int, len(nums1)+len(nums2))
  for i < len(nums1) && j < len(nums2){
    if nums1[i] < nums2[j]{
      res[cur] = nums1[i]
      i += 1
    }else{
      res[cur] = nums2[j]
      j += 1
    }
    cur += 1
  }
  if i < len(nums1){
    copy(res[cur:], nums1[i:])
  }
  if j < len(nums2){
    copy(res[cur:], nums2[j:])
  }
  return res
}

func print_slice(nums []int){
  for i:=0; i<len(nums); i++{
    fmt.Printf("%d ", nums[i])
  }
  fmt.Printf("\n")
}

/*
func main(){
  nums1 := make([]int, 20)
  for i:=0; i<len(nums1); i++{
    nums1[i] = rand.Intn(20)
  }
  fmt.Println("Before merge sort: ")
  print_slice(nums1)
  merged := MergeSort(nums1)
  fmt.Println("After merge sort: ")
  print_slice(merged)
}
*/
