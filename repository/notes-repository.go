package repository

import (
	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"gorm.io/gorm"
)

type NotesRepository interface {
	InsertNotes(b *entity.Notes) entity.Notes
	UpdateNotes(b *entity.Notes) entity.Notes
	DeleteNotes(b *entity.Notes)
	AllNotes() []entity.Notes
	FindNotesById(notesUserID uint64) entity.Notes
}

type notesConnection struct {
	connection *gorm.DB
}

//NewNotesRepository .asd
func NewNotesRepository(dbConn *gorm.DB) NotesRepository {
	return &notesConnection{
		connection: dbConn,
	}
}

func (db *notesConnection) InsertNotes(b *entity.Notes) entity.Notes {
	db.connection.Save(&b)
	db.connection.Preload("UserID").Find(&b)
	return *b
}

func (db *notesConnection) UpdateNotes(b *entity.Notes) entity.Notes {
	db.connection.Save(&b)
	db.connection.Preload("UserID").Find(&b)
	return *b
}

func (db *notesConnection) DeleteNotes(b *entity.Notes) {
	db.connection.Delete(&b)
}

func (db *notesConnection) FindNotesById(UserID uint64) entity.Notes {
	var note entity.Notes
	db.connection.Preload("user_id").Find(&note, UserID)
	return note
}

func (db *notesConnection) AllNotes() []entity.Notes {
	var notes []entity.Notes
	db.connection.Preload("UserID").Find(&notes)
	return notes
}
