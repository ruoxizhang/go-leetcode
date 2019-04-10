package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"a", "b", "a"}
	s := "abababab"
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	totalWords := len(words)
	wordsLength := len(words[0])

	constWordsMap := make(map[string]int)
	for _, word := range words {
		if value, ok := constWordsMap[word]; ok {
			constWordsMap[word] = value + 1
		} else {
			constWordsMap[word] = 1
		}
	}

	result := make([]int, 0)

	stringLen := len(s)
	for i := 0; i < wordsLength; i++ {
		tempMap := makeACopyMap(constWordsMap)
		for j := i; j+wordsLength <= stringLen; j = j + wordsLength {
			thisWord := s[j : j+wordsLength]
			// check if it is a word
			if _, ok := constWordsMap[thisWord]; ok {
				// check if the word is in tempMap
				if value, ok := tempMap[thisWord]; ok {
					//reduce word number by 1
					if value <= 1 {
						delete(tempMap, thisWord)
					} else {
						tempMap[thisWord] = value - 1
					}

					if len(tempMap) == 0 {
						result = append(result, j-(totalWords-1)*wordsLength)
						firstWordInThisSubString := s[j-(totalWords-1)*wordsLength : j-(totalWords-2)*wordsLength]
						tempMap[firstWordInThisSubString] = 1
					}
				} else {
					for k := totalWords - 1; k >= 0; k-- {
						if j-k*wordsLength < 0 {
							continue
						}
						if s[j-k*wordsLength:j-(k-1)*wordsLength] == thisWord {
							tempMap = makeACopyMap(constWordsMap)
							for m := j - (k-1)*wordsLength; m <= j; m = m + wordsLength {
								afterThisWord := s[m : m+wordsLength]
								if value, ok := tempMap[afterThisWord]; ok {
									if value <= 1 {
										delete(tempMap, afterThisWord)
									} else {
										tempMap[afterThisWord] = value - 1
									}
								}
							}
							break
						}
					}
				}
			} else {
				tempMap = makeACopyMap(constWordsMap)
			}
		}
	}
	return result
}

func findSubstringNew(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	if len(words[0]) == 0 || len(words[0])*len(words) > len(s) {
		return []int{}
	}

	totalWords := len(words)
	wordsLength := len(words[0])
	constWordsMap := make(map[string]int)
	for _, word := range words {
		constWordsMap[word]++
	}

	result := make([]int, 0)

	stringLen := len(s)
	for i := 0; i < wordsLength; i++ {
		tempMap := makeACopyMap(constWordsMap)
		for j := i; j+wordsLength <= stringLen; j = j + wordsLength {
			thisWord := s[j : j+wordsLength]
			// check if it is a word
			if _, ok := constWordsMap[thisWord]; ok {
				// check if the word is in tempMap
				if value, ok := tempMap[thisWord]; ok {
					//reduce word number by 1
					if value <= 1 {
						delete(tempMap, thisWord)
					} else {
						tempMap[thisWord] = value - 1
					}

					if len(tempMap) == 0 {
						result = append(result, j-(totalWords-1)*wordsLength)
						firstWordInThisSubString := s[j-(totalWords-1)*wordsLength : j-(totalWords-2)*wordsLength]
						tempMap[firstWordInThisSubString] = 1
					}
				} else {
					for k := totalWords - 1; k >= 0; k-- {
						if j-k*wordsLength < 0 {
							continue
						}
						if s[j-k*wordsLength:j-(k-1)*wordsLength] == thisWord {
							tempMap = makeACopyMap(constWordsMap)
							for m := j - (k-1)*wordsLength; m <= j; m = m + wordsLength {
								afterThisWord := s[m : m+wordsLength]
								if value, ok := tempMap[afterThisWord]; ok {
									if value <= 1 {
										delete(tempMap, afterThisWord)
									} else {
										tempMap[afterThisWord] = value - 1
									}
								}
							}
							break
						}
					}
				}
			} else {
				tempMap = makeACopyMap(constWordsMap)
			}
		}
	}
	return result
}

func makeACopyMap(originalMap map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

func findSubstringGood(s string, words []string) []int {

	res := []int{}

	if len(words) == 0 || len(s) == 0 {
		return res
	}

	if len(words[0]) == 0 || len(words[0])*len(words) > len(s) {
		return res
	}

	wordLen := len(words[0])
	wordNum := len(words)
	slen := wordLen * wordNum

	m := make(map[string]byte)

	for i := 0; i < wordNum; i++ {
		m[words[i]]++
	}

	for i := 0; i < wordLen; i++ {

		b := i
		e := i + wordLen

		opm := make(map[string]byte)

		for len(s)-b >= slen {
			k := s[e-wordLen : e]
			// 不存在这个单词
			//fmt.Println("word is ",k)
			if v, ok := m[k]; !ok {
				b = e
				e += wordLen
				if len(opm) != 0 {
					opm = make(map[string]byte)
				}

				continue
			} else {
				//fmt.Println("m[k]" , m[k], " omp[k]", opm[k])
				if w, ok := opm[k]; ok && v == w {
					//fmt.Println("reach  is ", k , " mk is ", m[k], " opmk is ", opm[k])
					e1 := strings.Index(s[b:e], s[e-wordLen:e])

					for y := b; y <= e1; y += wordLen {
						//fmt.Println("i is ",i," b is ",b," sub ss is e1 is ", s[y:y+wordLen], e1, " e is ", e)

						z := s[y : y+wordLen]
						if z != k { //避免第一个就是那个类似的单词，反倒是减一了
							opm[z]--
						}

					}
					b += e1 + wordLen //跳过第一个相同的单词，这样才能保持总的这个单词的数量不变
					//fmt.Println(" b update is ", b)

				} else {
					opm[k]++

				}
			}
			//fmt.Println("b is ", b , " e is ",e, " len is ", e -b  )
			if e-b == slen {

				res = append(res, b)
				opm[s[b:b+wordLen]]--
				b += wordLen
				e += wordLen

			} else {
				e += wordLen
			}
		}
	}
	return res
}
