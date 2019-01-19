package main

type ReasonCode int

const (
	UNKNOWN ReasonCode = 0 + iota
	CANNOT_POST
	TIMEOUT
	PARSE_ERROR
)

func (r ReasonCode) String() string {
	var types = [...]string{
		"UNKNOWN",
		"CANNOT_POST",
		"TIMEOUT",
		"PARSE_ERROR",
	}
	return types[r]
}
