package utils

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/goofinator/usersHttp/internal/model"
)

// DecodeUser decodes user from incoming json
func DecodeUser(r io.Reader) (*model.User, error) {
	decoder := json.NewDecoder(r)
	var user model.User

	if err := decoder.Decode(&user); err != nil {
		return nil, fmt.Errorf("error on json.Decode: %s", err)
	}
	return &user, nil
}

// EncodeUsers encode users to json
func EncodeUsers(w io.Writer, users []*model.User) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(users); err != nil {
		return fmt.Errorf("error on jsonEncode: %s", err)
	}
	return nil
}
