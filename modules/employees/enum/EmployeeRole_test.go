package enum

import (
	"testing"
)

func TestEmployeeRole_String(t *testing.T) {
	if EmployeeRole(ADMIN).String() != ADMIN.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", EmployeeRole(ADMIN).String(), ADMIN.String())
	}
}

func TestEmployeeRole_Scan(t *testing.T) {

	value := "admin"
	arr := make([]uint8, len(value))
	for i := range arr {
		arr[i] = uint8(value[i])
	}

	role := EmployeeRole(ADMIN)
	if err := role.Scan(arr); err != nil {
		t.Errorf("Returned error when method should be successful")
	}
}

func TestEmployeeRole_Value(t *testing.T) {
	value, _ := EmployeeRole(ADMIN).Value()
	if value != ADMIN.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", value, ADMIN.String())
	}
}

func TestEmployeeRole_MarshalJSON(t *testing.T) {
	json, _ := EmployeeRole(ADMIN).MarshalJSON()
	if string(json) != "\"" + ADMIN.String() + "\"" {
		t.Errorf("Incorrect return value, got: %s, want: %s", string(json), ADMIN.String())
	}
}

func TestEmployeeRole_UnmarshalJSON(t *testing.T) {
	role := EmployeeRole(NORMAL)
	role.UnmarshalJSON([]byte("normal"))

	if role.String() != NORMAL.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", role.String(), NORMAL.String())
	}
}