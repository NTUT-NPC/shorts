package shorts

import (
	"errors"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type RedirectsMap struct {
	Temporary map[string]string
	Permanent map[string]string
}

var Redirects RedirectsMap

const RedirectsFile string = "config/redirects.toml"

func ReadRedirects() {
	file, err := os.ReadFile(RedirectsFile)
	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal([]byte(file), &Redirects)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d temporary and %d permanent redirects", len(Redirects.Temporary), len(Redirects.Permanent))
}

func WriteRedirects() error {
	file, err := os.Create(RedirectsFile)
	if err != nil {
		return errors.New("unable to create redirects.toml")
	}
	defer file.Close()

	err = toml.NewEncoder(file).Encode(Redirects)
	if err != nil {
		return errors.New("unable to encode redirects.toml")
	}

	return nil
}
