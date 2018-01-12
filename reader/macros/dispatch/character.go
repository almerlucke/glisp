package dispatch

import (
	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/characters"
)

// CharacterDispatch character dispatch macro
func CharacterDispatch(arg uint64, rd reader.Reader) (types.Object, error) {
	err := rd.UnreadChar()
	if err != nil {
		return nil, err
	}

	token, err := rd.ParseToken(true)
	if err != nil {
		return nil, err
	}

	return characters.NewWithToken(token)
}
