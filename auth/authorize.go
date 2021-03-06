package auth

var accesses = map[string][]Access{
	"admin": []Access{
		Access{
			Type: "registry",
			Name: "catalog",
			Actions: []string{
				"*",
			},
		},
	},
	"user": []Access{
		Access{
			Type: "repository",
			Name: "hello-world",
			Actions: []string{
				"pull",
			},
		},
	},
}

// Authorize func
func (u *User) Authorize() []Access {
	return accesses[u.Username]
}
