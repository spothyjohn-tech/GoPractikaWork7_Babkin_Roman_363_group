package main

import "fmt"

type OrderItem struct {
	ID       int
	ItemName string
	Price    float64
	Quantity int
}
type Customer struct {
	ID       int
	Username string
	Adress   string
}
type Order struct {
	ID        int
	customer  Customer
	orderItem []OrderItem
	Status    string
}

func (order *Order) OrderSum() float64 {
	var SumAllOrder float64
	for _, Item := range order.orderItem {
		fmt.Println(Item.ItemName)
		SumAllOrder += Item.Price * float64(Item.Quantity)
	}
	return SumAllOrder
}

func (order *Order) AddItem(AddItem OrderItem) {
	if (AddItem.Quantity > 0) {
		order.orderItem = append(order.orderItem, AddItem)
		fmt.Println("Товар:", AddItem.ItemName, "был добавлен в заказ")
	} else {
		fmt.Println("Для добавления товара в заказ его колличество должно быть больше нуля")
	}
}

func (order *Order) RemoveItem(RemItem OrderItem) {
	for i, Item := range order.orderItem {
		if Item.ID == RemItem.ID {
			order.orderItem = append(order.orderItem[:i], order.orderItem[i+1:]...)
			fmt.Println("Товар:", RemItem.ItemName, "был удалён из заказа")
			return
		}
	}
	fmt.Println("Данного товара нет в заказе")
}

func (order *Order) ChangeStatus(NewStatus string) {
	if NewStatus != "" {
		order.Status = NewStatus
	}
}

func main() {
	TestCustomer := Customer{
		ID:       1,
		Username: "Иван Иванов",
		Adress:   "ул. Ленина 5",
	}
	Bread := OrderItem{
		ID:       1,
		ItemName: "Хлеб",
		Price:    20,
		Quantity: 2,
	}
	Eggs := OrderItem{
		ID:       2,
		ItemName: "Яйца",
		Price:    40,
		Quantity: 3,
	}
	Milk := OrderItem{
		ID:       3,
		ItemName: "Молоко",
		Price:    50,
		Quantity: 1,
	}
	TestOrder := Order{
		ID:        1,
		customer:  TestCustomer,
		orderItem: []OrderItem{},
		Status:    "В обработке",
	}
	TestOrder.ChangeStatus("Ожидайте курьера")
	TestOrder.AddItem(Milk)
	TestOrder.AddItem(Eggs)
	TestOrder.AddItem(Bread)
	TestOrder.RemoveItem(Bread)
	fmt.Println("Сумма за весь заказ:")
	fmt.Println(TestOrder.OrderSum())
}
