package user

type GetUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// TODO should implement orders dto here
}
