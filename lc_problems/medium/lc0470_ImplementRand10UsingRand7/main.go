func rand10() int {
	return helper()
}

func helper() int {
	ret := rand7() + (rand7()-1)*7
	if ret > 40 {
		return helper()
	}
	return (ret-1)/4 + 1
}