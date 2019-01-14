package app

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/nehayward/metadata/models"
	yaml "gopkg.in/yaml.v2"
)

// IsValid Checks that data is a valid yaml
func IsValid(data []byte) (models.App, error) {
	var app models.App

	fmt.Println(string(data))
	if err := yaml.UnmarshalStrict(data, &app); err != nil {
		fmt.Println(err)
		return app, err
	}

	if isEmpty(app.Title) {
		return app, errors.New("application title field required")
	}

	if isEmpty(app.Version) {
		return app, errors.New("application version field required")
	}

	if err := isValidMaintainer(app.Maintainers); err != nil {
		return app, err
	}

	if isEmpty(app.Company) {
		return app, errors.New("application company field required")
	}

	if err := isValidURL(app.Website); err != nil {
		return app, err
	}

	if err := isValidURL(app.Source); err != nil {
		return app, err
	}

	if isEmpty(app.License) {
		return app, errors.New("application license field required")
	}

	if isEmpty(app.Description) {
		return app, errors.New("application description field required")
	}

	return app, nil
}

func isValidMaintainer(maintainers []models.Maintainer) error {
	for _, m := range maintainers {
		if isEmpty(m.Name) {
			return errors.New("application name field required for maintaner")
		}

		if err := isValidEmail(m.Email); err != nil {
			return err
		}
	}

	return nil
}

func isValidEmail(email string) error {
	if isEmpty(email) {
		return errors.New("application email field required for maintaner")
	}

	validEmailRe := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !validEmailRe.MatchString(email) {
		return errors.New("not valid email: " + email)
	}

	return nil
}

func isValidURL(website string) error {
	if _, err := url.ParseRequestURI(website); err != nil {
		return errors.New("not valid url: " + website)
	}

	return nil
}

func isEmpty(field string) bool {
	if field == "" {
		return true
	}

	return false
}
