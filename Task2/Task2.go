package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	products map[int]*Product
}

func (inventory *Inventory) AddProduct(product *Product) {
	if inventory.products == nil {
		inventory.products = make(map[int]*Product)
	}
	fmt.Println("продукт:", product.Name, "был добавлен в инвентарь")
	inventory.products[product.ID] = product
}

func (inventory *Inventory) WriteOff(productID int, quantity int) error {
	value, ok := inventory.products[productID]
	if !ok {
		return errors.New("Данного предмета на складе нет")
	}
	if quantity < 0 {
		return errors.New("Колличество товаров для списания должно быть больше нуля")
	}
	if value.Quantity < quantity {
		return errors.New("Недостаточно товаров для списания")
	}
	fmt.Println(value.Name, ":", quantity, "кол.", "было списано со склада")
	value.Quantity -= quantity
	return nil

}

func (inventory *Inventory) RemoveProduct(productID int) error {
	value, ok := inventory.products[productID]
	if ok {
		fmt.Println(value.Name, "- предмет был выброшен")
		delete(inventory.products, productID)
		return nil
	} else {
		return errors.New("Данного предмета на складе нет")
	}
}

func (inventory *Inventory) GetTotalValue() float64 {
	var Summ float64
	for product := range inventory.products {
		Summ += float64(inventory.products[product].Quantity) * inventory.products[product].Price
	}
	return Summ
}

func main() {
	var Inventory1 Inventory
	Bread := Product{
		ID:       1,
		Name:     "Хлэбчик",
		Price:    50,
		Quantity: 15,
	}
	Rabbit := Product{
		ID:       2,
		Name:     "Зайчик",
		Price:    100,
		Quantity: 20,
	}
	Inventory1.AddProduct(&Bread)
	Inventory1.WriteOff(Bread.ID, 5)
	Inventory1.RemoveProduct(Bread.ID)
	Inventory1.AddProduct(&Bread)
	Inventory1.WriteOff(Bread.ID, 5)
	Inventory1.AddProduct(&Rabbit)
	fmt.Println("Сумма всех товаров", Inventory1.GetTotalValue())
}
