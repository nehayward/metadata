package app

import (
	"errors"

	"github.com/nehayward/metadata/models"
)

var apps []models.App

// Save app data
func Save(app models.App) error {
	for _, a := range apps {
		if a.Title == app.Title && a.Version == app.Version {
			return errors.New("app with '" + app.Version + "' and '" + app.Title + "' already exist")
		}
	}
	apps = append(apps, app)
	return nil
}
