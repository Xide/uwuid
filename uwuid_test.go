package uwuid_test

import (
	"reflect"
	"testing"
	"uwuid"
)

func TestNewRandom(t *testing.T) {
	_, err := uwuid.NewUwUID()
	if err != nil {
		t.Errorf("Couldn't generate a new UwUID: %v", err)
	}
}

func TestParseBytes(t *testing.T) {
	bts := []byte{
		0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
		0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	}
	if _, err := uwuid.FromBytes(bts); err != nil {
		t.Errorf("Couldn't parse UwUID from bytes: %v", err)
	}
}

func TestParseInvalidBytesSize(t *testing.T) {
	bts := []byte{
		0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	}
	if _, err := uwuid.FromBytes(bts); err == nil {
		t.Errorf("Expected failure when using 8 bytes instead of 16")
	}
}

func TestParseString(t *testing.T) {
	uwu := "UwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUwUUw"
	_, err := uwuid.FromString(uwu)
	if err != nil {
		t.Errorf("Couldn't parse UwUID from string: %v", err)
	}
}

func TestParseInvalidStringSize(t *testing.T) {
	uwu := ""
	if _, err := uwuid.FromString(uwu); err == nil {
		t.Errorf("Expected failure when parsing empty string")
	}
}

func TestIdempotentParse(t *testing.T) {
	uwu, _ := uwuid.NewUwUID()

	uwuRepr := uwu.String()

	parsedUwU, err := uwuid.FromString(uwuRepr)
	if err != nil {
		t.Errorf("failed to parse serialized UwUID")
	}

	if !reflect.DeepEqual(*uwu, *parsedUwU) {
		t.Errorf("UwUID's are different : %v != %v", uwu, parsedUwU)
	}
}
