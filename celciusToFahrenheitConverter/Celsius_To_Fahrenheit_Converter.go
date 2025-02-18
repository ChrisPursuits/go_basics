package celciusToFahrenheitConverter

func Converter(celciusValue float64) float64 {
	const FAHRENHEIT_CONVERSION_VALUE float64 = 9.0 /5.0;

	if (celciusValue == 0) {
		return 32;

	} else {
		return (celciusValue * FAHRENHEIT_CONVERSION_VALUE) + 32;
	}
}

