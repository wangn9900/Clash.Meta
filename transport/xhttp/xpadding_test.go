package xhttp

import "testing"

func TestGetNormalizedXPaddingBytesDefault(t *testing.T) {
	r, err := (&Config{}).GetNormalizedXPaddingBytes()
	if err != nil {
		t.Fatal(err)
	}
	if r.Min != 100 || r.Max != 1000 {
		t.Fatalf("unexpected default xpadding range: got %d-%d, want 100-1000", r.Min, r.Max)
	}
}
