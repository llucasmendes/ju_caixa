package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Criar uma nova conta -> ", reader)

	b := newBill(name)
	fmt.Println("Nova conta criada - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nEscolha uma opcao:\n\na - adicionar item\nd - adicionar desconto\ns - salvar conta    | ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Nome do produto: ", reader)
		price, _ := getInput("Preço do produto: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("O preço deve ser um número!!!")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item adicionado: ", name, price)
		promptOptions(b)
	case "s":
		fmt.Println(b.formatBill())
		b.save()
	case "d":
		tip, _ := getInput("Entre com a porcentagem do desconto: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("O desconto deve ser um número!!!")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Desconto de:", tip)
		promptOptions(b)
	default:
		fmt.Println("Opcao incorreta...")
		promptOptions(b)
	}
}
