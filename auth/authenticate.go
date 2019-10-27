package auth

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var users = map[string]string{
	"admin": "123qweasd",
	"user":  "password",
}

// User presentation
type User struct {
	Username string
	Password string
}

// ParseHeader parses "Basic c2ZkOmFzZmQ=" texts
func ParseHeader(headerValue string) (*User, error) {
	encrypted := strings.Split(headerValue, " ")[1]
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, fmt.Errorf("Basic Auth header parse error: %s", err)
	}
	credentials := strings.Split(string(data), ":")
	u := new(User)
	u.Username = credentials[0]
	u.Password = credentials[1]
	return u, nil
}

// Authenticate func
func (u *User) Authenticate() bool {
	password, ok := users[u.Username]
	if !ok || password != u.Password {
		return false
	}
	return true
}
