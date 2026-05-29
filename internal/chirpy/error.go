package chirpy

import "errors"

var (
	ErrChirpNotCreated = errors.New("error while creating the chirp")
	ErrChirpTooLong    = errors.New("chirp is too long")
	ErrEmptyUserID     = errors.New("empty user_id")
	ErrEmptyBody       = errors.New("empty chirp body")
)
