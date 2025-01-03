package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	
	_ "github.com/singl3focus/em_testtask/docs" // docs is generated by Swag CLI.
)

func (h *Handler) Router() http.Handler {
	public := mux.NewRouter()

	public.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
	)).Methods(http.MethodGet)

	public.HandleFunc("/healthy", h.Healthy).Methods(http.MethodGet)
	public.HandleFunc("/song/add", h.AddSong).Methods(http.MethodPost)
	public.HandleFunc("/song/remove", h.RemoveSong).Methods(http.MethodDelete)
	public.HandleFunc("/song/update", h.UpdateSongInfo).Methods(http.MethodPost)
	public.HandleFunc("/song/info/get", h.GetSongInfo).Methods(http.MethodGet)
	public.HandleFunc("/song/text/by-verses", h.GetSongTextByVerses).Methods(http.MethodGet)

	return public
}