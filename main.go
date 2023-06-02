package main

import (
	"fmt"
	"math/rand"
)

//O programa deve gerar um número aleatório entre 1 e 100 que será a resposta.
//O usuário deve tentar adivinhar o número, digitando um valor inteiro.
//O programa deve comparar o valor digitado com a resposta e dar feedback ao usuário, informando se o número é maior ou menor que a resposta.
//O usuário deve continuar tentando até acertar o número.
//O programa deve armazenar o número de tentativas do usuário em um vetor.
//Quando o usuário acertar o número, o programa deve exibir a quantidade de tentativas que o usuário utilizou para acertar.
//O programa deve perguntar ao usuário se ele deseja jogar novamente.
//O programa deve armazenar o número de tentativas utilizadas em todas as jogadas em uma matriz.
//O usuário pode jogar quantas vezes quiser.
//O programa deve permitir que o usuário visualize o número de tentativas utilizadas em todas as jogadas.

func main() {
	var chute int
	var resposta string
	tentativas := make([]int, 1)
	r := 0 //rodadas
	numero := rand.Intn(100)

	fmt.Println("Bem-vindo ao jogo da adivinhação!")
	for {
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&chute)
		if numero < chute {
			fmt.Println("O número é menor que", chute, ".")
			tentativas[r] = tentativas[r] + 1
		} else if numero > chute {
			fmt.Println("O número é maior que", chute, ".")
			tentativas[r] = tentativas[r] + 1
		} else {
			tentativas[r] = tentativas[r] + 1
			fmt.Print("Parabéns, você acertou!\nVocê utilizou ", tentativas[r], " tentativas.\nVocê deseja jogar novamente? (s/n): ")
			fmt.Scan(&resposta)
			if resposta == "s" {
				numero = rand.Intn(100)
				r++
				tentativas = append(tentativas, 0)
				continue
			} else {
				for i := 0; i <= r; i++ {
					fmt.Println("Você utilizou ", tentativas[i], " tentativas na ", i+1, "ª jogada.")
				}
				break
			}
		}
	}
}
