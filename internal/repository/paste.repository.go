//Package repository interacts with database
package repository

import (
	"github.com/google/uuid"
	"github.com/jirugutema/gopastebin/internal/config"
	"github.com/jirugutema/gopastebin/internal/dto"
)

func CreatePaste(paste dto.PasteDTO) error {
	query := `
        INSERT INTO pastes (
            id, title, content, language, password, burn_after_read, expires_at, created_at, views, user_id, is_public
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
    `
	_, err := config.DB.Exec(
		query,
		paste.ID,
		paste.Title,
		paste.Content,
		paste.Language,
		paste.Password,
		paste.BurnAfterRead,
		paste.ExpiresAt,
		paste.CreatedAt,
		paste.Views,
		paste.UserID,
		paste.IsPublic,
	)
	return err
}

func GetPaste(pasteID uuid.UUID) (dto.PasteDTO, error) {
	var res dto.PasteDTO
	query := `
        SELECT id, title, content, language, password, burn_after_read, expires_at, created_at, views, user_id, is_public
        FROM pastes
        WHERE id = $1
    `
	row := config.DB.QueryRow(query, pasteID)
	err := row.Scan(
		&res.ID,
		&res.Title,
		&res.Content,
		&res.Language,
		&res.Password,
		&res.BurnAfterRead,
		&res.ExpiresAt,
		&res.CreatedAt,
		&res.Views,
		&res.UserID,
		&res.IsPublic,
	)
	return res, err
}

func IncrementViews(pasteID uuid.UUID) error {
	_, err := config.DB.Exec("UPDATE pastes SET views = views + 1 WHERE id = $1", pasteID)
	return err
}

func DeletePaste(pasteID uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM pastes WHERE id = $1", pasteID)
	return err
}

func PasteUserExists(userID uuid.UUID) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", userID).Scan(&exists)
	return exists, err
}
