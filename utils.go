package main


func getPosition(slice []string, val string) (position int) {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
