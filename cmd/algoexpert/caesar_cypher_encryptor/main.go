package caesarcypherencryptor

func CaesarCipherEncryptor(str string, key int) string {
	runes := []rune{}

	for _, c := range str {
		idxOnAlphabet := (int(c-'a') + key) % 26
		idxOnUnicode := 'a' + idxOnAlphabet
		runes = append(runes, rune(idxOnUnicode))
	}

	return string(runes)
}
