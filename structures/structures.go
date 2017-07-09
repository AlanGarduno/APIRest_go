package structures

type User struct {
	Id 		int      `json:id`
	UserName string  `json:UserName`
	FirstName string `json:FirstName`
	LastName string `json:LastName`
}