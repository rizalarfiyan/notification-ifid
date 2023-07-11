package utils

func GetStateWording(state int) string {
	switch state {
	case 1, 2:
		return "Bahaya"
	case 3, 4, 5:
		return "Bahaya"
	default:
		return "Aman"
	}
}
