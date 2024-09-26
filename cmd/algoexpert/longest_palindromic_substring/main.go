package longestpalindromicsubstring

func LongestPalindromicSubstring(str string) string {
	longestStart, longestEnd := 0, 1

	for i := 1; i < len(str); i++ {
		oddS, oddE := checkPalindrom(str, i-1, i+1)
		evenS, evenE := checkPalindrom(str, i-1, i)
		odd := oddE - oddS
		even := evenE - evenS

		longest := longestEnd - longestStart

		if odd > even && odd > longest {
			longestStart, longestEnd = oddS, oddE
		} else if even > longest {
			longestStart, longestEnd = evenS, evenE
		}
	}

	return str[longestStart:longestEnd]
}

func checkPalindrom(str string, start int, end int) (int, int) {
	for start >= 0 && end < len(str) {
		if str[start] != str[end] {
			break
		}
		start--
		end++
	}
	return start + 1, end
}
