package enum

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type NoticeType int
const (
	SICK NoticeType = iota
	REMOTE
	VACATION
)

var noticeTypeToString = map[NoticeType]string{
	SICK: "sick",
	REMOTE: "remote",
	VACATION: "vacation",
}

var noticeTypeToID = map[string]NoticeType{
	"sick": SICK,
	"remote": REMOTE,
	"vacation": VACATION,
}

func (nt NoticeType) String() string {
	return noticeTypeToString[nt]
}

func (nt *NoticeType) Scan(value interface{}) error {
	v := ""
	switch value.(type) {
	case []uint8:
		v = string(value.([]uint8))
	default:
		v = value.(string)
	}
	*nt = NoticeType(noticeTypeToID[v])
	return nil
}

func (nt NoticeType) Value() (driver.Value, error) {
	return nt.String(), nil
}

func (nt NoticeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(noticeTypeToString[nt])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (nt *NoticeType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*nt = noticeTypeToID[j]
	return nil
}