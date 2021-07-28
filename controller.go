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

	name, _ := getInput("Criar uma nova conta --- ", reader)

	b := newBill(name)
	fmt.Println("\nNova conta criada - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nEscolha uma opcao:\n\na - adicionar item\ne - excluir item\nd - adicionar desconto\nv - ver conta atual\ns - salvar conta    | ", reader)
	switch opt {
	case "a":
		name, _ := getInput("\nNome do produto: ", reader)
		price, _ := getInput("Preço do produto: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("\nO preço deve ser um número!!!")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("\nItem adicionado: ", name, price)
		promptOptions(b)
	case "s":
		fmt.Println(b.formatBill())
		b.save()
		a, _ := getInput("\nCompra finalizada ", reader)
		fmt.Print(a)
	case "v":
		fmt.Println(b.formatBill())
		promptOptions(b)
	case "e":
		i := 0
		s := []string{}

		for k := range b.itens {
			fmt.Println(i, "--", k)
			s = append(s, k)
			i++
		}
		a, _ := getInput("\nEscolha um item:", reader)
		p, err := strconv.ParseInt(a, 0, 64)
		if err != nil {
			fmt.Println("\nO item deve ser um número!!!")
			promptOptions(b)
		}
		b.removeItem(s[p])

		promptOptions(b)
	case "d":
		tip, _ := getInput("Entre com a porcentagem do desconto: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("\nO desconto deve ser um número!!!")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("\nDesconto de:", tip)
		promptOptions(b)
	default:
		fmt.Println("\nOpcao incorreta...")
		promptOptions(b)
	}
}
