package storages

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

// Хранилище в оперативной памяти
type StoregeRam struct {
	lock sync.Mutex
	data map[string]int
}

func CreateStoregeRam () *StoregeRam {
	return &StoregeRam{data: map[string]int{}, lock: sync.Mutex{}}
}

// Получить данные из хранилища
func (s *StoregeRam) GetValue(key string) int {
	s.lock.Lock()
	log.Println("Хранилище заблокировано для Чтения")
	log.Println(s.data)
	defer s.lock.Unlock()
	return s.data[key]
}

// Записать данные в хранилище
func (s *StoregeRam) SetValue(key string, value int) {
	s.lock.Lock()
	log.Println("Хранилище заблокировано для записи")
	defer s.lock.Unlock()
	s.data[key] = value 
}

// Удаляет данные из хранилища
func (s *StoregeRam) UnsetValue(key string) {
	s.lock.Lock()
	log.Println("Хранилище заблокировано для Удаления")
	defer s.lock.Unlock()
	delete(s.data, key)
	log.Println("Хранилище заблокировано для Удаления ключа", key)
}

func TestSet (storage *StoregeRam) {
	i := 0
	for {
		value, err := fmt.Printf("%d!", i)
		if err != nil {
			continue
		}
		key := strconv.Itoa(i)
		if err != nil {
			continue
		} else {
			storage.SetValue(key, value)
			i += 1
		}
		time.Sleep(time.Second * 1)
	}
}

func TestGet (storage *StoregeRam) {
	i := 0
	for {
		key := strconv.Itoa(i)
		log.Println(storage.GetValue(key))
		i += 1
		time.Sleep(time.Second * 1)
	}
}

func TestUnset (storage *StoregeRam) {
	i := 0
	for {
		key := strconv.Itoa(i)
		storage.UnsetValue(key)
		i += 1
		time.Sleep(time.Second * 2)
	}
}