package ParseData

var OldWeightValue = 0
var faultWeight = 40

func CheckData(weightBox int) (checkData bool) {
	// значение нуля попадает сюда когда весы не готовы, если весы не готовы значит ждем калибровки с последующим автозабитием
	if weightBox <= 0 {
		OldWeightValue = 0
	}

	// значение весов регулирует авто забитие, оно происходит только при изменение весе, если вес не изменялся то автозабитие не происходит

	if weightBox > 0 {
		if (OldWeightValue - faultWeight) <= weightBox && weightBox <= (OldWeightValue + faultWeight) {
			return false
		} else {
			OldWeightValue = weightBox
			return true
		}
	} else {
		return false
	}
}
