// Задача 6: проверить, является ли срез палиндромом.
package main

import "fmt"

func isPalindrome(nums []int) bool {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 2, 1}
	b := []int{1, 2, 3, 4, 5}

	fmt.Println(a, "-> палиндром:", isPalindrome(a))
	fmt.Println(b, "-> палиндром:", isPalindrome(b))
}
