package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"regexp"
)

// GET /artists
func getArtists(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, born, genre FROM artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var artists []Artist
	for rows.Next() {
		var artist Artist
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.Born, &artist.Genre); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		artists = append(artists, artist)
	}

	resp, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

// POST /artists
func postArtist(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
		http.Error(w, "Неверные входные данные", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO artists (id, name, born, genre) VALUES (?, ?, ?, ?)",
		artist.ID, artist.Name, artist.Born, artist.Genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// PUT /atrist
func updateArtist(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var artist Artist
	if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
		http.Error(w, "Неверные входные данные", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE artists SET name = ?, born = ?, genre = ? WHERE id = ?",
		artist.Name, artist.Born, artist.Genre, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DELETE /artist
func deleteArtist(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := db.Exec("DELETE FROM artists WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Ошибка при проверке результата удаления", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Артист с указанным ID не найден", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GET /albums
func getAlbums(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, year, artist_id FROM albums")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Year, &album.ArtistID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		albums = append(albums, album)
	}

	resp, err := json.Marshal(albums)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// POST /albums
func postAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация данных
	if !isValidYear(album.Year) {
		http.Error(w, "Неправильный формат года", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO albums (id, title, year, artist_id) VALUES (?, ?, ?, ?)",
		album.ID, album.Title, album.Year, album.ArtistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
func isValidYear(year string) bool {
	// Проверка года (4 цифры)
	re := regexp.MustCompile(`^\d{4}$`)
	return re.MatchString(year)
}

// DELETE /album/{id}
func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Получаем ID альбома из маршрута

	// Выполняем запрос на удаление
	result, err := db.Exec("DELETE FROM albums WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Ошибка при удалении альбома", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Ошибка при проверке результата удаления", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Альбом с указанным ID не найден", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Альбом удалён"))
}

// PUT /album/{id}
func updateAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	id := chi.URLParam(r, "id") // Получаем ID альбома из маршрута

	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Обновляем данные в базе
	_, err := db.Exec("UPDATE albums SET title = ?, year = ?, artist_id = ?  WHERE id = ?",
		album.Title, album.ArtistID, album.Year, id)
	if err != nil {
		http.Error(w, "Ошибка при обновлении альбома", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Альбом обновлён"))
}
