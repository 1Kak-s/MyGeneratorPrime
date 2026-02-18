package pirate

import (
	"errors"
	"strings"
)

// Struct Pirate
type Pirate struct {
	Name  string
	Prime string
	Img   string
}

// Interface pour sauve un pirate
type Saver interface {
	Save(p Pirate) error
}

// create a new pirate and validate the input
func New(name, prime, img string) (*Pirate, error) {
	if name == "" {
		return nil, errors.New("le nom du pirate ne peut pas être vide")
	}
	if prime == "" {
		return nil, errors.New("la prime ne peut pas être vide")
	}

	name = strings.ToUpper(name)

	return &Pirate{
		Name:  name,
		Prime: prime,
		Img:   img,
	}, nil
}
