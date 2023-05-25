package doctor

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterDoctor(input RegisterDoctorInput) (Doctor, error)
	Login(input LoginDoctorInput) (Doctor, error)
	IsEmailAvailable(input CheckEmailDoctorInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (Doctor, error)
	GetDoctorByID(ID int) (Doctor, error)
	GetDoctorByName(name string) (Doctor, error)
	GetDoctorByCity(city string) (Doctor, error)
	GetDoctorBySpeciality(speciality string) (Doctor, error)
	UpdateDoctor(input RegisterDoctorInput) (Doctor, error)
	DeleteDoctor(input CheckNameDoctorInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterDoctor(input RegisterDoctorInput) (Doctor, error) {
	doctor := Doctor{
		Name:        input.Name,
		Gender:      input.Gender,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Speciality:  input.Speciality,
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return Doctor{}, err
	}

	doctor.Password = string(PasswordHash)

	newDoctor, err := s.repository.Create(doctor)
	if err != nil {
		return Doctor{}, err
	}

	return newDoctor, nil
}

func (s *service) Login(input LoginDoctorInput) (Doctor, error) {
	email := input.Email
	password := input.Password

	doctor, err := s.repository.FindByEmail(email)
	if err != nil {
		return Doctor{}, err
	}

	if doctor.ID == 0 {
		return Doctor{}, errors.New("No doctor found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(password))
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (Doctor, error) {
	doctor, err := s.repository.FindByID(ID)
	if err != nil {
		return Doctor{}, err
	}

	doctor.AvatarFileName = fileLocation

	updatedDoctor, err := s.repository.Update(doctor)
	if err != nil {
		return Doctor{}, err
	}

	return updatedDoctor, nil
}

func (s *service) GetDoctorByID(ID int) (Doctor, error) {
	doctor, err := s.repository.FindByID(ID)
	if err != nil {
		return Doctor{}, err
	}

	if doctor.ID == 0 {
		return Doctor{}, errors.New("No doctor found on that ID")
	}

	return doctor, nil
}

func (s *service) GetDoctorByName(name string) (Doctor, error) {
	doctor, err := s.repository.FindByName(name)
	if err != nil {
		return Doctor{}, err
	}

	if doctor.Name == "" {
		return Doctor{}, errors.New("No doctor found on that name")
	}

	return doctor, nil
}

func (s *service) GetDoctorByCity(city string) (Doctor, error) {
	doctor, err := s.repository.FindByCity(city)
	if err != nil {
		return Doctor{}, err
	}

	if doctor.City == "" {
		return Doctor{}, errors.New("No doctor found on that city")
	}

	return doctor, nil
}
func (s *service) GetDoctorBySpeciality(speciality string) (Doctor, error) {
	doctor, err := s.repository.FindBySpeciality(speciality)
	if err != nil {
		return Doctor{}, err
	}

	if doctor.Speciality == "" {
		return Doctor{}, errors.New("No doctor found on that speciality")
	}

	return doctor, nil
}

func (s *service) UpdateDoctor(inputName CheckNameDoctorInput, inputData UpdateDoctorInput) (Doctor, error) {
	doctor, err := s.repository.FindByName(inputName.Name)
	if err != nil {
		return Doctor{}, err
	}

	doctor.Email = inputData.Email
	doctor.PhoneNumber = inputData.PhoneNumber
	doctor.Address = inputData.Address
	doctor.City = inputData.City
	doctor.Speciality = inputData.Speciality

	updatedDoctor, err := s.repository.Update(doctor)
	if err != nil {
		return Doctor{}, err
	}

	return updatedDoctor, nil
}
