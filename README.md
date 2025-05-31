# quotes_service


Мини-сервис “Цитатник”

- Добавление цитаты
- Получение всех цитат
- Получение случайной цитаты
- Поиск цитат по автору
- Удаление цитаты по ID

# Стек
- Go 1.21+
- Gorilla Mux

# Перед запуском установите зависимости:
go mod tidy
#Запуск из корня проекта командой:
go run cmd/main.go


# Примеры использования:

* Добавление цитаты
  curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author": "Confucius", "quote": "Life is simple, but we insist on making it complicated."}'


* Получение всех цитат
  curl http://localhost:8080/quotes


* Получение цитат по автору
  curl "http://localhost:8080/quotes?author=Confucius"

  
* Получение случайной цитаты
  curl http://localhost:8080/quotes/random 


# Удаление цитаты
curl -X DELETE http://localhost:8080/quotes/1


