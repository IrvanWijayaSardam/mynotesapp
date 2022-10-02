package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password"`
	Profile  string `json:"profile" form:"profile" binding:"required"`
	Jk       string `json:"jk" form:"jk" binding:"required"`
}

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password" form:"password" validate:"min:5" binding:"required"`
// }
