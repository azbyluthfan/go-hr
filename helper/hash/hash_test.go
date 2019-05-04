package hash

import "testing"

func TestDecodeHashedPassword(t *testing.T) {
	encodedPassword := "ZDMyMjQ4ZWQwNzI3YTJlMmYzMGViNjEwNDliNmFhNzI1MWI5ODVhY2E2NjE1ZTE2MGNlNzI5MDBkMjllNGY0Ny52YXpvNW9xZXM5PQ=="
	hashedPassword, salt, _ := DecodeHashedPassword(encodedPassword)

	if hashedPassword != "d32248ed0727a2e2f30eb61049b6aa7251b985aca6615e160ce72900d29e4f47" {
		t.Errorf("Incorrect hashed password, got: %s, want: %s", hashedPassword, "d32248ed0727a2e2f30eb61049b6aa7251b985aca6615e160ce72900d29e4f47")
	}

	if salt != "vazo5oqes9" {
		t.Errorf("Incorrect salt, got: %s, want: %s", salt, "vazo5oqes9")
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
		t.Errorf("Bad base64 string error not captured")
	}
}

func TestHashPasswordWithSalt(t *testing.T) {
	hashedPassword := HashPasswordWithSalt("123456", "vazo5oqes9")
	if hashedPassword != "d32248ed0727a2e2f30eb61049b6aa7251b985aca6615e160ce72900d29e4f47" {
		t.Errorf("Incorrect hashed password, got: %s, want: %s", hashedPassword, "d32248ed0727a2e2f30eb61049b6aa7251b985aca6615e160ce72900d29e4f47")
	}
}
