package database

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"time"
)

var nameRE = regexp.MustCompile(`^[0-9a-z_\-]+$`)

func genMigration(name string) (string, error) {
	const timeFormat = "20060102150405"

	if name == "" {
		return "", fmt.Errorf("migrate: migration name can't be empty")
	}
	if !nameRE.MatchString(name) {
		return "", fmt.Errorf("migrate: invalid migration name: %q", name)
	}

	version := time.Now().UTC().Format(timeFormat)
	return fmt.Sprintf("%s_%s", version, name), nil
}

func CreateMigration(destination string) error {
	template, err := genMigration("migration")
	if err != nil {
		return err
	}
	m := []string{"up", "down"}
	for _, x := range m {
		filename := fmt.Sprintf("%s.%s.sql", template, x)
		target := path.Join(destination, filename)
		if err := ioutil.WriteFile(target, nil, 0o644); err != nil {
			return err
		}
		if err := ioutil.WriteFile(path.Join(destination, filename), nil, 0o644); err != nil {
			return err
		}
	}

	return nil
}
