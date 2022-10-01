package dto

type NotesUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Date        string `json:"date" form:"date" binding:"required"`
	UserID      uint64 `json:"userid" form:"userid" binding:"required"`
	Image       string `json:"image" form:"image" binding:"required"`
}

type NotesCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Date        string `json:"date" form:"date" binding:"required"`
	UserID      uint64 `json:"userid" form:"userid" binding:"required"`
	Image       string `json:"image" form:"image" binding:"required"`
}
