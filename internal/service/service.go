package service

import (
	"log/slog"

	"github.com/singl3focus/em_testtask/internal/models"
)

type Repository interface {
	AddSong(song models.Song) error
	RemoveSong(groupName, songTitle string) error
	UpdateSongInfo(oldGroupName, oldSongTitle string, newGroupName, newSongTitle string) error
	GetSongTextByVerses(groupName, songTitle string, offset, limit int) (string, error)
	GetSongsInfo(groupName, songTitle string, offset, limit int) ([]models.SongInfo, error)
}

type Service struct {
	repo Repository
	logger *slog.Logger
}

func NewService(repo Repository, logger *slog.Logger) *Service {
	return &Service{
		repo: repo,
		logger: logger,
	}
}

func (s *Service) AddSong(song models.Song) error {
	return s.repo.AddSong(song)
}

func (s *Service) RemoveSong(groupName, songTitle string) error {
	return s.repo.RemoveSong(groupName, songTitle)
}

func (s *Service) UpdateSongInfo(oldGroupName, oldSongTitle string, newGroupName, newSongTitle string) error {
	return s.repo.UpdateSongInfo(oldGroupName, oldSongTitle, newGroupName, newSongTitle)
}

func (s *Service) GetSongTextByVerses(groupName, songTitle string, offset, limit int) (string, error) {
	return s.repo.GetSongTextByVerses(groupName, songTitle, offset, limit)
}

func (s *Service) GetSongsInfo(groupName, songTitle string, offset, limit int) ([]models.SongInfo, error) {
	return s.repo.GetSongsInfo(groupName, songTitle, offset, limit)
}
