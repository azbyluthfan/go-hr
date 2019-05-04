package hash

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

func DecodeHashedPassword(hashed string) (hashedPassword string, salt string, err error) {
	decoded, err := b64.StdEncoding.DecodeString(hashed)
	if err != nil {
		return "", "", err
	}

	str := strings.Split(string(decoded), ".")
	if len(str) != 2 {
		return "", "", errors.New("Bad hashed password")
	}

	// remove trailing =
	str[1] = str[1][:len(str[1])-1]

	return str[0], str[1], nil
}

func HashPasswordWithSalt(password, salt string) string {
	hashed := sha256.New()
	hashed.Write([]byte(password + "." + salt))
	return hex.EncodeToString(hashed.Sum(nil))
}
