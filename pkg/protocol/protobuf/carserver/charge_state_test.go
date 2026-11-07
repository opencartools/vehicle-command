package carserver

import (
	"encoding/hex"
	"testing"

	"google.golang.org/protobuf/proto"
)

// The four last-charge fields exactly as a Model 3 (firmware 2026.38.3.1) sent them over BLE.
const lastChargeWire = "bd0b3d0a8740c20b03455552cd0b295c8f3ef20b0608cbdcfbd006"

func TestChargeStateLastChargeSummary(t *testing.T) {
	raw, err := hex.DecodeString(lastChargeWire)
	if err != nil {
		t.Fatal(err)
	}
	var state ChargeState
	if err := proto.Unmarshal(raw, &state); err != nil {
		t.Fatal(err)
	}
	if got := state.GetLastChargeCost(); got != float32(4.22) {
		t.Errorf("cost = %v, want 4.22", got)
	}
	if got := state.GetLastChargeCurrency(); got != "EUR" {
		t.Errorf("currency = %q, want EUR", got)
	}
	if got := state.GetLastChargePricePerKwh(); got != float32(0.28) {
		t.Errorf("price per kWh = %v, want 0.28", got)
	}
	if got := state.GetLastChargeTimestamp().GetSeconds(); got != 1780411979 {
		t.Errorf("timestamp = %d, want 1780411979 (2026-06-02T14:52:59Z)", got)
	}
	if len(state.ProtoReflect().GetUnknown()) != 0 {
		t.Errorf("fields left unknown: %x", state.ProtoReflect().GetUnknown())
	}
	back, err := proto.Marshal(&state)
	if err != nil {
		t.Fatal(err)
	}
	if hex.EncodeToString(back) != lastChargeWire {
		t.Errorf("re-encoded %x, want %s", back, lastChargeWire)
	}
}
