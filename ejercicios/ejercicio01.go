package ejercicios

import "strconv"

func ConviertoString(palabra string) (int, string) {
	numero, _ := strconv.Atoi(palabra)
	return numero, palabra
}
