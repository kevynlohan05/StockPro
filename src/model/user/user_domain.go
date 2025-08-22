package model

type userDomain struct {
	iD       string
	name     string
	email    string
	password string
	sector   string
	role     string
	isActive bool
}

func (ud *userDomain) GetID() string {
	return ud.iD
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetSector() string {
	return ud.sector
}

func (ud *userDomain) GetRole() string {
	return ud.role
}

func (ud *userDomain) GetIsActive() bool {
	return ud.isActive
}
