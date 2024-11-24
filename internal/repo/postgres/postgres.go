package postgres

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/singl3focus/em_testtask/internal/models"
)

type Repository struct {
	logger *slog.Logger
	db *sqlx.DB
}

func NewPostgresDB(authLink string, logger *slog.Logger) (*Repository, error) {
	db, err := sqlx.Open("postgres", authLink)
	if err != nil {
		return nil, fmt.Errorf("db connect error: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping error: %s", err.Error())
	}

	dbMigrations := db.DB
	if err := goose.Run("up", dbMigrations, "migrations"); err != nil {
        return nil, fmt.Errorf("goose up migrations error: %s", err.Error())
    }

	return &Repository{
		db: db,
		logger: logger,
	}, nil
}

func (r *Repository) AddSong(song models.Song) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

    query1 := `INSERT INTO songs (group_name, song_title) VALUES ($1, $2) RETURNING id`

	var songID int
	err = tx.Get(&songID, query1, song.Group, song.Title)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting songs: %v", err)
	}

	query2 := `INSERT INTO verses (song_id, verse_number, text) VALUES ($1, $2, $3)`
    for _, verse := range song.Verses {
		_, err := tx.Exec(query2, songID, verse.Number, verse.Text)
        if err != nil {
            tx.Rollback()
            return fmt.Errorf("error inserting verses: %v", err)
        }
	}

    if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

    return nil
}

func (r *Repository) RemoveSong(groupName, songTitle string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

    query1 := `SELECT id FROM songs WHERE group_name = $1 AND song_title = $2`

	var songID int
	err = tx.Get(&songID, query1, groupName, songTitle)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error get song_id from songs: %v", err)
	}

	query2 := `DELETE FROM verses WHERE song_id = $1`
    _, err = tx.Exec(query2, songID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete verses: %v", err)
	}

	query3 := `DELETE FROM songs WHERE id = $1`
	_, err = tx.Exec(query3, songID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete song info: %v", err)
	}	

    if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

    return nil
}

func (r *Repository) UpdateSongInfo(oldGroupName, oldSongTitle string, newGroupName, newSongTitle string) error {
	query := `UPDATE songs SET group_name = $1, song_title = $2 WHERE group_name = $3 AND song_title = $4`

	_, err := r.db.Exec(query, newGroupName, newSongTitle, oldGroupName, oldSongTitle)
	if err != nil {
		return fmt.Errorf("error inserting verses: %v", err)
	}

	return nil
}

func (r *Repository) GetSongTextByVerses(groupName, songTitle string, offset, limit int) (string, error) {
	query1 := `SELECT id FROM songs WHERE group_name = $1 AND song_title = $2`

	var songID int
	err := r.db.Get(&songID, query1, groupName, songTitle)
	if err != nil {
		return "", fmt.Errorf("error get song_id from songs: %v", err)
	}

	query2 := `SELECT text FROM verses WHERE song_id = $1 ORDER BY verse_number LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query2, songID, limit, offset)
	if err != nil {
		return "", fmt.Errorf("error get song_id from songs: %v", err)
	}
	defer rows.Close()

	songText := ""
    for rows.Next() {
        var verseText string

        err := rows.Scan(&verseText)
        if err != nil {
            return "", fmt.Errorf("error scanning verse row: %v", err)
        }

		songText += verseText
		if rows.Next() {
			songText += "\n\n"
		}
    }

    if err = rows.Err(); err != nil {
        return "", fmt.Errorf("error reading rows: %v", err)
    }

    return songText, nil
}

func (r *Repository) GetSongsInfo(groupName, songTitle string, offset, limit int) ([]models.SongInfo, error) {
	query := `SELECT group_name, song_title FROM songs WHERE group_name ILIKE $1 AND song_title ILIKE $2 LIMIT $3 OFFSET $4`
	rows, err := r.db.Query(query, groupName, songTitle, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error get from songs: %v", err)
	}
	defer rows.Close()

	songsInfo := make([]models.SongInfo, 0, 1)
    for rows.Next() {
        var song models.SongInfo

        err := rows.Scan(&song.Group, &song.Title)
        if err != nil {
            return nil, fmt.Errorf("error scanning songs row: %v", err)
        }

		songsInfo = append(songsInfo, song)
    }

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error reading rows: %v", err)
    }

    return songsInfo, nil
}