package main

import (
	"errors"
	"fmt"
)

var errMissingUsername = errors.New("username is required")

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errMissingUsername
	}
	username := cmd.arguments[0]

	s.cfg.SetUser(username)

	fmt.Printf("user has been set to %q\n", username)

	return nil
}
