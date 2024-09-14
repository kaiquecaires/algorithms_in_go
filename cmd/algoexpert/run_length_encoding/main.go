package runlengthencoding

import "fmt"

func RunLengthEncoding(str string) string {
	var (
		lastChar   = rune(str[0])
		times      = 0
		encodedStr = ""
	)

	for _, c := range str {
		if times == 9 || lastChar != c {
			encodedStr += fmt.Sprintf("%d%s", times, string(lastChar))
			times = 0
		}
		times++
		lastChar = c
	}

	if times > 0 {
		encodedStr += fmt.Sprintf("%d%s", times, string(lastChar))
	}

	return encodedStr
}
