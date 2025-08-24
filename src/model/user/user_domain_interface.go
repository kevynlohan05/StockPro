package model

type UserDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetSector() string
	GetRole() string
	GetIsActive() bool

	EncryptPassword()
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
