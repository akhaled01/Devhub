# About the `Types` directory

Each file here houses non-DB entity specific methods and structs. For example, the `user.go` file contains the main user entity (struct) which is:

```go
type User struct {
	ID        uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	Avatar    string    `json:"avatar"`
}
```

Structs in this directory are also the main interface between the `vite.js` frontend and the backend
