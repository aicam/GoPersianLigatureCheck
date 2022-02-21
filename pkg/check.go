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

	// context define a 3 letter combination of each rune
	context := FullContext{}

	// first of all check if there is a non-persian character in the context
	if pointer < len(sentence)-1 {
		// check if the rune is Persian
		if _, ok := CharacterToUnicodeMap[sentence[pointer]]; !ok {
			correctStr = append(correctStr, sentence[pointer])
			return CorrectRune(sentence, pointer+1, correctStr)
		}

		// check if the next rune is not persian
		if _, ok := CharacterToUnicodeMap[sentence[pointer+1]]; !ok {
			context.IsPostRunePersian = false
			context.IsPostRuneConnectable = false
		} else {
			context.IsPostRunePersian = true
			if CharacterConnectable[sentence[pointer+1]][IsConnectableBefore] {
				context.IsPostRuneConnectable = true
			} else {
				context.IsPostRuneConnectable = false
			}
		}
	} else {
		context.IsPostRunePersian = false
		context.IsPostRuneConnectable = false
	}

	// check if the previous rune is not persian
	if pointer != 0 {
		if _, ok := CharacterToUnicodeMap[sentence[pointer-1]]; !ok {
			context.IsPreRunePersian = false
			context.IsPreRuneConnectable = false
		} else {
			context.IsPreRunePersian = true
			if CharacterConnectable[sentence[pointer-1]][IsConnectableAfter] {
				context.IsPreRuneConnectable = true
			} else {
				context.IsPreRuneConnectable = false
			}
		}
	} else {
		context.IsPreRunePersian = false
		context.IsPreRuneConnectable = false
	}

	if pointer == 0 {
		if context.IsPostRunePersian && sentence[pointer+1] != ' ' {
			if CharacterConnectable[sentence[pointer]][IsConnectableAfter] {
				correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstCharacterLast])
			} else {
				correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstSpaceLast])
			}
		} else {
			correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstSpaceLast])
		}
		return CorrectRune(sentence, pointer+1, correctStr)
	}

	if pointer == len(sentence)-1 {
		if context.IsPreRunePersian && context.IsPreRuneConnectable {
			correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][CharacterFirstSpaceLast])
		} else {
			correctStr = append(correctStr, sentence[pointer])
		}
		return correctStr
	}

	// Check if before and after the current alphabet are connectable alphabets
	if context.IsPreRuneConnectable && context.IsPostRuneConnectable {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][CharacterFirstCharacterLast])
	}

	// Check if before is connectable and after not
	if context.IsPreRuneConnectable && !context.IsPostRuneConnectable {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][CharacterFirstSpaceLast])
	}

	// Check if before is not connectable and after is
	if !context.IsPreRuneConnectable && context.IsPostRuneConnectable {
		correctStr = append(correctStr, CharacterToUnicodeMap[sentence[pointer]][SpaceFirstCharacterLast])
	}

	// Check if before and after are not connectable
	if !context.IsPreRuneConnectable && !context.IsPostRuneConnectable {
		correctStr = append(correctStr, sentence[pointer])
	}

	return CorrectRune(sentence, pointer+1, correctStr)
}
