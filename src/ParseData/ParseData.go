package ParseData

import 	"../TransportData"


func ParseScaleData(data *TransportData.ScaleResponse) (weightBox float64) {
	/*
	   80 00
	   EC 00

	80             - готовность 128 - готов, 0 - не готов
	00 (1я строка) - 00-1г, 01-0.1г, 04-0.01г, 05-0.1кг
	EC             - вес в определившейся дискретности
	00 (2я строка) - 0 это (+) 80 это (-) отрицательный , положительный вес.
	*/

	if data.ReadyAndDiscreteness[0] == 128 && (data.Weight[0] !=0 || data.Weight[1] !=0) {

		// data.ReadyAndDiscreteness[0] - готовность
		// data.ReadyAndDiscreteness[1] - дискретность

		if data.ReadyAndDiscreteness[1] == 0 {
			if  data.Weight[1] == 0 { // вес уместился в 1н байт
				weightBox = float64(data.Weight[0]) * 1
				return
			}

			if data.Weight[1] != 0 { // не уместился
				weightBox = ((256 * float64(data.Weight[1])) + float64(data.Weight[0])) * 1
				return
			}
		}

		if data.ReadyAndDiscreteness[1] == 4 {
			if  data.Weight[1] == 0 { // вес уместился в 1н байт
				weightBox = float64(data.Weight[0]) * 10
				return
			}

			if data.Weight[1] != 0 { // не уместился
				weightBox = ((256 * float64(data.Weight[1])) + float64(data.Weight[0])) * 10
				return
			}
		}
	}

	return
}