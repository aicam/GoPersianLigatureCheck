package pkg

func Correct(sentence string) string {
	return reverse(string(CorrectRune([]rune(sentence), 0, nil)))
}
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func CorrectRune(sentence []rune, pointer int, correctStr []rune) []rune {
	if pointer == len(sentence)-1 {
		if CharacterConnectable[sentence[pointer-1]][IsConnectableAfter] {
			correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][CharacterFirstSpaceLast])
		} else {
			correctStr = append(correctStr, sentence[pointer])
		}
		return correctStr
	}

	if pointer == 0 && sentence[pointer+1] != ' ' {
		if CharacterConnectable[sentence[pointer]][IsConnectableAfter] {
			correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstCharacterLast])
		}
		return CorrectRune(sentence, pointer+1, correctStr)
	}

	// handle starting space
	if sentence[pointer-1] == ' ' {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstCharacterLast])
		return CorrectRune(sentence, pointer+1, correctStr)
	}

	// Check if before and after the current alphabet are connectable alphabets
	if CharacterConnectable[sentence[pointer-1]][IsConnectableAfter] &&
		CharacterConnectable[sentence[pointer+1]][IsConnectableBefore] {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][CharacterFirstCharacterLast])
	}

	// Check if before is connectable and after not
	if CharacterConnectable[sentence[pointer-1]][IsConnectableAfter] &&
		!CharacterConnectable[sentence[pointer+1]][IsConnectableBefore] {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][CharacterFirstSpaceLast])
	}

	// Check if before is not connectable and after is
	if !CharacterConnectable[sentence[pointer-1]][IsConnectableAfter] &&
		CharacterConnectable[sentence[pointer+1]][IsConnectableBefore] {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstCharacterLast])
	}

	// Check if before and after are not connectable
	if sentence[pointer-1] != ' ' &&
		!CharacterConnectable[sentence[pointer-1]][IsConnectableAfter] &&
		!CharacterConnectable[sentence[pointer+1]][IsConnectableBefore] {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstSpaceLast])
	}

	return CorrectRune(sentence, pointer+1, correctStr)
}
