package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jirugutema/gopastebin/internal/dto"
	"github.com/jirugutema/gopastebin/internal/repository"
	"github.com/jirugutema/gopastebin/pkg"
)

var ErrUserNotExist = errors.New("user does not exist")

func CreatePasteService(paste dto.PasteDTO) (uuid.UUID, error) {
	paste.CreatedAt = time.Now()

	// Generate paste ID
	pasteID, err := pkg.IDGenerator()
	if err != nil {
		return uuid.Nil, err
	}
	paste.ID = pasteID

	// Check user existence
	if paste.UserID != nil {
		exists, err := repository.PasteUserExists(*paste.UserID)
		if err != nil {
			return uuid.Nil, err
		}
		if !exists {
			return uuid.Nil, ErrUserNotExist
		}
	}

	// Create paste in DB
	if err := repository.CreatePaste(paste); err != nil {
		return uuid.Nil, err
	}

	return paste.ID, nil
}

func GetPasteService(pasteID uuid.UUID) (dto.PasteDTO, error) {
	res, err := repository.GetPaste(pasteID)
	if err != nil {
		return res, err
	}

	// Increment views
	_ = repository.IncrementViews(pasteID)

	// Burn after read
	if res.BurnAfterRead {
		_ = repository.DeletePaste(pasteID)
	}

	return res, nil
}
