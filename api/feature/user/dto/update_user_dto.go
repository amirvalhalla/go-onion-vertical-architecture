package user

type UpdateUserDto struct {
	FirstName string `binding:"required,min=3,max=128" json:"firstName"`
	LastName  string `binding:"required,min=3,max=128" json:"lastName"`
	Suspend   bool   `binding:"required" json:"suspend"`
}
