package tag

import (
	"encoding/json"
	"net/http"

	neon "github.com/stationapi/station-tag-aggregator/db"
	"gorm.io/gorm"
)

func List(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	tags := neon.ListTags(db)

	stringified, err := json.Marshal(tags)

	if err != nil {
		http.Error(w, "there was an error listing the tags", http.StatusInternalServerError)

		return err
	}

	w.WriteHeader(200)
	w.Write(stringified)

	return nil
}
