package usecase

import (
	repo "rise-nostr/pkg/app/event/repo/postgre"
	"rise-nostr/pkg/models"
)

type Handler struct {
}

func NewEventHandler() *Handler {
	return &Handler{}
}

func (t *Handler) SaveEvent(data models.Event) error {
	return repo.SaveEvent(data)
}

func (t *Handler) GetEvent(limit int) []models.Event {
	return repo.GetEvent(limit)
}
