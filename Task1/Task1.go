package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	acountNumber int
	holderName   string
	balance      float64
}

func (Account *BankAccount) Deposit(amount float64) {
	fmt.Println("На аккаунт: ", Account.holderName, "было зачислено: ", amount)
	Account.balance += amount
}

func (Account *BankAccount) Withdraw(amount float64) error {
	if(amount < 0){
		return errors.New("Колличество средств для списания должно быть больше нуля")
	}
	if Account.balance-amount < 0 {
		return errors.New("Операция не прошла: недостаточно средств")
	} else {
		fmt.Println("С аккаунта: ", Account.holderName, "было снято: ", amount)
		Account.balance -= amount
		return nil
	}
}

func (Account *BankAccount) GetBalance() float64 {
	return Account.balance
}

func main() {
	Account1 := BankAccount{
		acountNumber: 1,
		holderName:   "Клиент_1",
		balance:      20,
	}
	fmt.Println("Сейчас на счету:")
	fmt.Println(Account1.GetBalance())
	Account1.Deposit(50)
	fmt.Println("Сейчас на счету:")
	fmt.Println(Account1.GetBalance())
	err1 := Account1.Withdraw(50)
	if err1 != nil {
		fmt.Println("Произошла ошибка:", err1)
	}
	fmt.Println("Сейчас на счету:")
	fmt.Println(Account1.GetBalance())
	err2 := Account1.Withdraw(50)
	if err2 != nil {
		fmt.Println("Произошла ошибка:", err2)
	}
	fmt.Println("Сейчас на счету:")
	fmt.Println(Account1.GetBalance())
}
