// Package triage implementa la clasificación determinista de urgencia
// del protocolo Nurse Canning 24 (Códigos Celeste/Amarillo/Rojo).
// Regla pura: sin dependencia de modelos externos => CI-friendly.
package triage

// Level representa el código de triage validado en el borde.
type Level string

// Códigos de la Matriz Determinista de Triage.
const (
	Celeste  Level = "CELESTE"
	Amarillo Level = "AMARILLO"
	Rojo     Level = "ROJO"
)

// Classify asigna el nivel de triage según signos vitales simplificados.
// Umbral Rojo: riesgo inminente => derivación humana forzada (Override).
func Classify(systolicBP, heartRate, spo2 int) Level {
	switch {
	case spo2 < 90 || systolicBP < 90 || heartRate > 130:
		return Rojo
	case spo2 < 95 || systolicBP < 100 || heartRate > 100:
		return Amarillo
	default:
		return Celeste
	}
}
