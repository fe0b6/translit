package translit

import (
	"regexp"
	"strings"
)

var (
	table map[string]string
	enNum *regexp.Regexp
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
		" ": "_",
	}

	enNum = regexp.MustCompile("[a-zA-Z0-9]")
}

func Transform(str string) string {
	sym := strings.Split(strings.TrimSpace(str), "")
	arr := make([]string, len(sym))

	for i, v := range sym {
		if enNum.MatchString(v) {
			arr[i] = v
		} else {
			arr[i] = table[v]
		}
	}

	return strings.ToLower(strings.Join(arr, ""))
}
