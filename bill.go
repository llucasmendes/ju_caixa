package main

import (
	"fmt"
	"os"
)

type bill struct {
	name     string
	itens    map[string]float64
	tip      float64
	quantity []float64
}

func newBill(name string) bill {
	b := bill{
		quantity: []float64{0},
		name:     name,
		itens:    map[string]float64{},
		tip:      0,
	}
	return b
}
func (b *bill) formatBill() string {
	fs := "\n\n        Ju Gourmet 🧁❤ \n"
	var total float64 = 0
	i := 0
	for k, v := range b.itens {
		fs += fmt.Sprintf("%vx %-20v ... R$%v \n", b.quantity[i], k+"", v)
		total += v
		i++
	}

	fs += fmt.Sprintf("\n%-20v ... %0.2f⌘\n", "Desconto: R$", b.tip)
	fs += fmt.Sprintf("%-20v ... %0.2f", "Valor Desconto: R$", (total * (b.tip / 100.0)))
	fs += fmt.Sprintln("\n-----------------------------")
	fs += fmt.Sprintf("%-20v ... %0.2f\n", "Total: R$", total)
	fs += fmt.Sprintf("%-20v ... %0.2f", "Total com desconto: R$", total-(total*(b.tip/100.0)))
	return fs
}
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}
func (b *bill) addItem(name string, price float64, quantity float64, counter int) {
	b.itens[name] = price * quantity
	//TODO: bug no slice de quantidades
	b.quantity[counter] = quantity
}
func (b *bill) removeItem(i string) {
	delete(b.itens, i)
}

func (b *bill) save() {
	data := []byte(b.formatBill())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("A conta foi salva no arquivo!")
}
