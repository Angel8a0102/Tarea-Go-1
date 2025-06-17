
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// aca usamos bufio para leer entrada del usuario
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese una lista de números separados por comas:")
	input, _ := reader.ReadString('\n')

	
	input = strings.TrimSpace(input)

	// se puede dividir la cadena en partes usando strings.Split
	numberStrings := strings.Split(input, ",")

	// aca convertimos strings a números usando strconv
	var numbers []int
	for _, ns := range numberStrings {
		n, err := strconv.Atoi(strings.TrimSpace(ns))
		if err != nil {
			fmt.Printf("'%s' no es un número válido, se omitirá\n", ns)
			continue
		}
		numbers = append(numbers, n)
	}

	// aca se puede mostrar la lista original
	fmt.Println("\nLista original:", numbers)

	// con el siguiente podemos ordenar usando sort.Ints
	sort.Ints(numbers)
	fmt.Println("Lista ordenada (ascendente):", numbers)

	// como arriba lo hemos colocado podemos ordenar en orden descendente usando sort.Slice
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	fmt.Println("Lista ordenada (descendente):", numbers)

	// algunas operaciones basicas con slice
	fmt.Println("\nOperaciones con slices:")
	fmt.Println("Longitud del slice:", len(numbers))
	fmt.Println("Capacidad del slice:", cap(numbers))

	if len(numbers) > 0 {
		fmt.Println("Primer elemento:", numbers[0])
		fmt.Println("Último elemento:", numbers[len(numbers)-1])

		// Sub-slice (primeros 3 elementos)
		if len(numbers) >= 3 {
			firstThree := numbers[:3]
			fmt.Println("Primeros 3 elementos:", firstThree)
		}

		// Sub-slice (últimos 2 elementos)
		if len(numbers) >= 2 {
			lastTwo := numbers[len(numbers)-2:]
			fmt.Println("Últimos 2 elementos:", lastTwo)
		}
	}

	// Buscar un elemento usando sort.SearchInts (requiere slice ordenado ascendentemente)
	sort.Ints(numbers) // Re-ordenamos ascendentemente para la búsqueda
	fmt.Println("\nLista reordenada para búsqueda:", numbers)

	fmt.Println("Ingrese un número a buscar:")
	searchInput, _ := reader.ReadString('\n')
	searchInput = strings.TrimSpace(searchInput)
	searchNum, err := strconv.Atoi(searchInput)
	if err != nil {
		fmt.Println("Número inválido para búsqueda")
		return
	}

	// buscar el número
	pos := sort.SearchInts(numbers, searchNum)
	if pos < len(numbers) && numbers[pos] == searchNum {
		fmt.Printf("El número %d se encuentra en la posición %d\n", searchNum, pos)
	} else {
		fmt.Printf("El número %d no se encuentra en la lista\n", searchNum)
	}

	// eliminar duplicados usando slices
	fmt.Println("\nEliminando duplicados...")
	uniqueNumbers := make([]int, 0, len(numbers))
	seen := make(map[int]bool)

	for _, num := range numbers {
		if !seen[num] {
			uniqueNumbers = append(uniqueNumbers, num)
			seen[num] = true
		}
	}

	fmt.Println("Lista sin duplicados:", uniqueNumbers)
}