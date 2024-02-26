package piscine

import "fmt"

// func Index(s string, toFind string) int {

// arr := len(toFind)

// sl := len(s)
// tl := []rune(toFind)
// flag := false
// foundIn := -1

// for i, ch := range s {
// 	if tl[i] == ch {
// 		for j, char := range tl {
// 			if char == tl[j] {
// 				if flag == false {
// 					foundIn = j
// 				}
// 				flag = true

// 			}
// 		}
// 	}
// }
// return foundIn

// n := len(toFind)
// switch {
// case n == 0:
// 	return 0
// case n == 1:
// 	return IndexByte(s, toFind[0])
// case n == len(s):
// 	if toFind == s {
// 		return 0
// 	}
// 	return -1
// case n > len(s):
// 	return -1
// }

// for i, ch := range s {
// 	if arr[i] == ch {

// 	}
// }

func Index(s, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		fmt.Printf("s[i : i+len(toFind)]: %v\n", s[i:i+len(toFind)])
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
