package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jabutech/simple-blog/util"
)

func EncodedToken(strToken string) (string, error) {
	// Load config
	config, err := util.LoadConfig("../") // "." as location file app.env in root folder
	FatalError("error load config: ", err)

	// If there is, create new variable with empty string value
	encodedToken := ""
	// Split authHeader with white space
	arrayToken := strings.Split(strToken, " ")
	// If length arrayToken is same the 2
	if len(arrayToken) == 2 {
		// Get arrayToken with index 1 / only token jwt
		encodedToken = arrayToken[1]
	}

	// Validation token
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(config.SecretKey), nil
	})

	// If error
	if err != nil {
		return "", err
	}

	// Get payload token
	claim, ok := token.Claims.(jwt.MapClaims)
	// If not `ok` and token invalid
	if !ok || !token.Valid {
		return "", errors.New("failed claim token")
	}

	// Get payload `user id` and convert to type `float64` and type `int`
	strUserId := fmt.Sprint(claim["user_id"])

	// Return user id string
	return strUserId, nil
}
