package util

import "testing"

func TestPercentageChange(t *testing.T) {
	old, new := 20, 60
	pct := PercentageChange(old, new)

	if int(pct) != 200.0 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}

func TestPercentageChangeFloat(t *testing.T) {
	old, new := 20.0, 60.0
	pct := PercentageChangeFloat(old, new)

	if int(pct) != 200.0 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}

func TestPercentageChangeString(t *testing.T) {
	old, new := "20.0", "60.0"
	pct, err := PercentageChangeString(old, new)
	if err != nil {
		t.Fatalf("percentage change error: %v", err)
	}

	if int(pct) != 200.0 {
		t.Fatalf("%f is wrong percent!", pct)
	}
}
