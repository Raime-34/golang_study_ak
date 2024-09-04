package main

import (
	"bytes"
	"testing"
)

func TestMyHashMap_CRC64(t *testing.T) {
	MeasureTime("CRC64", func() {
		hashMap := NewHashMap(10, CRC64())

		// Вставляем данные в хэш-таблицу
		hashMap.Set("key1", "value1")
		hashMap.Set("key2", "value2")
		hashMap.Set("key3", "value3")

		// Проверяем корректность данных
		value, found := hashMap.Get("key1")
		if !found || value.(*Item).value != "value1" {
			t.Errorf("Expected value1, got %v", value)
		}

		value, found = hashMap.Get("key2")
		if !found || value.(*Item).value != "value2" {
			t.Errorf("Expected value2, got %v", value)
		}

		value, found = hashMap.Get("key3")
		if !found || value.(*Item).value != "value3" {
			t.Errorf("Expected value3, got %v", value)
		}

		// Проверяем случай, когда ключа нет
		value, found = hashMap.Get("key4")
		if found || value.(*Item) != nil {
			t.Errorf("Expected no value, got %v", value)
		}
	})
}

// Тестирование функции Set и Get с использованием хэш-функции CRC32
func TestMyHashMap_CRC32(t *testing.T) {
	MeasureTime("CRC32", func() {
		hashMap := NewHashMap(10, CRC32())

		// Вставляем данные в хэш-таблицу
		hashMap.Set("key1", "value1")
		hashMap.Set("key2", "value2")
		hashMap.Set("key3", "value3")

		// Проверяем корректность данных
		value, found := hashMap.Get("key1")
		if !found || value.(*Item).value != "value1" {
			t.Errorf("Expected value1, got %v", value)
		}

		value, found = hashMap.Get("key2")
		if !found || value.(*Item).value != "value2" {
			t.Errorf("Expected value2, got %v", value)
		}

		value, found = hashMap.Get("key3")
		if !found || value.(*Item).value != "value3" {
			t.Errorf("Expected value3, got %v", value)
		}

		// Проверяем случай, когда ключа нет
		value, found = hashMap.Get("key4")
		if found || value.(*Item) != nil {
			t.Errorf("Expected no value, got %v", value)
		}
	})
}

// Тестирование функции Set и Get с использованием хэш-функции CRC16
func TestMyHashMap_CRC16(t *testing.T) {
	MeasureTime("CRC16", func() {
		hashMap := NewHashMap(10, CRC16())

		// Вставляем данные в хэш-таблицу
		hashMap.Set("key1", "value1")
		hashMap.Set("key2", "value2")
		hashMap.Set("key3", "value3")

		// Проверяем корректность данных
		value, found := hashMap.Get("key1")
		if !found || value.(*Item).value != "value1" {
			t.Errorf("Expected value1, got %v", value)
		}

		value, found = hashMap.Get("key2")
		if !found || value.(*Item).value != "value2" {
			t.Errorf("Expected value2, got %v", value)
		}

		value, found = hashMap.Get("key3")
		if !found || value.(*Item).value != "value3" {
			t.Errorf("Expected value3, got %v", value)
		}

		// Проверяем случай, когда ключа нет
		value, found = hashMap.Get("key4")
		if found || value.(*Item) != nil {
			t.Errorf("Expected no value, got %v", value)
		}
	})
}

// Тестирование разрешения коллизий
func TestMyHashMap_Collision(t *testing.T) {
	// Используем маленький размер хэш-таблицы, чтобы спровоцировать коллизии
	hashMap := NewHashMap(2, CRC32())

	hashMap.Set("key1", "value1")
	hashMap.Set("key2", "value2") // Эта запись вероятно столкнется с key1
	hashMap.Set("key3", "value3")

	// Проверяем, что все элементы правильно хранятся
	value, found := hashMap.Get("key1")
	if !found || value.(*Item).value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	value, found = hashMap.Get("key2")
	if !found || value.(*Item).value != "value2" {
		t.Errorf("Expected value2, got %v", value)
	}

	value, found = hashMap.Get("key3")
	if !found || value.(*Item).value != "value3" {
		t.Errorf("Expected value3, got %v", value)
	}
}

func TestToBytes(t *testing.T) {
	// Тестирование корректной работы с строками
	str := "test string"
	expected := []byte(str)
	output := toBytes(str)
	if !bytes.Equal(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func BenchmarkCRC64(b *testing.B) {
	hashFunc := CRC64()
	data := "some test data"
	size := 100

	for i := 0; i < b.N; i++ {
		hashFunc(data, size)
	}
}

func BenchmarkCRC32(b *testing.B) {
	hashFunc := CRC32()
	data := "some test data"
	size := 100

	for i := 0; i < b.N; i++ {
		hashFunc(data, size)
	}
}

func BenchmarkCRC16(b *testing.B) {
	hashFunc := CRC16()
	data := "some test data"
	size := 100

	for i := 0; i < b.N; i++ {
		hashFunc(data, size)
	}
}
