package main

import (
	"fmt"
	"time"
)

type Event struct {
	ID           int
	Title        string      
	Description  string      
	Date         time.Time  
	Location     string      
	MaxAttendees int        
	Attendees    []string    
}

type EventManager struct {
	events       []Event
	nextEventID  int
}
func (em *EventManager) AddEvent(title, description, location string, date time.Time, maxAttendees int) {
	newEvent := Event{
		ID:           em.nextEventID,
		Title:        title,
		Description:  description,
		Date:         date,
		Location:     location,
		MaxAttendees: maxAttendees,
		Attendees:    []string{},
	}
	em.events = append(em.events, newEvent)
	em.nextEventID++
	fmt.Println("Мероприятие добавлено:", title)
}

func (em *EventManager) RegisterAttendee(eventID int, name string) {
	for i := range em.events {
		if em.events[i].ID == eventID {
			if len(em.events[i].Attendees) >= em.events[i].MaxAttendees {
				fmt.Println("Мероприятие уже заполнено.")
				return
			}
			for _, attendee := range em.events[i].Attendees {
				if attendee == name {
					fmt.Println("Пользователь уже зарегистрирован.")
					return
				}
			}
			em.events[i].Attendees = append(em.events[i].Attendees, name)
			fmt.Println("Участник", name, "зарегистрирован на мероприятие:", em.events[i].Title)
			return
		}
	}
	fmt.Println("Мероприятие с ID", eventID, "не найдено.")
}

func (em *EventManager) CancelRegistration(eventID int, name string) {
	for i := range em.events {
		if em.events[i].ID == eventID {
			for j, attendee := range em.events[i].Attendees {
				if attendee == name {
					em.events[i].Attendees = append(em.events[i].Attendees[:j], em.events[i].Attendees[j+1:]...)
					fmt.Println("Регистрация отменена для", name, "на мероприятие:", em.events[i].Title)
					return
				}
			}
			fmt.Println("Пользователь", name, "не найден среди участников.")
			return
		}
	}
	fmt.Println("Мероприятие с таким ID не найдено.")
}

func (em *EventManager) PrintUpcomingEvents() {
	now := time.Now()
	found := false
	for _, event := range em.events {
		if event.Date.After(now) {
			fmt.Println("--------------------------------------------------")
			fmt.Println("ID:", event.ID)
			fmt.Println("Название:", event.Title)
			fmt.Println("Описание:", event.Description)
			fmt.Println("Дата:", event.Date.Format("02.01.2006 15:04"))
			fmt.Println("Место:", event.Location)
			fmt.Println("Макс. участников:", event.MaxAttendees)
			fmt.Println("Зарегистрировано:", len(event.Attendees))
			found = true
		}
	}
	if !found {
		fmt.Println("Нет предстоящих мероприятий.")
	}
	fmt.Println("--------------------------------------------------")
}

func (em *EventManager) PrintAllEvents() {
	for _, event := range em.events {
		fmt.Println("--------------------------------------------------")
		fmt.Println("ID:", event.ID)
		fmt.Println("Название:", event.Title)
		fmt.Println("Описание:", event.Description)
		fmt.Println("Дата:", event.Date.Format("02.01.2006 15:04"))
		fmt.Println("Место:", event.Location)
		fmt.Println("Макс. участников:", event.MaxAttendees)
		fmt.Println("Зарегистрировано:", len(event.Attendees))
		fmt.Println("Участники:", event.Attendees)
	}
	fmt.Println("--------------------------------------------------")
}

func main() {
	em := EventManager{
		events:      []Event{},
		nextEventID: 1,
	}
	event1Date := time.Date(2025, 12, 10, 18, 0, 0, 0, time.Local)
	event2Date := time.Date(2025, 11, 15, 14, 30, 0, 0, time.Local)
	em.AddEvent("Встреча Go-разработчиков", "Обсуждение решения 13-ой задачи на Go", "Москва", event1Date, 50)
	em.AddEvent("Фестиваль еды", "Вкусные блюда от шефов!", "Москва", event2Date, 100)
	em.RegisterAttendee(1, "Алексей")
	em.RegisterAttendee(1, "Мария")
	em.RegisterAttendee(2, "Иван")
	em.RegisterAttendee(1, "Мария") 
	em.CancelRegistration(1, "Алексей")
	em.PrintAllEvents()
	em.PrintUpcomingEvents()
}
