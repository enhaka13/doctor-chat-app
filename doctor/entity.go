package doctor

import (
	"time"
)

type Doctor struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Gender         int       `json:"gender"`
	Email          string    `json:"email"`
	PhoneNumber    int       `json:"phone_number"`
	Address        string    `json:"adress"`
	Speciality     string    `json:"speciality"`
	AvatarFileName string    `json:"avatar_file_name"`
	token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
