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
	GetAllDoctors() ([]Doctor, error)
	GetDoctorByID(ID int) (Doctor, error)
	GetDoctorByName(name string) (Doctor, error)
	GetDoctorByCity(city string) ([]Doctor, error)
	GetDoctorBySpeciality(speciality string) ([]Doctor, error)
	UpdateDoctor(inputName CheckNameDoctorInput, inputData UpdateDoctorInput) (Doctor, error)
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
		City:        input.City,
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

func (s *service) IsEmailAvailable(input CheckEmailDoctorInput) (bool, error) {
	email := input.Email

	doctor, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if doctor.ID == 0 {
		return true, nil
	}

	return false, nil
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

func (s *service) GetAllDoctors() ([]Doctor, error) {
	doctors, err := s.repository.FindAll()
	if err != nil {
		return []Doctor{}, err
	}

	return doctors, err
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

func (s *service) GetDoctorByCity(city string) ([]Doctor, error) {
	doctors, err := s.repository.FindByCity(city)
	if err != nil {
		return []Doctor{}, err
	}

	// for _, doctor := range doctors {
	// 	if doctor.City == "" {
	// 		return []Doctor{}, errors.New("No doctor found on that city")
	// 	}
	// }

	return doctors, nil
}
func (s *service) GetDoctorBySpeciality(speciality string) ([]Doctor, error) {
	doctors, err := s.repository.FindBySpeciality(speciality)
	if err != nil {
		return []Doctor{}, err
	}

	// for _, doctor := range doctors {
	// 	if doctor.Speciality == "" {
	// 		return []Doctor{}, errors.New("No doctor found on that speciality")
	// 	}
	// }

	return doctors, nil
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

func (s *service) DeleteDoctor(input CheckNameDoctorInput) error {
	doctor, err := s.repository.FindByName(input.Name)
	if err != nil {
		return err
	}

	err = s.repository.Delete(doctor.ID)
	if err != nil {
		return err
	}

	return nil
}
