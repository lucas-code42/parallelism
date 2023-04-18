package main

import "fmt"

// calcularSoma
// parâmetros:
// inicio e fim são os parâmetros que devem ser "constantes" inicio=0 e fim=250
// isso é para garantir que as threads (goroutines) vão estar fazendo o cálculo na mesma proporção para que a tarefa
// seja igualmente dividia
// já o ch é um "pipe" para que o valor da soma possa ser enviado de volta para a thread main dentro da func main()
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

	// Canal para receber os resultados, declarado na thread principal da aplicação
	// Essa é a variável que vai poder transitar entre as threads da aplicação
	ch := make(chan int)

	numGoroutines := 4         // Número de goroutines para dividir o trabalho
	passo := n / numGoroutines // aqui vamos de fato dividir a contagem para cada go routine

	for i := 0; i < numGoroutines; i++ {
		inicio := i*passo + 1  // aqui estamos garantindo que o início da contagem seja sempre 0
		fim := (i + 1) * passo // aqui estamos garantindo que o fim da contagem seja sempre o valor de passo

		// como queremos triggar nossa goroutine 4 vezes como declarado na variável numGoroutines
		// aqui garantimos que ela vai ser triggada 4 vezes mesmo
		if i == numGoroutines-1 {
			fim = n
		}

		// aqui é onde a magica acontece.
		// podemos entender uma goroutine em termos simplistas como uma thread isolada, isto é, que não está
		// compartilhando memória com as demais threads dentro deste processo ou PID.
		// O que garante que iremos conseguir paralelizar o processo vai ser o ch (chan)
		// que é um tipo de dado especial do Go, e ele sim vai conseguir transitar entre
		// a thread que o ch foi enviado e a thread principal
		go calcularSoma(inicio, fim, ch) // Inicia uma goroutine para cada intervalo
	}

	// Aguarda as goroutines finalizarem e recebe os resultados do canal
	for i := 0; i < numGoroutines; i++ {
		somaParcial := <-ch // Recebe o resultado do canal
		soma += somaParcial
	}

	fmt.Println("A soma de todos os números até", n, "é:", soma)
}

// Esse é o mesmo código sem o uso de goroutines e channels

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
