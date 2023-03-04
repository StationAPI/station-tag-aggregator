package category

import (
	"encoding/json"
	"net/http"

	neon "github.com/stationapi/station-tag-aggregator/db"
	"gorm.io/gorm"
)

func List(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	categories := neon.ListCategories(db)

	stringified, err := json.Marshal(categories)

	if err != nil {
		http.Error(w, "there was an error listing the categories", http.StatusInternalServerError)

		return err
	}

	w.WriteHeader(200)
	w.Write(stringified)

	return nil
}
