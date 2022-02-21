package pkg

func Correct(sentence []rune, pointer int, correctStr []rune) []rune {
	if pointer == len(sentence)-1 {
		return correctStr
	}
	if pointer == 0 && sentence[pointer+1] != ' ' {
		if CharacterConnectable[sentence[pointer]][IsConnectableAfter] {
			correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFisrtCharacterLast])
		}
	}

	return ""
}
