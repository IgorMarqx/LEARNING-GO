package calc

import (
	"errors"
	"fmt"
)

func Operation(operation string) (int, error) {
	num1 := 5
	num2 := 4
	var result int

	if operation == "dividir" && num2 == 0 {
		return 0, errors.New("Não é possivel dividir por 0")
	}

	arrayOperations := [4]string{"somar", "multiplicar", "dividir", "subtrair"}
	var validOperation bool

	for _, value := range arrayOperations {
		if value == operation {
			validOperation = true
			break
		}
	}

	if !validOperation {
		return 0, errors.New("A operação não reconhecida")
	}

	switch operation {
	case "somar":
		result = num1 + num2
	case "multiplicar":
		result = num1 * num2
	case "dividir":
		result = num1 / num2
	case "subtrair":
		result = num1 - num2
	}

	fmt.Println("O resultado da", operation, "é", result)
	return result, nil
}
