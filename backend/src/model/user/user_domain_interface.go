package model

import "github.com/kevynlohan05/StockPro/src/configuration/rest_err"

type UserDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetSector() string
	GetRole() string
	GetIsActive() bool

	SetID(id string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(name, email, password, sector, role string, isActive bool) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		sector:   sector,
		role:     role,
		isActive: isActive,
	}
}

func NewUserUpdateDomain(name, email, password, sector, role string, isActive bool) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
		sector:   sector,
		role:     role,
		isActive: isActive,
	}
}

func NewLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
