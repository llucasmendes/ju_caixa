package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var counter int = 0

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

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	fmt.Println(runtime.GOOS)
	if ok {
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("\nEscolha uma opcao:\n\na - adicionar item\ne - excluir item\nd - adicionar desconto\nv - ver conta atual\nf - forma de pagemento\ns - salvar conta    | ", reader)
	switch opt {
	case "a":
		name, _ := getInput("\nNome do produto: ", reader)
		quantity, _ := getInput("\nQuantidade do produto: ", reader)
		price, _ := getInput("\nPreço do produto: ", reader)
		q, erro := strconv.ParseFloat(quantity, 64)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil || erro != nil {
			fmt.Println("\nO preço/quantidade deve ser um número!!!")
			time.Sleep(2 * time.Second)
			CallClear()
			promptOptions(b)
		}
		if q <= 0.0 {
			q = 1
		}

		b.addItem(name, p, q, counter)
		counter++
		fmt.Println("\nItem adicionado: ", name, quantity, "x", price)
		time.Sleep(2 * time.Second)
		CallClear()
		promptOptions(b)

	case "s":
		fmt.Println(b.formatBill())
		b.save()
		a, _ := getInput("\nCompra finalizada ", reader)
		fmt.Print(a)
	case "v":
		fmt.Println(b.formatBill())
		promptOptions(b)
	case "f":
		pagamento, _ := getInput("Entre com a forma de pagamento: ", reader)
		b.payment = pagamento
		promptOptions(b)
	case "e":
		i := 0
		s := []string{}

		for k := range b.itens {
			fmt.Println(i, "--", k)
			s = append(s, k)
			i++
		}
		a, _ := getInput("\nEscolha um item: ", reader)
		p, err := strconv.ParseInt(a, 0, 64)
		if err != nil {
			fmt.Println("\nO item deve ser um número!!!")
			time.Sleep(2 * time.Second)
			CallClear()
			promptOptions(b)
		}
		b.removeItem(s[p])
		time.Sleep(2 * time.Second)
		CallClear()

		promptOptions(b)
	case "d":
		tip, _ := getInput("Entre com a porcentagem do desconto: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("\nO desconto deve ser um número!!!")
			time.Sleep(2 * time.Second)
			CallClear()
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("\nDesconto de:", tip)
		time.Sleep(2 * time.Second)
		CallClear()
		promptOptions(b)
	default:
		fmt.Println("\nOpcao incorreta...")
		time.Sleep(2 * time.Second)
		CallClear()
		promptOptions(b)
	}
}
