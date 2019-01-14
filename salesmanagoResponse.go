package main

type SalesmanagoResponse struct {
	Success bool
	Message []string
	Result  bool
}

type HasContactResponse struct {
	SalesmanagoResponse
	ContactId string
}
