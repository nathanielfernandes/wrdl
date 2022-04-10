package wrdl

import (
	"strings"
)

type Solver struct {
	must_contain     string
	must_not_contain string
	potential_words  []string
	letters          *[]Letter
}

func InitialSolver() Solver {
	letters := fresh()
	return Solver{
		must_contain:     "",
		must_not_contain: "",
		potential_words:  words,
		letters:          &letters,
	}
}

func Results(gs Guesses) []string {
	s := InitialSolver()
	for _, g := range gs {
		s.update(g)
	}

	return s.potential_words
}

func (s *Solver) update(guess []Guess) {
	for i, g := range guess {
		switch g.v {
		case Incorrect:
			s.must_not_contain += g.c
		case Valid:
			s.must_contain += g.c
			(*s.letters)[i].never_letters += g.c
		case Determined:
			(*s.letters)[i].determined = rune(g.c[0])
		}
	}

	filtered := []string{}
	for _, word := range s.potential_words {
		if s.filter(word) {
			filtered = append(filtered, word)
		}
	}

	s.potential_words = filtered
}

func (s Solver) filter(word string) bool {
	for _, c := range s.must_contain {
		if !strings.ContainsRune(word, c) {
			return false
		}
	}

	for i, c := range word {
		l := (*s.letters)[i]

		if !l.exact(c) && l.no_empty() || (strings.ContainsRune(s.must_not_contain, c) || l.invalid(c)) {
			return false
		}

	}

	return true
}
