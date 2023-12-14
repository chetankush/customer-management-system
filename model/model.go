package model

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	DateOfBirth string `json:"date_of_birth"`
	IsActive    bool   `json:"is_active"`
}
