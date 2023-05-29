package graph

// isColor(c) returns true if c is a valid color, false otherwise
// STUB: dipendente dalla rappresentazione scelta per i colori;
//
//	enumerazione, lista di color-names, codifica RGB/CMYK, hex code, interi,...
func isColor(color string) bool {
	return true
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