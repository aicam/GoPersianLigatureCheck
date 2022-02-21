package pkg

const SpaceFirstSpaceLast = 0
const SpaceFisrtCharacterLast = 1
const CharacterFirstCharacterLast = 2
const CharacterFirstSpaceLast = 3

var CharacterToUnicodeMap = map[rune][]rune{
	'ا': {'ا', 'ا', 'ﺎ', 'ﺎ'},
	'آ': {'آ', 'آ', 'آ', 'آ'},
	'ب': {'ب', 'ﺑ', 'ﺒ', 'ﺐ'},
	'پ': {'پ', 'ﭘ', 'ﭙ', 'ﭗ'},
	'ت': {'ت', '𞸵', 'ﺘ', 'ﺖ'},

}
