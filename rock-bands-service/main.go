package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	// Инициализация базы данных
	if err := InitDB(); err != nil {
		fmt.Printf("Ошибка при инициализации базы данных: %s\n", err.Error())
		return
	}
	defer db.Close()

	// Настройка роутера
	r := chi.NewRouter()

	// Регистрация маршрутов
	r.Get("/artists", getArtists)
	r.Post("/artists", postArtist)
	r.Put("/artist/{id}", updateArtist)
	r.Delete("/artist/{id}", deleteArtist)

	r.Get("/albums", getAlbums)
	r.Post("/albums", postAlbum)
	r.Put("/album/{id}", updateAlbum)
	r.Delete("/album/{id}", deleteAlbum)

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", r)
}
