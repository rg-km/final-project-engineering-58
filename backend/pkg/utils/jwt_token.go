package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

/**
 *	return (user_id, roles, error)
 */
func DecodeJwtToken(req *http.Request) (string, []string, error)  {
	header := req.Header.Get("Authorization")

	if header == "" {
		return "", nil, errors.New("Header authorization not found")
	}

	hash := strings.ReplaceAll(header, "Bearer ", "")

	token := strings.Split(hash, ".")

	if len(token) < 3 {
		return "", nil, errors.New("Token not failed")
	}

	decoded, errDecoded := base64.RawURLEncoding.DecodeString(token[1])
	if errDecoded != nil {
		return "", nil, errors.New("Error encoding token")
	}

	dataJwt := make(map[string]interface{})
	roles := make([]string, 0)

	json.Unmarshal(decoded, &dataJwt)

	sub := fmt.Sprintf("%v", dataJwt["sub"])

	roleOrg := reflect.ValueOf(dataJwt["role"])

	for i := 0; i < roleOrg.Len(); i++ {
		role := fmt.Sprintf("%v", roleOrg.Index(i))
		roles = append(roles, role)
	}

	return sub, roles, nil
}