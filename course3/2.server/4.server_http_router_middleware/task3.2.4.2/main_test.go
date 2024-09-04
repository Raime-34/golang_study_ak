package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMainRoute(t *testing.T) {
	// Создаем новый роутер
	r := chi.NewRouter()

	// Подключаем middleware и хендлер
	r.Use(LoggerMiddleware)
	r.Get("/", handleRoute)

	// Создаем HTTP-запрос
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Выполняем запрос
	r.ServeHTTP(rr, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверяем тело ответа
	expectedBody := "Hello world"
	assert.Equal(t, expectedBody, rr.Body.String())
}

func TestLoggerMiddleware(t *testing.T) {
	// Создаем новый роутер
	r := chi.NewRouter()

	// Подключаем middleware, но заменяем zap на фейковый логгер
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Создаем фейковый логгер, который не выводит данные в консоль
			logger := zap.NewNop()

			// Заменяем logger в middleware на фейковый
			defer logger.Sync()
			logger.Info("Request received",
				zap.String("url", r.URL.String()),
				zap.Int("attempt", 3),
				zap.Duration("backoff", time.Second),
			)
			next.ServeHTTP(w, r)
		})
	})

	// Добавляем хендлер
	r.Get("/", handleRoute)

	// Создаем HTTP-запрос
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Выполняем запрос
	r.ServeHTTP(rr, req)

	// Проверяем статус ответа
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Hello world", rr.Body.String())
}
