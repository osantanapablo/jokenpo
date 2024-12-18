package main

import (
	"testing"
)

func TestIsValidChoice_WithValidChoice(t *testing.T) {
	options := []string{"pedra", "papel", "tesoura"}

	if !isValidChoice("pedra", options) {
		t.Errorf("isValidChoice(pedra) = false; expected true")
	}

}

func TestIsValidChoice_WithInValidChoice(t *testing.T) {
	options := []string{"pedra", "papel", "tesoura"}

	if isValidChoice("spock", options) {
		t.Errorf("isValidChoice(spock) = true; expected false")
	}
}

func TestDetermineWinner_WithPlayer1Winner(t *testing.T) {
	if result := determineWinner("pedra", "tesoura"); result != "Jogador 1 ganhou!" {
		t.Errorf("determineWinner(pedra, tesoura) = %s; expected Jogador 1 ganhou!", result)
	}
}

func TestDetermineWinner_WithPlayer2Winner(t *testing.T) {
	if result := determineWinner("tesoura", "pedra"); result != "Jogador 2 ganhou!" {
		t.Errorf("determineWinner(tesoura, pedra) = %s; expected Jogador 2 ganhou!", result)
	}
}

func TestDetermineWinner_WithDraw(t *testing.T) {
	if result := determineWinner("tesoura", "tesoura"); result != "Empate!" {
		t.Errorf("determineWinner(tesoura, tesoura) = %s; expected Empate!", result)
	}
}
