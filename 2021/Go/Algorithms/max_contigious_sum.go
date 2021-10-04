package main

import "fmt"

func findMaxSubarraySum(inputArr[]int) int{
    maxSum := 0
    totalSum := 0
    for i:=0; i< len(inputArr); i++ {
        totalSum += inputArr[i]
        if totalSum > maxSum{
            maxSum = totalSum
        }
        if totalSum < 0{
            totalSum = 0
        }
    }
    return maxSum

}
func main(){
    newArr := []int {-2, -3, 4, -1, -2, 1, 5, -3}
    maxSum := findMaxSubarraySum(newArr)
    fmt.Printf("Maximum sum is %d\n", maxSum)
}