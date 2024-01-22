package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresoPorTeclado() {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero para obtener su tabla de multiplicar: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i < 11; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
	}
}
