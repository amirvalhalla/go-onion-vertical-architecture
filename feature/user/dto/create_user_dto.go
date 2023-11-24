package user

type CreateDto struct {
	FirstName string `binding:"required,min=3,max=128" json:"firstName"`
	LastName  string `binding:"required,min=3,max=128" json:"lastName"`
}
