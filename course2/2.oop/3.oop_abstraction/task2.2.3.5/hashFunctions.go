package main

import (
	"hash/crc32"
	"hash/crc64"
)

func toBytes(data interface{}) []byte {
	// Проверяем, что data является строкой
	str, _ := data.(string)

	// Преобразуем строку в байтовый срез
	return []byte(str)
}

func CRC64() func(data interface{}, size int) int {
	return func(data interface{}, size int) int {
		bytes := toBytes(data)

		table := crc64.MakeTable(crc64.ECMA)
		checksum := crc64.Checksum(bytes, table) % uint64(size)
		return int(checksum)
	}
}

func CRC32() func(data interface{}, size int) int {
	return func(data interface{}, size int) int {
		bytes := toBytes(data)

		table := crc32.MakeTable(crc32.IEEE)
		checksum := crc32.Checksum(bytes, table) % uint32(size)
		return int(checksum)
	}
}

func CRC16() func(data interface{}, size int) int {
	return func(data interface{}, size int) int {
		bytes := toBytes(data)

		var polynomial uint16 = 0xA001
		var crc uint16 = 0xFFFF

		for _, b := range bytes {
			crc ^= uint16(b)
			for i := 0; i < 8; i++ {
				if (crc & 1) != 0 {
					crc = (crc >> 1) ^ polynomial
				} else {
					crc >>= 1
				}
			}
		}
		crc = crc % uint16(size)
		return int(crc)
	}
}
