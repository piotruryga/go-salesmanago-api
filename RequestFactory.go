package main

type RequestFactory interface {
	CreateRequest() AuthRequest
}

type HasContactRequest struct {
	AuthRequest
	Email string `json:"email"`
}

/*func (h *HasContactRequest) CreateRequest() AuthRequest {
	return new(HasContactRequest())
}*/

type ContactDeleteRequest struct {
	AuthRequest
	Email string `json:"email"`
}

var database map[string]interface{}

func InitRF() {
	database = make(map[string]interface{})
	hcR := new(HasContactRequest)
	database["hasContactRequest"] = hcR

	cdR := new(ContactDeleteRequest)
	database["contactDeleteRequest"] = cdR
}

func ReturnImplementation(requestName string) interface{} {
	return database[requestName]
}

/*func (h *ContactDeleteRequest) CreateRequest() AuthRequest {
	return new(ContactDeleteRequest{})
}

func HasContactRequestFactory() HasContactRequest {
	return new(HasContactRequest{})
}

func (h *AuthRequest) ContactDeleteRequestFactory() ContactDeleteRequest {
	return new(ContactDeleteRequest{})
}

func GetFactory(factoryType string) RequestFactory {
	switch factoryType {
	case "hasContact":
			return new(HasContactRequest)
	case "contactDelete":
		return new(ContactDeleteRequest)
	}

	return nil
}
*/
