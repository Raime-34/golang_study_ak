package task2_5_2_1

import "testing"

func BenchmarkJSON(b *testing.B) {
	users := GenerateUSER(1000)

	for i := 0; i < b.N; i++ {
		JSON(users)
	}
}

func BenchmarkEasyJSON(b *testing.B) {
	users := GenerateUSER(1000)

	for i := 0; i < b.N; i++ {
		EasyJSON(users)
	}
}

func BenchmarkJSONiter(b *testing.B) {
	users := GenerateUSER(1000)

	for i := 0; i < b.N; i++ {
		JSONiter(users)
	}
}
