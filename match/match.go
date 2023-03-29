package match

import "strings"

type Match struct {
	Matches []string

	Fuzzy        []string
	Prefixes     []string
	Suffixes     []string
	DoubleRights []string

	Musts    []string
	Excludes []string
}

func (m Match) Match(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	for _, exclude := range m.Excludes {
		if strings.Contains(s, exclude) {
			return false
		}
	}
	for _, must := range m.Musts {
		if !strings.Contains(s, must) {
			return false
		}
	}

	for _, match := range m.Matches {
		if strings.Contains(s, match) {
			return true
		}
	}

	for _, fuzzy := range m.Fuzzy {
		for _, prefix := range m.Prefixes {
			if strings.Contains(s, prefix+fuzzy) {
				return true
			}
		}
		for _, suffix := range m.Suffixes {
			if strings.Contains(s, fuzzy+suffix) {
				return true
			}
		}
		for _, doubleRight := range m.DoubleRights {
			if strings.Contains(s, fuzzy) && strings.Contains(s, doubleRight) {
				return true
			}
		}
	}
	return false
}
