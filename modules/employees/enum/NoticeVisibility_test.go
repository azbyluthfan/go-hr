package enum

import "testing"

func TestNoticeVisibility_String(t *testing.T) {
	if NoticeVisibility(PUBLIC).String() != PUBLIC.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", NoticeVisibility(PUBLIC).String(), PUBLIC.String())
	}
}

func TestNoticeVisibility_Scan(t *testing.T) {

	value := "public"
	arr := make([]uint8, len(value))
	for i := range arr {
		arr[i] = uint8(value[i])
	}

	nv := NoticeVisibility(PUBLIC)
	if err := nv.Scan(arr); err != nil {
		t.Errorf("Returned error when method should be successful")
	}
}

func TestNoticeVisibility_Value(t *testing.T) {
	value, _ := NoticeVisibility(PRIVATE).Value()
	if value != PRIVATE.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", value, PRIVATE.String())
	}
}

func TestNoticeVisibility_MarshalJSON(t *testing.T) {
	json, _ := NoticeVisibility(PUBLIC).MarshalJSON()
	if string(json) != "\"" + PUBLIC.String() + "\"" {
		t.Errorf("Incorrect return value, got: %s, want: %s", string(json), PUBLIC.String())
	}
}

func TestNoticeVisibility_UnmarshalJSON(t *testing.T) {
	nv := NoticeVisibility(PRIVATE)
	nv.UnmarshalJSON([]byte("private"))

	if nv.String() != PRIVATE.String() {
		t.Errorf("Incorrect return value, got: %s, want: %s", nv.String(), PRIVATE.String())
	}
}

