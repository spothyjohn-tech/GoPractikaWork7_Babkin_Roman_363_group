package main

import (
	"fmt"
	"sync"
	"time"
)

type EventBus struct {
	handlers map[string][]func(interface{}) 
	mutex    sync.RWMutex
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()
	handlers, exists := eb.handlers[event]
	if !exists {
		fmt.Println("Данного события нет")
		return
	}
	for _, handler := range handlers {
		go handler(data)
	}
}

func main() {
	bus := EventBus{
		handlers: make(map[string][]func(interface{})),
	}
	bus.Subscribe("Купоны", func(data interface{}) {
		fmt.Println("Пришли новые купоны:", data)
	})
	bus.Subscribe("Регистрация", func(data interface{}) {
		fmt.Println("Вы успешно прошли регистрацию", data)
	})
	bus.Publish("Купоны", "пришёл новый скидочный купон")
	bus.Publish("Регистрация", "Иван Иванов")
	time.Sleep(1 * time.Second)
}
