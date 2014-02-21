package job

import (
	"testing"

	"github.com/coreos/fleet/unit"
)

func TestNewJobPayloadBadType(t *testing.T) {
	j := NewJobPayload("foo.unknown", *unit.NewSystemdUnitFile("echo"))
	_, err := j.Type()

	if err == nil {
		t.Fatal("Expected non-nil error")
	}
}

func TestNewJobPayload(t *testing.T) {
	payload := NewJobPayload("echo.service", *unit.NewSystemdUnitFile("Echo"))

	if payload.Name != "echo.service" {
		t.Errorf("Payload has unexpected name '%s'", payload.Name)
	}

	if pt, _ := payload.Type(); pt != "systemd-service" {
		t.Errorf("Payload has unexpected Type '%s'", pt)
	}
}

func TestJobPayloadServiceDefaultPeers(t *testing.T) {
	unitFile := unit.NewSystemdUnitFile("")
	payload := NewJobPayload("echo.service", *unitFile)
	peers := payload.Peers()

	if len(peers) != 0 {
		t.Fatalf("Unexpected number of peers %d, expected 0", len(peers))
	}
}

func TestJobPayloadSocketDefaultPeers(t *testing.T) {
	unitFile := unit.NewSystemdUnitFile("")
	payload := NewJobPayload("echo.socket", *unitFile)
	peers := payload.Peers()

	if len(peers) != 1 {
		t.Fatalf("Unexpected number of peers %d, expected 1", len(peers))
	}

	if peers[0] != "echo.service" {
		t.Fatalf("Unexpected peers: %v", peers)
	}
}
