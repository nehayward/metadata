package models

type App struct {
	Title       string
	Version     string
	Maintainers []Maintainer
	Company     string
	Website     string
	Source      string
	License     string
	Description string
}
