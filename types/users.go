package types

// Credentials is the data users enter in the FE to login
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Users is a slice of User
type Users []User

//User is the format of a User in the database
type User struct {
	ID             int
	Username       string
	HashedPassword string
}
