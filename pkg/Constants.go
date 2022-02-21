package pkg

const SpaceFirstSpaceLast = 0
const SpaceFirstCharacterLast = 1
const CharacterFirstCharacterLast = 2
const CharacterFirstSpaceLast = 3

var CharacterToUnicodeMap = map[rune][]rune{
	'ا': {'ا', 'ا', 'ﺎ', 'ﺎ'},
	'آ': {'آ', 'آ', 'آ', 'آ'},
	'ب': {'ب', 'ﺑ', 'ﺒ', 'ﺐ'},
	'پ': {'پ', 'ﭘ', 'ﭙ', 'ﭗ'},
	'ت': {'ت', '𞸵', 'ﺘ', 'ﺖ'},
	'ث': {'ث', 'ﺛ', 'ﺜ', 'ﺚ'},
	'ج': {'ج', 'ﺟ', 'ﺠ', 'ﺞ'},
	'چ': {'چ', 'ﭼ', 'ﭽ', 'ﭻ'},
	'ح': {'ح', 'ﺣ', 'ﺤ', 'ﺢ'},
	'خ': {'خ', 'ﺧ', 'ﺨ', 'ﺦ'},
	'د': {'د', 'د', 'د', 'د'},
	'ذ': {'ذ', 'ذ', 'ذ', 'ذ'},
	'ر': {'ر', 'ر', 'ر', 'ر'},
	'ز': {'ز', 'ز', 'ز', 'ز'},
	'ژ': {'ژ', 'ژ', 'ژ', 'ژ'},
	'س': {'س', '𞸮', 'ﺴ', 'ﺲ'},
	'ش': {'ش', 'ﺷ', 'ﺸ', 'ﺶ'},
	'ص': {'ص', 'ﺻ', 'ﺼ', 'ﺺ'},
	'ض': {'ض', 'ﺿ', 'ﻀ', 'ﺾ'},
	'ط': {'ط', 'ط', 'ﻂ', 'ﻂ'},
	'ظ': {'ظ', 'ظ', 'ﻆ', 'ﻆ'},
	'ع': {'ع', 'ﻋ', 'ﻌ', 'ﻊ'},
	'غ': {'غ', '𞸻', 'ﻐ', 'ﻎ'},
	'ف': {'ف', 'ﻓ', 'ﻔ', 'ﻒ'},
	'ق': {'ق', 'ﻗ', 'ﻘ', 'ﻖ'},
	'ک': {'ک', 'ﮐ', 'ﮑ', 'ﮏ'},
	'گ': {'گ', 'ﮔ', 'ﮕ', 'ﮓ'},
	'ل': {'ل', 'ﻟ', 'ﻠ', 'ﻞ'},
	'م': {'م', 'ﻣ', 'ﻤ', 'ﻢ'},
	'ن': {'ن', 'ﻧ', 'ﻨ', 'ﻦ'},
	'و': {'و', 'و', 'و', 'و'},
	'ه': {'ه', 'ﻫ', 'ﻬ', 'ﻪ'},
	'ی': {'ی', 'ﯾ', 'ﯿ', 'ﯽ'},
	' ': {' ', ' ', ' ', ' '},
}

const IsConnectableBefore = 0
const IsConnectableAfter = 1

var CharacterConnectable = map[rune][]bool{
	'ا': {true, false},
	'آ': {false, false},
	'ب': {true, true},
	'پ': {true, true},
	'ت': {true, true},
	'ث': {true, true},
	'ج': {true, true},
	'چ': {true, true},
	'ح': {true, true},
	'خ': {true, true},
	'د': {true, false},
	'ذ': {true, false},
	'ر': {true, false},
	'ز': {true, false},
	'ژ': {true, false},
	'س': {true, true},
	'ش': {true, true},
	'ص': {true, true},
	'ض': {true, true},
	'ط': {true, true},
	'ظ': {true, true},
	'ع': {true, true},
	'غ': {true, true},
	'ف': {true, true},
	'ق': {true, true},
	'ک': {true, true},
	'گ': {true, true},
	'ل': {true, true},
	'م': {true, true},
	'ن': {true, true},
	'و': {true, false},
	'ه': {true, true},
	'ی': {true, true},
	' ': {false, false},
}
