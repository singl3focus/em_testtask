package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/singl3focus/em_testtask/internal/models"
)

type Service interface {
	AddSong(song models.Song) error
	RemoveSong(groupName, songTitle string) error
	UpdateSongInfo(oldGroupName, oldSongTitle string, newGroupName, newSongTitle string) error
	GetSongTextByVerses(groupName, songTitle string, offset, limit int) (string, error)
	GetSongsInfo(groupName, songTitle string, offset, limit int) ([]models.SongInfo, error)
}

type Handler struct {
	logger *slog.Logger
	service Service
}

func NewHandler(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
		service: service,
	}
}

// Healthy godoc
// @Summary Check service healthy or not
// @Success 200
func (h *Handler) Healthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// AddSong godoc
// @Summary Add song
// @Tags song
// @Accept json
// @Produce json
// @Param song body models.Song true "Song"
// @Success 200 {object} Result "Song was succesfully added"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /song/add [post]
func (h *Handler) AddSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		h.NewErrorResponse(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	h.logger.Debug("[AddSong]", "(input)", song)

	err := h.service.AddSong(song)
	if err != nil {
		h.NewErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	h.NewTextResponse(w, "Song was succesfully added")
}

// RemoveSong godoc
// @Summary Remove song by groupName, songTitle
// @Tags song
// @Produce json
// @Param groupName query string true "Group Name"
// @Param songTitle query string true "Song Title"
// @Success 200 {object} Result "Song was successfully removed"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /song/remove [delete]
func (h *Handler) RemoveSong(w http.ResponseWriter, r *http.Request) {
	groupName := r.URL.Query().Get("groupName")
	songTitle := r.URL.Query().Get("songTitle")
	if groupName == "" || songTitle == ""{
		h.NewErrorResponse(w, http.StatusBadRequest, "Bad request", "")
		return
	}

	h.logger.Debug("[RemoveSong]", "(params)",
			struct{
				groupName string
				songTitle string
			}{
				groupName: groupName,
				songTitle: songTitle})

	err := h.service.RemoveSong(groupName, songTitle)
	if err != nil {
		h.NewErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	h.NewTextResponse(w, "Song was successfully removed")
}

type UpdateSongInfoRequest struct {
	OldGroupName string `json:"oldGroupName"`
	OldSongTitle string `json:"oldSongTitle"`
	NewGroupName string `json:"newGroupName"`
	NewSongTitle string `json:"newSongTitle"`
}

// UpdateSongInfo godoc
// @Summary Update song by groupName, songTitle
// @Tags song
// @Accept json
// @Produce json
// @Param data body UpdateSongInfoRequest true "Old and new song info"
// @Success 200 {object} Result "Song's info was successfully update"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /song/update [post]
func (h *Handler) UpdateSongInfo(w http.ResponseWriter, r *http.Request) {
	var data UpdateSongInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.NewErrorResponse(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	h.logger.Debug("[RemoveSong]", "(input)", data)

	err := h.service.UpdateSongInfo(data.OldGroupName, data.OldSongTitle, data.NewGroupName, data.NewSongTitle)
	if err != nil {
		h.NewErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	h.NewTextResponse(w, "Song's info was successfully update")
}

// GetSongTextByVerses godoc
// @Summary Get song's text by groupName,songTitle with verses pagination
// @Tags song
// @Produce json
// @Param groupName query string true "Group Name"
// @Param songTitle query string true "Song Title"
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {object} Result "Song text"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /song/text/by-verses [get]
func (h *Handler) GetSongTextByVerses(w http.ResponseWriter, r *http.Request) {
	groupName := r.URL.Query().Get("groupName")
	songTitle := r.URL.Query().Get("songTitle")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	if (groupName == "") || (songTitle == "") || (limitStr == "")|| (offsetStr == "") {
		h.NewErrorResponse(w, http.StatusBadRequest, "Bad request", "")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.NewErrorResponse(w, http.StatusBadRequest, "'limit' is not integer", err.Error())
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		h.NewErrorResponse(w, http.StatusBadRequest, "'offset' is not integer", err.Error())
		return
	}

	h.logger.Debug("[GetSongTextByVerses]", "(params)",
			struct{
				groupName string
				songTitle string
				limit int
				offset int
			}{
				groupName: groupName,
				songTitle: songTitle,
				limit: limit,
				offset: offset,})


	text, err := h.service.GetSongTextByVerses(groupName, songTitle, offset, limit)
	if err != nil {
		h.NewErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	h.NewTextResponse(w, text)
}

// GetSongTextByVerses godoc
// @Summary Get songs info by matched params
// @Tags song
// @Produce json
// @Param groupName query string true "Group Name"
// @Param songTitle query string true "Song Title"
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Success 200 {array} models.SongInfo "Songs info"
// @Failure 400 {object} Error "Bad request"
// @Failure 500 {object} Error "Internal server error"
// @Router /song/info/get [get]
func (h *Handler) GetSongInfo(w http.ResponseWriter, r *http.Request) {
	groupName := r.URL.Query().Get("groupName")
	songTitle := r.URL.Query().Get("songTitle")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	if (groupName == "") || (songTitle == "") || (limitStr == "")|| (offsetStr == "") {
		h.NewErrorResponse(w, http.StatusBadRequest, "Bad request", "")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.NewErrorResponse(w, http.StatusBadRequest, "'limit' is not integer", err.Error())
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		h.NewErrorResponse(w, http.StatusBadRequest, "'offset' is not integer", err.Error())
		return
	}

	h.logger.Debug("[GetSongInfo]", "(params)",
			struct{
				groupName string
				songTitle string
				limit int
				offset int
			}{
				groupName: groupName,
				songTitle: songTitle,
				limit: limit,
				offset: offset,})


	models, err := h.service.GetSongsInfo(groupName, songTitle, offset, limit)
	if err != nil {
		h.NewErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(models); err != nil {
		h.NewErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}
}