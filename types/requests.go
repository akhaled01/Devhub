package types

// to be decoded from a signup request
type SignupRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int64  `json:"age"`
	Gender    string `json:"gender"`
	Avatar    string `json:"image"`
}

type LoginRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`  
}
