package triage

import "testing"

func TestClassifyRojo(t *testing.T) {
	if got := Classify(80, 140, 88); got != Rojo {
		t.Fatalf("esperaba ROJO, obtuve %s", got)
	}
}

func TestClassifyAmarillo(t *testing.T) {
	if got := Classify(95, 110, 93); got != Amarillo {
		t.Fatalf("esperaba AMARILLO, obtuve %s", got)
	}
}

func TestClassifyCeleste(t *testing.T) {
	if got := Classify(120, 80, 98); got != Celeste {
		t.Fatalf("esperaba CELESTE, obtuve %s", got)
	}
}
