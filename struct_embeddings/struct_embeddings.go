package structembeddings

import "fmt"

type BaseForm struct {
	AccountNumber string 
	Name string
	Age int
	Country string
}

type WaterBillForm struct { 
	BaseForm
	Liters int
}

type ElectricityBillForm struct { 
	BaseForm
	Kwh int
}

func (b *BaseForm) UserDetails() BaseForm { 
	return *b
}

func (w *WaterBillForm) WaterBillDetails() WaterBillForm { 
	return *w
}

type Receipt interface { 
	UserDetails() BaseForm
	WaterBillDetails() WaterBillForm
}

func StructEmbeddings() {

	account := BaseForm{
		AccountNumber: "123456789",
		Name: "John",
		Age: 25,
		Country: "USA",
	}

	waterBillForm := WaterBillForm{
		BaseForm: account,
		Liters: 100,
	}

	// var receipt Receipt = &waterBillForm
	// fmt.Printf("Water Bill Receipt: %+v\n", receipt.WaterBillDetails())

	receipt := Receipt(&waterBillForm)
	fmt.Printf("Water Bill Receipt: %+v\n", receipt.WaterBillDetails())
}