package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Save(user User) (User, error){
	data := r.db.Create(&user)

	if data.Error != nil {
		return user, data.Error
	}
	return user, nil
}

