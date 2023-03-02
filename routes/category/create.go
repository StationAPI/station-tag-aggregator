package category

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	neon "github.com/stationapi/station-tag-aggregator/db"
	"github.com/stationapi/station-tag-aggregator/routes"
	"gorm.io/gorm"
)

func Create(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	category := neon.Category{}

	err := routes.ProcessBody(r.Body, &category)

	if err != nil {
		http.Error(w, "there was an error processing the request body", http.StatusInternalServerError)

		return err
	}

	if category.Id != "" {
		http.Error(w, "the id was already defined in the request body", http.StatusBadRequest)

		return errors.New("the id was already defined in the request body")
	}

	category.Id = uuid.NewString()

	neon.CreateCategory(category, db)

	return nil
}
