package enum

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type EmployeeRole int
const (
	ADMIN EmployeeRole = iota
	NORMAL
)

var employeeRoleToString = map[EmployeeRole]string{
	ADMIN: "admin",
	NORMAL: "normal",
}

var employeeRoleToID = map[string]EmployeeRole{
	"admin": ADMIN,
	"normal": NORMAL,
}

func (er EmployeeRole) String() string {
	return employeeRoleToString[er]
}

func (er *EmployeeRole) Scan(value interface{}) error {
	v := ""
	switch value.(type) {
	case []uint8:
		v = string(value.([]uint8))
	default:
		v = value.(string)
	}
	*er = EmployeeRole(employeeRoleToID[v])
	return nil
}

func (er EmployeeRole) Value() (driver.Value, error) {
	return er.String(), nil
}

func (er EmployeeRole) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(employeeRoleToString[er])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (er *EmployeeRole) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*er = employeeRoleToID[j]
	return nil
}