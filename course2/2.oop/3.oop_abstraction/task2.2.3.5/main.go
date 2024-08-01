package main

import (
	"fmt"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type Item struct {
	key   string
	value interface{}
	next  *Item // указатель на следующий элемент имеющий аналогичный хэш, в итоге получаем связный список
}

type MyHashMap struct {
	table    []*Item
	hashFunc func(data interface{}, size int) int
}

func MeasureTime(name string, f func()) {
	start := time.Now()

	f()

	time.Sleep(1 * time.Second)
	duration := (time.Since(start) - 1*time.Second).Nanoseconds()
	fmt.Printf("Execution time of %s: %v nanoseconds\n", name, duration)
}

func (m MyHashMap) Set(key string, value interface{}) {
	hash := m.hashFunc(key, len(m.table))

	item := Item{
		key:   key,
		value: value,
	}

	// Проверка наличия значения с таким хэшем
	// если оно имеется, то попытка подставить новое звено поле next
	// если его нет, то прсото подставляем новое значение
	if i := m.table[hash]; i != nil && i.key != key {
		for {
			if i.next == nil {
				i.next = &item
				break
			} else {
				i = i.next
			}
		}
	} else {
		m.table[hash] = &item
	}
}

func (m MyHashMap) Get(key string) (interface{}, bool) {
	hash := m.hashFunc(key, len(m.table))

	next := m.table[hash]

	// поиск нужного значения по ключу
	// в связном списке
	for next != nil {
		if next.key == key {
			break
		}

		next = next.next
	}

	isFound := next != nil

	return next, isFound
}

// Конструктор Хэш таблицы
func NewHashMap(size int, hashFunc func(data interface{}, size int) int) MyHashMap {
	return MyHashMap{
		table:    make([]*Item, size),
		hashFunc: hashFunc,
	}
}
