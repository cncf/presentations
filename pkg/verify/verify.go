package verify

import (
	"fmt"
	"os"

	"net/http"
	"net/url"

	types "github.com/cncf/presentations/pkg/types"
	"gopkg.in/yaml.v3"
)

func Verify(f string) error {
	yamlFile, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("yamlFile.Get err   #%v ", err)
	}

	data := types.Presentations{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return fmt.Errorf("unmarshal error: %v", err)
	}

	errors := []error{}
	for key, entry := range data {
		if entry.Name == "" {
			errors = append(errors, fmt.Errorf("empty name detected for entry %d", key))
			continue
		}

		if entry.Description == "" {
			errors = append(errors, fmt.Errorf("empty description detected for '%s'", entry.Name))
		}
		if entry.Date.IsZero() {
			errors = append(errors, fmt.Errorf("empty date detected for entry '%s'", entry.Name))
		}

		if entry.Slides != "" {
			if _, err := url.Parse(entry.Slides); err != nil {
				errors = append(errors, fmt.Errorf("invalid slides URL for %s: %v", entry.Name, err))
			} else {
				if resp, err := http.Get(entry.Slides); err != nil || resp.StatusCode >= 400 {
					errors = append(errors, fmt.Errorf("broken slides URL for %s (%s): %v", entry.Name, entry.Slides, err))
				}
			}
		} else {
			errors = append(errors, fmt.Errorf("slides URL empty for %s", entry.Name))
		}

		if entry.Video != "" {
			if _, err := url.Parse(entry.Video); err != nil {
				errors = append(errors, fmt.Errorf("invalid video URL for %s: %v", entry.Name, err))
			} else {
				// This won't properly check YouTube links because YouTube doesn't 404 with a broken URL... :'(
				// TODO: Better validate YouTube URLs
				if resp, err := http.Get(entry.Video); err != nil || resp.StatusCode >= 400 {
					errors = append(errors, fmt.Errorf("broken video URL for %s (%s):  %v", entry.Name, entry.Video, err))
				}
			}
		}

		for _, repo := range entry.Repos {
			if resp, err := http.Get(repo); err != nil || resp.StatusCode >= 400 {
				errors = append(errors, fmt.Errorf("invalid repo URL: %s", repo))
			}
		}

		for _, presenter := range entry.Presenters {
			if presenter.Github != "" {
				githubURL := fmt.Sprintf("https://github.com/%s", presenter.Github)
				if resp, err := http.Get(githubURL); err != nil || resp.StatusCode >= 400 {
					errors = append(errors, fmt.Errorf("invalid GitHub user: %s", presenter.Github))
				}
			}
		}
	}
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		return fmt.Errorf("errors found: %d", len(errors))
	}

	return nil
}
