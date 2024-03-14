package main

import "fmt"

func longPrefix(strs []string) string {
	prefix := ""

	if len(strs) == 0 {
		return prefix
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}

	return prefix
}

func main() {
	strs1 := []string{"floower", "floow", "flight"}
	fmt.Println("Input:", strs1)
	fmt.Println("Output:", longPrefix(strs1))

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println("Input:", strs2)
	fmt.Println("Output:", longPrefix(strs2))
}
