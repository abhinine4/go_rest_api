package service

import (
	"go-fiber-crud/data/request"
	"go-fiber-crud/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteResponse)
	Delete(nodeId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
