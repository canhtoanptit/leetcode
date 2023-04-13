package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	down := true
	arr := make([]string, numRows)
	row := 0
	for i := 0; i < len(s); i++ {
		sTemp := arr[row] + string(s[i])
		arr[row] = sTemp
		if row == numRows-1 {
			down = false
		} else if row == 0 {
			down = true
		}
		if down {
			row++
		} else {
			row--
		}
	}
	res := ""
	for i := 0; i < numRows; i++ {
		res += arr[i]
	}
	return res
}
