package main

import (
	"fmt"
	"os"
)

type bill struct {
	name     string
	itens    map[string]float64
	tip      float64
	quantity [100]float64
	payment  string
}

func newBill(name string) bill {
	b := bill{
		quantity: [100]float64{1},
		name:     name,
		itens:    map[string]float64{},
		tip:      0,
		payment:  "Pix",
	}
	return b
}
func (b *bill) formatBill() string {
	fs := "\n\n        Ju Gourmet 🧁❤ \n"
	var total float64 = 0
	i := 0
	for k, v := range b.itens {
		fs += fmt.Sprintf("%vx %-20v ... R$ %0.2f \n", (*b).quantity[i], k+"", v)
		total += v
		i++
	}

	fs += fmt.Sprintf("\n%-20v ... %0.2f⌘\n", "Desconto: ⌘", b.tip)
	fs += fmt.Sprintf("%-20v ... %0.2f", "Valor Desconto: R$", (total * (b.tip / 100.0)))
	fs += fmt.Sprintln("\n-----------------------------")
	fs += fmt.Sprintf("%-20v ...R$ %0.2f\n", "Total: ", total)
	fs += fmt.Sprintf("%-20v ... R$ %0.2f\n", "Total com desconto: ", total-(total*(b.tip/100.0)))
	fs += fmt.Sprintf("FORMA DE PAGAMENTO: %v", b.payment)
	return fs
}
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}
func (b *bill) addItem(name string, price float64, quantity float64, counter int) {
	(*b).itens[name] = price * quantity

	(*b).quantity[counter] = quantity
}
func (b *bill) removeItem(i string) {
	delete(b.itens, i)
}

// TODO: funcao para mudar a quantidade de produtos

func (b *bill) save() {
	data := []byte(b.formatBill())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("A conta foi salva no arquivo!")
}
