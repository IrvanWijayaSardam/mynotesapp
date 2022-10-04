package entity

type Notes struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Date        string `gorm:"type:varchar(255)" json:"date"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
}
