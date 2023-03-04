package tag

import (
	"encoding/json"
	"errors"
	"net/http"

	neon "github.com/stationapi/station-tag-aggregator/db"
	"gorm.io/gorm"
)

func Get(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	id := r.URL.Query().Get("id")	

	if id == "" {
		http.Error(w, "there is no id in the query parameters", http.StatusBadRequest)
		
		return errors.New("there is no id in the query parameters")
	} 

	tag, ok := neon.GetTag(id, db)

	if !ok {
		http.Error(w, "that tag was not found", http.StatusNotFound)

		return errors.New("that tag was not found")
	}

	stringified, err := json.Marshal(tag)

	if err != nil {
		http.Error(w, "there was an error fetching the tag", http.StatusInternalServerError)

		return err
	}

	w.WriteHeader(200)
	w.Write(stringified)

	return nil
}
