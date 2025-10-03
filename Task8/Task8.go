package main

import (
	"fmt"
)

type Room struct {
	ID       int
	Number   string
	Capacity int
	Price    float64
}
type Reservation struct {
	ID     int
	RoomID int
	Name   string
	Nights int
	Total  float64
}
type Hotel struct {
	ID           int
	Name         string
	Rooms        []Room
	Reservations []Reservation
}

func (hotel *Hotel) AddRoom(room Room) {
	hotel.Rooms = append(hotel.Rooms, room)
	fmt.Println("Добавлен номер:", room.Number)
}

func (hotel *Hotel) IsRoomAvailable(roomID int) bool {
	for _, res := range hotel.Reservations {
		if res.RoomID == roomID {
			return false
		}
	}
	return true
}

func (hotel *Hotel) CreateReservation(roomID int, name string, nights int) {
	for _, room := range hotel.Rooms {
		if room.ID == roomID {
			if hotel.IsRoomAvailable(roomID) {
				total := room.Price * float64(nights)
				newRes := Reservation{
					ID:     len(hotel.Reservations) + 1,
					RoomID: roomID,
					Name:   name,
					Nights: nights,
					Total:  total,
				}
				hotel.Reservations = append(hotel.Reservations, newRes)
				fmt.Println("Бронь успешно создана для", name)
				fmt.Println("Номер:", room.Number)
				fmt.Println("Ночей:", nights)
				fmt.Println("Сумма:", total)
				return
			} else {
				fmt.Println("Номер уже забронирован")
				return
			}
		}
	}
	fmt.Println("Номер не найден")
}

func (hotel *Hotel) ShowReservations() {
	if len(hotel.Reservations) == 0 {
		fmt.Println("Нет активных броней")
		return
	}
	for _, res := range hotel.Reservations {
		fmt.Println("Бронь ID:", res.ID, "Имя клиента:", res.Name, "Номер ID:", res.RoomID, "Сумма:", res.Total)
	}
}

func main() {
	myHotel := Hotel{Name: "Отель"}

	myHotel.AddRoom(Room{ID: 1, Number: "101", Capacity: 2, Price: 2500})
	myHotel.AddRoom(Room{ID: 2, Number: "102", Capacity: 3, Price: 3000})
	myHotel.AddRoom(Room{ID: 3, Number: "103", Capacity: 1, Price: 1800})

	myHotel.CreateReservation(1, "Иван Иванов", 3)
	fmt.Println()
	myHotel.CreateReservation(1, "Мария Петрова", 2)
	fmt.Println()
	myHotel.CreateReservation(2, "Анна Смирнова", 1)
	fmt.Println()
	myHotel.ShowReservations()
}
