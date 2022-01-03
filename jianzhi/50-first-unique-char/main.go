package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
}

func firstUniqChar(s string) byte {
	b := []byte(s)
	repeatMap := make(map[byte]int)
	for _, v := range b {
		if count, ok := repeatMap[v]; ok {
			repeatMap[v] = count + 1
		} else {
			repeatMap[v] = 1
		}
	}

	for _, v := range b {
		if count, ok := repeatMap[v]; ok {
			if count == 1 {
				return v
			}
		}
	}

	return ' '
}

// 数组
// func firstUniqChar(s string) byte {
// 	cnt := [26]int{}
// 	for _, ch := range s {
// 		cnt[ch-'a']++
// 	}

// 	for i, ch := range s {
// 		if cnt[ch-'a'] == 1 {
// 			return s[i]
// 		}
// 	}

// 	return ' '
// }
