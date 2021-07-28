package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	itens map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		itens: map[string]float64{},
		tip:   0,
	}
	return b
}
func (b *bill) formatBill() string {
	fs := "        Ju Gourmet \n"
	var total float64 = 0

	for k, v := range b.itens {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+"", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ... %0.2f\n", "desconto:", b.tip)
	fs += fmt.Sprintf("------------------------------------\n")
	fs += fmt.Sprintf("%-25v ... %0.2f", "total:", total-(total*(b.tip/100.0)))
	return fs
}
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}
func (b *bill) addItem(name string, price float64) {
	b.itens[name] = price
}

func (b *bill) save() {
	data := []byte(b.formatBill())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("A conta foi salva no arquivo!")
}
