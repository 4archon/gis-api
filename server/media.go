package server

import (
	"encoding/json"
	"io"
	"log"
	"map/business"
	"net/http"
	"strconv"
	"errors"
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

	// data.Medias = data.Medias[0:3]
	data.Medias = recentMedia(data.Medias)

	resutl, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	response.Header().Set("Content-Type", "applicaton/json")
	response.WriteHeader(http.StatusOK)
	response.Write(resutl)
}


func recentMediaByName(data []business.Media, name string) (business.Media, error) {
	for _, i := range data {
		if i.MediaName == name {
			return i, nil
		}
	}

	return business.Media{}, errors.New("nothing has been found")
}

func recentMedia(data []business.Media) []business.Media {
	var medias []business.Media

	if media, err := recentMediaByName(data, "Фото слева"); err == nil {
		medias = append(medias, media)
	}
	if media, err := recentMediaByName(data, "Фото спереди"); err == nil {
		medias = append(medias, media)
	}
	if media, err := recentMediaByName(data, "Фото справа"); err == nil {
		medias = append(medias, media)
	}
	if media, err := recentMediaByName(data, "Видео"); err == nil {
		medias = append(medias, media)
	}
	
	return medias
}