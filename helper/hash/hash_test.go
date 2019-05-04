package hash

import "testing"

var (
	salt = "vazo5oqes9"
	password = "123456"
	encryptedPasswordWithSalt = "d32248ed0727a2e2f30eb61049b6aa7251b985aca6615e160ce72900d29e4f47"
)


func TestDecodeHashedPassword(t *testing.T) {
	encodedPassword := "ZDMyMjQ4ZWQwNzI3YTJlMmYzMGViNjEwNDliNmFhNzI1MWI5ODVhY2E2NjE1ZTE2MGNlNzI5MDBkMjllNGY0Ny52YXpvNW9xZXM5PQ=="
	hashedPassword, decodedSalt, _ := DecodeHashedPassword(encodedPassword)

	if hashedPassword != encryptedPasswordWithSalt {
		t.Errorf("Incorrect hashed password, got: %s, want: %s", hashedPassword, encryptedPasswordWithSalt)
	}

	if decodedSalt != salt {
		t.Errorf("Incorrect salt, got: %s, want: %s", decodedSalt, salt)
	}
}

func TestDecodeHashedPasswordWithBadString(t *testing.T) {
	encodedPassword := "ZDMyMjQ4ZWQwNzI3YTJlMmYzMGViNjEwNDliNmFhNzI1MWI5ODVhY2E2NjE1ZTE2MGNlNzI5MDBkMjllNGY0N3Zhem81b3Flczk9"
	_, _, err := DecodeHashedPassword(encodedPassword)
	if err.Error() != "Bad hashed password" {
		t.Errorf("Incorrect error, got: %s, want: %s", err.Error(), "Bad hashed password")
	}

	encodedPassword = "abc"
	_, _, err = DecodeHashedPassword(encodedPassword)
	if err == nil {
		t.Errorf("Bad base64 string error is not captured")
	}
}

func TestHashPasswordWithSalt(t *testing.T) {
	hashedPassword := HashPasswordWithSalt(password, salt)
	if hashedPassword != encryptedPasswordWithSalt {
		t.Errorf("Incorrect hashed password, got: %s, want: %s", hashedPassword, encryptedPasswordWithSalt)
	}
}
