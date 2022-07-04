package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//Создаем мап
	stringMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
	}

	//Создаем объект структуры с mutex
	mapStructInstance := MapStruct{
		m: stringMap,
	}

	rand.Seed(time.Now().UnixNano())
	//Запускаем три горутины с указателем на наш объект
	for i := 0; i < 3; i++ {
		go Worker(&mapStructInstance)
	}

	for {

	}
}

//Так как идет и чтение и запись используем RWMutex
type MapStruct struct {
	mu sync.RWMutex
	m  map[int]int
}

// Иногда нужно чтобы все читали, но когда приходит поток который пишет, он ждет когда читатели уйдут и лочит, отрабатывает и потом снова можно читать
func (m *MapStruct) WriteToMap(key int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.m[key]++
}

//Просто записывает в рандомное место
func Worker(mapStruct *MapStruct) {
	for {
		key := rand.Intn(3)
		mapStruct.WriteToMap(key)
		fmt.Println(key, "  ", mapStruct.m[key])
		time.Sleep(500 * time.Millisecond)
	}
}
