package graph

import "strconv"

// isColor(c) returns true if c is a valid color, false otherwise
// STUB: dipendente dalla rappresentazione scelta per i colori;
//	enumerazione, lista di color-names, codifica RGB/CMYK, hex code, interi,...
func IsColor(color string) bool {
	return true
}

// toColor(c) returns the color associated to the integer c
// STUB: dipendente dalla rappresentazione scelta per i colori;
//	enumerazione, lista di color-names, codifica RGB/CMYK, hex code, interi,...
func ToColor(c int) string{
	return strconv.Itoa(c)
}

// toColors(cs) takes an array of integers cs and returns an array of strings s such that
// s[i] = toColor(cs[i])
func ToColors(colors []int) (s []string) {
	s = make([]string, len(colors))
	for i,color := range colors{
		s[i] = ToColor(color)
	}

	return
}

// contains(e, v) returns true if v is an element of e, false otherwise
func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
			if v == s {
					return true
			}
	}
	return false
}