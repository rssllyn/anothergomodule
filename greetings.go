package anothergomodule

import (
	"errors"
	"fmt"

	"github.com/rssllyn/thirdlevel"
)

func Hello() (string, error) {
	// Return a greeting that embeds the name in a message.
	name := thirdlevel.Name()
	if len(name) == 0 {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
