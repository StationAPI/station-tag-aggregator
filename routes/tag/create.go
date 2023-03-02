package tag

import (
	"errors"
	"net/http"

	neon "github.com/stationapi/station-tag-aggregator/db"
	"github.com/stationapi/station-tag-aggregator/routes"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	tag := neon.Tag{} 		

	err := routes.ProcessBody(r.Body, &tag)

	if err != nil {
		http.Error(w, "there was an error processing the body", http.StatusInternalServerError)

		return err
	}

	if tag.Id != "" {
		http.Error(w, "the id is already defined in the request", http.StatusBadRequest)

		return errors.New("the id is already defined in the request")
	}

	id := uuid.NewString() 

	tag.Id = id

	_, ok := neon.GetCategory(tag.CategoryId, db)

	if !ok {
		http.Error(w, "that category does not exist", http.StatusBadRequest)

		return errors.New("that category does not exist")
	}

	neon.CreateTag(tag, db)

	return nil
}
