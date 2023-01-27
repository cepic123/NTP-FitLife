package main

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}
