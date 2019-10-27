package auth

var accesses = map[string][]access{
	"admin": []access{
		access{
			Type: "repository",
			Name: "hello-world",
			Actions: []string{
				"pull",
				"push",
			},
		},
	},
	"user": []access{
		access{
			Type: "repository",
			Name: "hello-world",
			Actions: []string{
				"pull",
			},
		},
	},
}

// Authorize func
func (u *User) Authorize() []access {
	return accesses[u.Username]
}
