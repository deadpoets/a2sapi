package steam

import (
	"strings"
	"testing"
)

func TestParseServerInfo(t *testing.T) {
	data := []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x49, 0x11, 0x71, 0x6C, 0x2E, 0x73, 0x79, 0x6E,
		0x63, 0x6F, 0x72, 0x65, 0x2E, 0x6F, 0x72, 0x67, 0x20, 0x2D, 0x20, 0x55,
		0x53, 0x20, 0x43, 0x45, 0x4E, 0x54, 0x52, 0x41, 0x4C, 0x20, 0x23, 0x31,
		0x00, 0x74, 0x68, 0x75, 0x6E, 0x64, 0x65, 0x72, 0x73, 0x74, 0x72, 0x75,
		0x63, 0x6B, 0x00, 0x62, 0x61, 0x73, 0x65, 0x71, 0x33, 0x00, 0x43, 0x6C,
		0x61, 0x6E, 0x20, 0x41, 0x72, 0x65, 0x6E, 0x61, 0x00, 0x00, 0x00, 0x02,
		0x10, 0x00, 0x64, 0x6C, 0x00, 0x01, 0x31, 0x30, 0x36, 0x33, 0x00, 0xB1,
		0x38, 0x6D, 0x02, 0xF8, 0xC1, 0x4D, 0x7B, 0x17, 0x40, 0x01, 0x63, 0x6C,
		0x61, 0x6E, 0x61, 0x72, 0x65, 0x6E, 0x61, 0x2C, 0x73, 0x79, 0x6E, 0x63,
		0x6F, 0x72, 0x65, 0x2C, 0x74, 0x65, 0x78, 0x61, 0x73, 0x00, 0x48, 0x4F,
		0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	sinfo, err := parseServerInfo(data)
	if err != nil {
		t.Fatalf("Unexpected error when parsing server info")
	}
	if !strings.EqualFold(sinfo.Name, "ql.syncore.org - US CENTRAL #1") {
		t.Fatalf("Expected server name: ql.syncore.org - US CENTRAL #1 got: %s",
			sinfo.Name)
	}
	if !strings.EqualFold(sinfo.Environment, "Linux") {
		t.Fatalf("Expected server environment: Linux got: %s",
			sinfo.Environment)
	}
	if sinfo.Players != 2 {
		t.Fatalf("Expected server to contain 2 players, got: %d", sinfo.Players)
	}
	if !strings.EqualFold(sinfo.Folder, "baseq3") {
		t.Fatalf("Expected server's game folder to be baseq3, got: %s", sinfo.Folder)
	}
}
