package main

type SalesmanagoResponse struct {
	Success bool
	Message []string
	Result  bool
}

type hasContactResponse struct {
	SalesmanagoResponse
	ContactId string
}
