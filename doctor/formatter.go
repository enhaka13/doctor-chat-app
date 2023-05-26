package doctor

type DoctorFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Speciality  string `json:"speciality"`
	ImageURL    string `json:"image_url"`
	Token       string `json:"token"`
}

func FormatDoctor(doctor Doctor, token string) DoctorFormatter {
	formatter := DoctorFormatter{
		ID:          doctor.ID,
		Name:        doctor.Name,
		Gender:      doctor.Gender,
		Email:       doctor.Email,
		PhoneNumber: doctor.PhoneNumber,
		Address:     doctor.Address,
		City:        doctor.City,
		Speciality:  doctor.Speciality,
		ImageURL:    doctor.AvatarFileName,
		Token:       token,
	}

	return formatter
}
