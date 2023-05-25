package doctor

import "gorm.io/gorm"

type Repository interface {
	Create(doctor Doctor) (Doctor, error)
	FindByID(ID int) (Doctor, error)
	FindByEmail(email string) (Doctor, error)
	FindByName(name string) (Doctor, error)
	FindByAddress(address string) (Doctor, error)
	FindBySpeciality(speciality string) (Doctor, error)
	Update(doctor Doctor) (Doctor, error)
	Delete(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(doctor Doctor) (Doctor, error) {
	if err := r.db.Create(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) FindByID(ID int) (Doctor, error) {
	var doctor Doctor

	if err := r.db.Where("id = ?", ID).Find(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) FindByEmail(email string) (Doctor, error) {
	var doctor Doctor

	if err := r.db.Where("email = ?", email).Find(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) FindByName(name string) (Doctor, error) {
	var doctor Doctor

	if err := r.db.Where("name = ?", name).Find(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) FindByAddress(address string) (Doctor, error) {
	var doctor Doctor

	if err := r.db.Where("address = ?", address).Find(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) FindBySpeciality(speciality string) (Doctor, error) {
	var doctor Doctor

	if err := r.db.Where("speciality = ?", speciality).Find(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) Update(doctor Doctor) (Doctor, error) {
	if err := r.db.Save(&doctor).Error; err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (r *repository) Delete(ID int) error {
	var doctor Doctor
	if err := r.db.Where("id = ?", ID).Delete(&doctor).Error; err != nil {
		return err
	}

	return nil
}
