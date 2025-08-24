package entity

type UserEntity struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	Sector   string `json:"sector" gorm:"type:varchar(100);not null"`
	Role     string `json:"role" gorm:"type:varchar(50);not null"`
	IsActive bool   `json:"is_active" gorm:"type:boolean;default:true"`
}
