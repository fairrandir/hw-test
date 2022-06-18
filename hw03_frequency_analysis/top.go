package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

func Top10(text string) (res []string) {
	var (
		freqs map[string]int
		words []string
		i     int
		ok    bool
	)
	if len(text) == 0 {
		return
	}
	words = strings.Fields(text)
	// Сразу аллоцируем память, чтобы точно хватило даже в худшем случае: если все слова уникальны
	freqs = make(map[string]int, len(words))
	for i = 0; i < len(words); i++ {
		// Нормализуем слово - переведем в нижний регистр, сносим спецсимволы в начале и конце слова
		tmpWord := strings.TrimFunc(strings.ToLower(words[i]), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		// Если всё слово и состояло из одних спецсимволов - игнорируем его
		if tmpWord == "" {
			continue
		}
		// Если слово ещё не встречалось в словаре - добавим его в итоговый слайс
		if _, ok = freqs[tmpWord]; !ok {
			res = append(res, tmpWord)
		}
		// Увеличим в словаре частотность слова
		freqs[tmpWord]++
	}
	// Отсортируем итоговый слайс по частотности слов
	sort.Slice(res, func(i, j int) bool {
		// Если частотность одинаковая - то первым будет то, которое раньше по алфавиту
		if freqs[res[i]] == freqs[res[j]] {
			return res[i] < res[j]
		}
		// Иначе - которое встречается в тексте чаще
		return freqs[res[i]] > freqs[res[j]]
	})
	// Если слов в итоге больше 10 - усечем итоговый слайс
	if len(res) > 10 {
		res = res[:10]
	}
	return
}
