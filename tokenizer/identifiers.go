package tokenizer

import "unicode"

func isStartOfIdentifier(r rune) bool {
	return unicode.IsLetter(r) || (r == '_')
}

func isIdentifierLetter(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func (tokenizer *Tokenizer) parseIdentifier(firstLetter rune) (*Token, error) {
	// rs, err := tokenizer.nextRunes(math.MaxInt32, func (r rune) (bool, error) {
	//   if isIdentifierLetter(r) {
	//     return false, nil
	//   }
	//
	//   if r
	// })

	return nil, nil
}
