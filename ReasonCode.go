package main

type ReasonCode int

const (
	OK ReasonCode = 0 + iota
	CANNOT_POST
	TIMEOUT
	PARSE_ERROR
)

func (r ReasonCode) String() string {
	var types = [...]string{
		"OK",
		"CANNOT_POST",
		"TIMEOUT",
		"PARSE_ERROR",
	}
	return types[r]
}
