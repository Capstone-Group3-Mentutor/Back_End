package repository

import (
	"be12/mentutor/features/login"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Name      string
	Images    string
	Email     string
	Password  string
	Role      string
	IdClass   uint
	ClassName string
}

type Mentee struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Images    string
	Role      string
	IdClass   uint
	ClassName string
}

type GoogleToken struct {
	gorm.Model
	IdMentee     uint
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
}

type Class struct {
	ClassName string
}

func ToDomainMentor(u Mentor) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Images:   u.Images,
		Role:     u.Role,
		IdClass:  u.IdClass,
		Class:    u.ClassName,
	}
}

func ToDomainMentee(u Mentee) login.Core {
	return login.Core{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Images:   u.Images,
		Role:     u.Role,
		IdClass:  u.IdClass,
		Class:    u.ClassName,
	}
}

func FromDomain(du login.Core) Mentor {
	return Mentor{
		Email:    du.Email,
		Password: du.Password,
	}
}

func FromDomainMentee(du login.Core) Mentee {
	return Mentee{
		Email:    du.Email,
		Password: du.Password,
	}
}
