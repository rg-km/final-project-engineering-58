package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

/**
 *	return (user_id, roles, error)
 */
func DecodeJwtToken(req *http.Request) (string, string, error)  {
	header := req.Header.Get("Authorization")

	if header == "" {
		return "", "", errors.New("Header authorization not found")
	}

	hash := strings.ReplaceAll(header, "Bearer ", "")

	token := strings.Split(hash, ".")

	if len(token) < 3 {
		return "", "", errors.New("Token not failed")
	}

	decoded, errDecoded := base64.RawURLEncoding.DecodeString(token[1])
	if errDecoded != nil {
		return "", "", errors.New("Error encoding token")
	}

	dataJwt := make(map[string]interface{})

	json.Unmarshal(decoded, &dataJwt)

	sub := fmt.Sprintf("%v", dataJwt["sub"])

	role := fmt.Sprintf("%v", dataJwt["role"])

	return sub, role, nil
}