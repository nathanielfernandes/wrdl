package wrdl

import "strings"

type Letter struct {
	determined    rune
	never_letters string
}

func (l Letter) exact(c rune) bool {
	return l.determined == c
}

func (l Letter) invalid(c rune) bool {
	return strings.ContainsRune(l.never_letters, c)
}

func (l Letter) no_empty() bool {
	return l.determined != ' '
}

func fresh() []Letter {
	batch := []Letter{}
	for i := 0; i < 5; i++ {
		batch = append(batch, Letter{' ', ""})
	}

	return batch
}
