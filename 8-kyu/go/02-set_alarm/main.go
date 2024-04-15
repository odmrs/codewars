package main

func main() {

}

func SetAlarm(employed, vacation bool) bool {
	if employed && !vacation {
		return true
	}
	return false
}
