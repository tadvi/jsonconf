/*
Package for saving and loading JSON from files.
*/
package jsonconf

import (
	"encoding/json"
	"os"
)

// Load configuration.
func Load(file string, conf interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		return err
	}
	return nil
}

// Save configuration.
func Save(file string, conf interface{}) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	en := json.NewEncoder(f)
	en.SetIndent("", "    ")
	if err := en.Encode(conf); err != nil {
		return err
	}
	return nil
}
