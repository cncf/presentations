package types

import (
	"time"
)

type Presentations []struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Video       string    `yaml:"video"`
	Date        time.Time `yaml:"date"`
	Slides      string    `yaml:"slides"`
	Language    string    `yaml:"language"`
	License     string    `yaml:"license"`
	Presenters  []struct {
		Name   string `yaml:"name"`
		Github string `yaml:"github"`
	}
	Event struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	}
	Projects []string `yaml:"projects"`
	Tags     []string `yaml:"tags"`
}
