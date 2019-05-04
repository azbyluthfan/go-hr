package enum

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type NoticeVisibility int
const (
	PUBLIC NoticeVisibility = iota
	PRIVATE
)


var noticeVisibilityToString = map[NoticeVisibility]string{
	PUBLIC: "public",
	PRIVATE: "private",
}

var noticeVisibilityToID = map[string]NoticeVisibility{
	"public": PUBLIC,
	"private": PRIVATE,
}

func (nv NoticeVisibility) String() string {
	return noticeVisibilityToString[nv]
}

func (nv *NoticeVisibility) Scan(value interface{}) error {
	v := ""
	switch value.(type) {
	case []uint8:
		v = string(value.([]uint8))
	default:
		v = value.(string)
	}
	*nv = NoticeVisibility(noticeVisibilityToID[v])
	return nil
}

func (nv NoticeVisibility) Value() (driver.Value, error) {
	return nv.String(), nil
}

func (nv NoticeVisibility) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(noticeVisibilityToString[nv])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (nv *NoticeVisibility) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*nv = noticeVisibilityToID[j]
	return nil
}