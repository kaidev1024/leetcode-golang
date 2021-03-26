func angleClock(hour int, minutes int) float64 {
	hourAngle := float64(hour % 12 * 30)

	minuteAngle := float64(minutes * 6)
	hourAngle += minuteAngle / 12

	dif := hourAngle - minuteAngle
	if dif < 0 {
		dif = -dif
	}
	if dif > 180 {
		return 360 - dif
	}
	return dif
}