package main

import (
	"fmt"
	"strings"
)

type ContactInfo struct {
	Type  string 
	Value string 
}

type Contact struct {
	ID      int
	Name    string
	Info    []ContactInfo
}

type ContactManager struct {
	contacts []Contact
}

func (cm *ContactManager) AddContact(name string, infoList []ContactInfo) {
	newContact := Contact{
		ID:   len(cm.contacts) + 1,
		Name: name,
		Info: infoList,
	}
	cm.contacts = append(cm.contacts, newContact)
	fmt.Println("Контакт добавлен:", name)
}

func (cm *ContactManager) PrintAllContacts() {
	for _, contact := range cm.contacts {
		fmt.Println("--------------------------------------------------")
		fmt.Println("ID:", contact.ID)
		fmt.Println("Имя:", contact.Name)
		for _, info := range contact.Info {
			fmt.Println(info.Type,":", info.Value)
		}
	}
	fmt.Println("--------------------------------------------------")
}

func (cm *ContactManager) DeleteContact() {
	cm.PrintAllContacts()
	fmt.Println("Введите ID контакта для удаления:")
	var id int
	fmt.Scanln(&id)
	for i, contact := range cm.contacts {
		if contact.ID == id {
			fmt.Println("Контакт", contact.Name, "удалён.")
			cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
			return
		}
	}
	fmt.Println("Контакт с таким ID не найден.")
}

func (cm *ContactManager) SearchContacts() {
	fmt.Println("Введите строку для поиска (имя или значение информации):")
	var query string
	fmt.Scanln(&query)
	query = strings.ToLower(query)
	found := false
	for _, contact := range cm.contacts {
		if strings.Contains(strings.ToLower(contact.Name), query) {
			fmt.Println("Найден контакт:", contact.Name)
			for _, info := range contact.Info {
				fmt.Println(info.Type+":", info.Value)
			}
			fmt.Println("--------------------------------------------------")
			found = true
			continue
		}
		for _, info := range contact.Info {
			if strings.Contains(strings.ToLower(info.Value), query) {
				fmt.Println("Найден контакт:", contact.Name)
				for _, i := range contact.Info {
					fmt.Println(i.Type+":", i.Value)
				}
				fmt.Println("--------------------------------------------------")
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("Ничего не найдено по запросу:", query)
	}
}

func main() {
	var cm ContactManager
	cm.AddContact("Алексей Иванов", []ContactInfo{
		{Type: "телефон", Value: "+79991234567"},
		{Type: "email", Value: "alex@example.com"},
	})
	cm.AddContact("Мария Смирнова", []ContactInfo{
		{Type: "email", Value: "maria.smirnova@mail.com"},
		{Type: "адрес", Value: "Москва, ул. Пушкина, д.10"},
	})
	cm.PrintAllContacts()
	cm.SearchContacts()
	cm.DeleteContact()
	cm.PrintAllContacts()
}
