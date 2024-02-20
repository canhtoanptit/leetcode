package leetcode

import "strings"

func revertWord(word string) string {
	arrOfWord := strings.Split(word, " ")
	totalWord := len(arrOfWord)
	n := totalWord / 2

	revertedArr := make([]string, totalWord, totalWord)
	for i := 0; i <= n; i++ {
		// swtich the index
		indexOfSource := n*2 - i - 1
		revertedArr[i] = arrOfWord[indexOfSource]
		revertedArr[indexOfSource] = arrOfWord[i]
	}

	return strings.Join(revertedArr, " ")
}

func findDuplicateChar(str string) []string {
	rs := make([]string, 0, len(str))

	// for each char
	// put to the map
	// check key if exist -> duplicate
	// add duplicate char to the array
	chars := []rune(str)

	mapOfChar := make(map[string]struct{})
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if _, ok := mapOfChar[char]; ok {
			rs = append(rs, char)
		} else {
			mapOfChar[char] = struct{}{}
		}
	}

	return rs
}
