package service

import (
	"fmt"
	"log"

	"github.com/IrvanWijayaSardam/mynotesapp/dto"
	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"github.com/IrvanWijayaSardam/mynotesapp/repository"
	"github.com/mashingan/smapping"
)

type NoteService interface {
	Insert(b dto.NotesCreateDTO) entity.Notes
	Update(b dto.NotesUpdateDTO) entity.Notes
	Delete(b entity.Notes)
	All() []entity.Notes
	FindByID(bookID uint64) entity.Notes
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type noteService struct {
	noteRepository repository.NotesRepository
}

func NewNoteService(noteRepo repository.NotesRepository) NoteService {
	return &noteService{
		noteRepository: noteRepo,
	}
}

func (service *noteService) Insert(b dto.NotesCreateDTO) entity.Notes {
	note := entity.Notes{}
	err := smapping.FillStruct(&note, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.noteRepository.InsertNotes(&note)
	return res
}

func (service *noteService) Update(b dto.NotesUpdateDTO) entity.Notes {
	notes := entity.Notes{}
	err := smapping.FillStruct(&notes, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v : ", err)
	}
	res := service.noteRepository.UpdateNotes(&notes)
	return res
}

func (service *noteService) Delete(b entity.Notes) {
	service.noteRepository.DeleteNotes(&b)
}

func (service *noteService) All() []entity.Notes {
	return service.noteRepository.AllNotes()
}

func (service *noteService) FindByID(notesID uint64) entity.Notes {
	return service.noteRepository.FindNotesById(notesID)
}

func (service *noteService) IsAllowedToEdit(userID string, noteID uint64) bool {
	b := service.noteRepository.FindNotesById(noteID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
