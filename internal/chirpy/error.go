package chirpy

import "errors"

var (
	ErrChirpNotCreated    = errors.New("error while creating the chirp")
	ErrChirpTooLong       = errors.New("chirp is too long")
	ErrEmptyUserID        = errors.New("empty user_id")
	ErrEmptyBody          = errors.New("empty chirp body")
	ErrChirpsListNotFound = errors.New("chirps list not found")
	ErrChirpIDNotParsed   = errors.New("error parsing id")
	ErrChirpNotFound      = errors.New("chirp not found")
	ErrChirpIDNotFound    = errors.New("chirp id not found")
)
