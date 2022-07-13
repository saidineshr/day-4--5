package Models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	LastName string `json:"last_name"`
	DOB      string `json:"dob"`
	Subject  string `json:"subject"`
	Marks    uint   `json:"marks"`
	Address  string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}
