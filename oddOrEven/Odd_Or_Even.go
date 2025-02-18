package oddoreven

func OddOrEven(value int) bool {

	remainder := value % 2

	if remainder == 0 {
		return true

	} else {
		return false
	}
}
