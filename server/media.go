package server

import (
	"encoding/json"
	"io"
	"log"
	"map/business"
	"net/http"
	"strconv"
	"errors"
	"slices"
)


func (s Server) postPointRecentMedia(response http.ResponseWriter, req *http.Request) {
	_, _, err := s.checkUser(response, req)
	if err != nil {
		return
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}

	id, err := strconv.Atoi(string(body))
	if err != nil {
		log.Println(err)
		return
	}

	data, err := s.DB.GetPointMedia(id)
	if err != nil {
		return
	}

	data.Medias = recentMedia(data.Medias)

	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}


func recentMediaByName(data []business.Media, name string) (business.Media, error) {
	switch (name) {
	case "left":
		targetValue := []string{"Фото слева(старое место)", "Фото слева(новое место)",
		"Фото слева", "Фото слева после покраски", "Фото разметки слева"}
		for _, i := range data {
			if slices.Contains(targetValue, i.MediaName) {
				return i, nil
			}
		}
	case "front":
		targetValue := []string{"Фото спереди(старое место)", "Фото спереди(новое место)",
		"Фото спереди", "Фото спереди после покраски", "Фото разметки спереди"}
		for _, i := range data {
			if slices.Contains(targetValue, i.MediaName) {
				return i, nil
			}
		}
	case "right":
		targetValue := []string{"Фото справа(старое место)", "Фото справа(новое место)",
		"Фото справа", "Фото справа после покраски", "Фото разметки справа"}
		for _, i := range data {
			if slices.Contains(targetValue, i.MediaName) {
				return i, nil
			}
		}
	case "video":
		targetValue := []string{"Видео", "Видео после покраски", "Видео разметки"}
		for _, i := range data {
			if slices.Contains(targetValue, i.MediaName) {
				return i, nil
			}
		}
	}

	return business.Media{}, errors.New("nothing has been found")
}

func recentMedia(data []business.Media) []business.Media {
	var medias []business.Media

	if media, err := recentMediaByName(data, "left"); err == nil {
		medias = append(medias, media)
	}
	if media, err := recentMediaByName(data, "front"); err == nil {
		medias = append(medias, media)
	}
	if media, err := recentMediaByName(data, "right"); err == nil {
		medias = append(medias, media)
	}
	if media, err := recentMediaByName(data, "video"); err == nil {
		medias = append(medias, media)
	}
	
	return medias
}