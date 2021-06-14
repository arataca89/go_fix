/*
Corrigindo comportamento errado ao usar fmt.Scanf() no Windows

arataca89@gmail.com
20210614
*/
package main

import "fmt"

func main() {
	var n int
	opcao := -1

	fmt.Println("\n>>> Erro no fmt.Scanf() no Windows <<<")

	fmt.Printf("\nEntre com o valor de n: ")
	fmt.Scanf("%d", &n)

	for {
		fmt.Println()
		fmt.Println("<< 1 >> Opção 1")
		fmt.Println("<< 2 >> Opção 2")
		fmt.Println("<< 3 >> Opção 3")
		fmt.Println("<< 0 >> Encerrar")

		fmt.Print("Entre com sua opção: ")

		// Usar esta fmt.Scanf() causou comportamento indesejado
		//fmt.Scanf("%d", &opcao)

		fmt.Scan(&opcao)

		if opcao == 0 {
			break
		} else {
			fmt.Println("\nOpção escolhida:", opcao)
			fmt.Println("Valor de n:", n)
		}
	}

	fmt.Println("\nPrograma encerrado!")
}

/********************************************************************
	SAÍDA USANDO 		fmt.Scanf("%d", &opcao)
*********************************************************************
>>> Erro no fmt.Scanf() no Windows <<<

Entre com o valor de n: 13

<< 1 >> Opção 1
<< 2 >> Opção 2
<< 3 >> Opção 3
<< 0 >> Encerrar
Entre com sua opção:
Opção escolhida: -1
Valor de n: 13

<< 1 >> Opção 1
<< 2 >> Opção 2
<< 3 >> Opção 3
<< 0 >> Encerrar
Entre com sua opção:


*********************************************************************
	SAÍDA USANDO	fmt.Scan(&opcao)
*********************************************************************
>>> Erro no fmt.Scanf() no Windows <<<

Entre com o valor de n: 13

<< 1 >> Opção 1
<< 2 >> Opção 2
<< 3 >> Opção 3
<< 0 >> Encerrar
Entre com sua opção:


*********************
*	COMENTÁRIOS		*
*********************
Observe que ao usar fmt.Scanf() houve um erro. O programa não
esperou a entrada do úsuario na primeira interação do laço,
exibindo assim o menu duas vezes.

Isto deve-se a que no Windows a tecla ENTER equivale a "\r\n" e o
pacote fmt trata apenas o "\n".

A correção deste comportamento foi usar a função fmt.Scan() no lugar
de fmt.Scanf().

Referências:
https://socketloop.com/tutorials/golang-fix-fmt-scanf-on-windows-will-scan-input-twice-problem
https://github.com/golang/go/issues/5391
*/
