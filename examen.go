package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//ejemplo de strconv
	numeroEntero := "1234"
	base := 10
	bits := 64

	entero64, err := strconv.ParseInt(numeroEntero, base, bits)
	if err != nil {
		fmt.Println("Error al convertir a entero:", err)
	} else {
		fmt.Printf("Entero convertido: %d (tipo: %T)\n", entero64, entero64)
	}

	// Ejemplo con base diferente (binario)
	binario := "1010"
	enteroBinario, err := strconv.ParseInt(binario, 2, 64)
	if err != nil {
		fmt.Println("Error al convertir binario:", err)
	} else {
		fmt.Printf("Binario '%s' convertido: %d\n", binario, enteroBinario)
	}

	// ejemplo de parsefloat
	numeroDecimal := "3.141592"
	decimal64, err := strconv.ParseFloat(numeroDecimal, 64)
	if err != nil {
		fmt.Println("Error al convertir a float:", err)
	} else {
		fmt.Printf("Float convertido: %f (tipo: %T)\n", decimal64, decimal64)
	}

	// algunos pequeños ejemplos de strings
	cadena := "   Hola Mundo!   "
	fmt.Println("\nEjemplos con funciones strings:")
	fmt.Println("Original:", cadena)
	fmt.Println("TrimSpace:", strings.TrimSpace(cadena))
	fmt.Println("ToUpper:", strings.ToUpper(cadena))
	fmt.Println("ToLower:", strings.ToLower(cadena))
	fmt.Println("Prefix 'Hola':", strings.HasPrefix(cadena, "Hola"))
	fmt.Println("Suffix 'Ejemplo!':", strings.HasSuffix(cadena, "Go!"))
	fmt.Println("Contiene 'Mundo':", strings.Contains(cadena, "Mundo"))
	fmt.Println("Reemplazar:", strings.Replace(cadena, "Go", "Golang", 1))

	// como convertir un numero a un string
	valorInt := 42
	valorFloat := 3.1416
	fmt.Println("\nConversión de números a strings:")
	fmt.Println("Entero a string:", strconv.Itoa(valorInt))
	fmt.Println("Entero a string (FormatInt):", strconv.FormatInt(int64(valorInt), 10))
	fmt.Println("Float a string:", strconv.FormatFloat(valorFloat, 'f', 2, 64))

	// aca usamos los Split y Join
	frases := "Go,Python,Java,JavaScript"
	splitResult := strings.Split(frases, ",")
	fmt.Println("\nSplit:", splitResult)
	fmt.Println("Join con |:", strings.Join(splitResult, " | "))

	// aca convertimos el atoi (más simple que ParseInt para enteros)
	edadStr := "30"
	edad, err := strconv.Atoi(edadStr)
	if err != nil {
		fmt.Println("Error al convertir edad:", err)
	} else {
		fmt.Printf("\nEdad convertida: %d (tipo: %T)\n", edad, edad)
	}
}