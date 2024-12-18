package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	options := []string{"pedra", "papel", "tesoura"}

	fmt.Println("Jogador 1, escolha pedra, papel ou tesoura:")
	player1Choice := getPlayerChoice(options)

	fmt.Println("Jogador 2, escolha pedra, papel ou tesoura:")
	player2Choice := getPlayerChoice(options)

	fmt.Printf("Jogador 1 escolheu: %s\n", player1Choice)
	fmt.Printf("Jogador 2 escolheu: %s\n", player2Choice)

	result := determineWinner(player1Choice, player2Choice)
	fmt.Println(result)
}

func getPlayerChoice(options []string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada.")
			continue
		}
		choice := strings.TrimSpace(strings.ToLower(userInput))

		if isValidChoice(choice, options) {
			return choice
		}

		fmt.Println("Escolha inválida! Por favor, escolha entre pedra, papel ou tesoura:")
	}
}

func isValidChoice(choice string, options []string) bool {
	for _, option := range options {
		if choice == option {
			return true
		}
	}
	return false
}

func determineWinner(player1Choice, player2Choice string) string {
	winMap := map[string]string{
		"pedra":   "tesoura",
		"papel":   "pedra",
		"tesoura": "papel",
	}

	if player1Choice == player2Choice {
		return "Empate!"
	}

	if winMap[player1Choice] == player2Choice {
		return "Jogador 1 ganhou!"
	}

	return "Jogador 2 ganhou!"
}
