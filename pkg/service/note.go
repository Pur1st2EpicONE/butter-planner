package service

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
)

type NoteService struct {
	storage repository.NoteMaker
}

func NewNoteService(storage repository.NoteMaker) *NoteService {
	return &NoteService{storage: storage}
}

func (ns *NoteService) CreateNote(userId int, note butterplanner.Note) (int, error) {
	return ns.storage.CreateNote(userId, note)
}

func (ns *NoteService) GetAllNotes(userId int) ([]butterplanner.Note, error) {
	return ns.storage.GetAllNotes(userId)
}
