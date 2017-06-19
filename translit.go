package translit

import (
	"strings"
)

var (
	table map[string]string
)

func init() {
	table = map[string]string{
		"А": "A", "а": "a",
		"Б": "B", "б": "b",
		"В": "V", "в": "v",
		"Г": "G", "г": "g",
		"Д": "D", "д": "d",
		"Е": "E", "е": "e",
		"Ё": "E", "ё": "e",
		"Ж": "J", "ж": "j",
		"З": "Z", "з": "z",
		"И": "I", "и": "i",
		"Й": "I", "й": "i",
		"К": "K", "к": "k",
		"Л": "L", "л": "l",
		"М": "M", "м": "m",
		"Н": "N", "н": "n",
		"О": "O", "о": "o",
		"П": "P", "п": "p",
		"Р": "R", "р": "r",
		"С": "S", "с": "s",
		"Т": "T", "т": "t",
		"У": "U", "у": "u",
		"Ф": "F", "ф": "f",
		"Х": "H", "х": "h",
		"Ц": "C", "ц": "c",
		"Ч": "Ch", "ч": "ch",
		"Ш": "Sh", "ш": "sh",
		"Щ": "Sc", "щ": "sc",
		"Ъ": "", "ъ": "",
		"Ы": "Y", "ы": "y",
		"Ь": "", "ь": "",
		"Э": "E", "э": "e",
		"Ю": "Iu", "ю": "iu",
		"Я": "Ia", "я": "ia",
		" ": "_", "1": "1",
		"2": "2", "3": "3",
		"4": "4", "5": "5",
		"6": "6", "7": "7",
		"8": "8", "9": "9",
		"0": "0",
	}
}

func Transform(str string) string {
	arr := []string{}
	for _, v := range strings.Split(strings.TrimSpace(str), "") {
		arr = append(arr, table[v])
	}

	return strings.ToLower(strings.Join(arr, ""))
}
