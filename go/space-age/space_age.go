package space

// import "time"

type Planet string

// Age returns
func Age(age float64, p Planet) float64 {
	years := age / (3600 * 24 * 365.25)
	switch p {
	case "Earth":
		age = years
	case "Mercury":
		age = years / .2408467
	case "Venus":
		age = years / .61519726
	case "Mars":
		age = years / 1.8808158
	case "Jupiter":
		age = years / 11.862615
	case "Saturn":
		age = years / 29.447498
	case "Uranus":
		age = years / 84.016846
	case "Neptune":
		age = years / 164.79132
	}
	return age
}
