package doctor

type RegisterDoctorInput struct {
	Name        string `json:"name" binding:"required"`
	Gender      int    `json:"gender"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	Speciality  string `json:"speciality" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginDoctorInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailDoctorInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckNameDoctorInput struct {
	Name string `uri:"name" binding:"required, email"`
}

type UpdateDoctorInput struct {
	Email       string `json:"email" binding:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"adress"`
	City        string `json:"city"`
	Speciality  string `json:"speciality"`
	Password    string `json:"password"`
}
