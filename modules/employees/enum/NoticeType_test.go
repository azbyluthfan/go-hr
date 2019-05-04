package enum

import "testing"

func TestNoticeType_String(t *testing.T) {
	if NoticeType(ADMIN).String() != SICK.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", NoticeType(SICK).String(), SICK.String())
	}
}

func TestNoticeType_Scan(t *testing.T) {

	value := "remote"
	arr := make([]uint8, len(value))
	for i := range arr {
		arr[i] = uint8(value[i])
	}

	nt := NoticeType(REMOTE)
	if err := nt.Scan(arr); err != nil {
		t.Errorf("Returned error when method should be successful")
	}
}

func TestNoticeType_Value(t *testing.T) {
	value, _ := NoticeType(SICK).Value()
	if value != SICK.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", value, SICK.String())
	}
}

func TestNoticeType_MarshalJSON(t *testing.T) {
	json, _ := NoticeType(REMOTE).MarshalJSON()
	if string(json) != "\"" + REMOTE.String() + "\"" {
		t.Errorf("Incorrect return value, got: %s, want: %s", string(json), REMOTE.String())
	}
}

func TestNoticeType_UnmarshalJSON(t *testing.T) {
	nt := NoticeType(VACATION)
	nt.UnmarshalJSON([]byte("vacation"))

	if nt.String() != VACATION.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", nt.String(), VACATION.String())
	}
}
