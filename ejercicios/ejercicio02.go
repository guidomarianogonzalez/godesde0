package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func IngresoPorTeclado() int {
	fmt.Println("Ingrese numero para obtener su tabla de multiplicar: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado esta mal " + err.Error())
		}
		return numero
	}
	return 0
}
