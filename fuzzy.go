package fuzzy

import (
	"errors"
	"sort"
)

type Fuzzy struct {
	words     []string
	syllables map[string]uint64
}

func New() Fuzzy {
	return Fuzzy{
		words:     make([]string, 0, 64),
		syllables: make(map[string]uint64),
	}
}

func (fuzzy *Fuzzy) Add(word string) error {
	if len(fuzzy.words) >= 64 {
		return errors.New("too many words")
	}

	idx := len(fuzzy.words)

	fuzzy.words = append(fuzzy.words, word)

	for i := 0; i < len(word)-1; i++ {
		syllable := word[i : i+2]
		if _, ok := fuzzy.syllables[syllable]; ok {
			fuzzy.syllables[syllable] |= 1 << idx
		} else {
			fuzzy.syllables[syllable] = 1 << idx
		}
	}

	return nil
}

func (fuzzy *Fuzzy) Search(word string) []string {

	if len(word) < 2 {
		return []string{}
	}

	words := make([]int, 0, 64)
	ranks := map[int]int{}
	for i := 0; i < len(word)-1; i++ {
		syllable := word[i : i+2]
		if num, ok := fuzzy.syllables[syllable]; ok {
			idx := 0
			for num > 0 {
				if num&1 == 1 {
					if _, ok := ranks[idx]; ok {
						ranks[idx]++
					} else {
						ranks[idx] = 1
						words = append(words, idx)
					}
				}
				num >>= 1
				idx++
			}
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return ranks[words[i]] > ranks[words[j]]
	})

	result := make([]string, 0, len(words))
	for _, word := range words {
		result = append(result, fuzzy.words[word])
	}

	return result
}
