package main

import "fmt"

func calcularSoma(inicio, fim int, ch chan int) {
	soma := 0
	for i := inicio; i <= fim; i++ {
		soma += i
	}
	ch <- soma // Envia o resultado para o canal
}

func main() {
	n := 100 // Valor até o qual queremos calcular a soma
	soma := 0

	ch := make(chan int) // Canal para receber os resultados

	numGoroutines := 4 // Número de goroutines para dividir o trabalho
	passo := n / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		inicio := i*passo + 1
		fim := (i + 1) * passo
		if i == numGoroutines-1 {
			fim = n
		}
		go calcularSoma(inicio, fim, ch) // Inicia uma goroutine para cada intervalo
	}

	// Aguarda as goroutines finalizarem e recebe os resultados do canal
	for i := 0; i < numGoroutines; i++ {
		somaParcial := <-ch // Recebe o resultado do canal
		soma += somaParcial
	}

	fmt.Println("A soma de todos os números até", n, "é:", soma)
}

//package main
//
//import "fmt"
//
//func main() {
//	n := 1000 // Valor até o qual queremos calcular a soma
//	soma := 0
//
//	for i := 1; i <= n; i++ {
//		soma += i
//	}
//
//	fmt.Println("A soma de todos os números até", n, "é:", soma)
//}
