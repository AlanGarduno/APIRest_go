package structures

type User struct {
	Id 		int      `json:"id"`
	UserName string  `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
}

type Response struct{
	Status string `json:"status"`
	Data User     `json:"data"`
	Message string `json:"message"`
}