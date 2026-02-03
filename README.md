# Meal Plan

[English](#english) | [Русский](#русский)

---

## English

A backend service for meal planning built with **Clean Architecture** principles. This project serves as a demonstration of writing idiomatic, maintainable, and production-ready Go code.

### 🚀 Key Features
- **Clean Architecture:** Strict separation between Transport, Business Logic, and Data Access layers.
- **Dependency Injection:** Components are decoupled via interfaces for better testability.
- **Context Management:** Proper use of `context.Context` for request timeouts and cancellations.
- **Database:** High-performance PostgreSQL interaction using `pgx`.
- **Validation:** Custom model validation logic without heavy external dependencies.

### 🛠 Tech Stack
- **Language:** Go 1.22+
- **Router:** [go-chi/chi](https://github.com/go-chi/chi)
- **Database:** PostgreSQL with [jackc/pgx](https://github.com/jackc/pgx)
- **Containerization:** Docker & Docker Compose

### 📂 Project Structure
```text
.
├── cmd/
│   └── main.go          # Entry point & Dependency Injection container
├── internal/
│   ├── handler/         # Transport layer (HTTP Handlers & Routing)
│   ├── service/         # Business logic layer
│   ├── repository/      # Data access layer (PostgreSQL implementation)
│   └── models/          # Domain entities & Validation logic
├── docker-compose.yaml
└── go.mod
```

### 🏁 Quick Start
- **Clone & Run:**

```Bash
git clone [https://github.com/Sarahttheory/meal-plan.git](https://github.com/Sarahttheory/meal-plan.git)
cd meal-plan
docker-compose up --build
```
- **Access API:**
The server will be available at http://localhost:8080

---

## Русский

Бэкенд-сервис для планирования рациона питания, разработанный с применением принципов Clean Architecture. Проект демонстрирует подход к написанию идиоматичного, поддерживаемого и готового к продакшену кода на Go.

### 🚀 Ключевые особенности
- **Clean Architecture:** Строгое разделение ответственности между слоями Transport, Business Logic и Data Access.
- **Dependency Injection:** Компоненты связаны через интерфейсы, что обеспечивает высокую тестируемость.
- **Context Management:** Корректное использование context.Context для управления таймаутами и отменой запросов.
- **Database:** Высокопроизводительное взаимодействие с PostgreSQL через драйвер pgx.
- **Validation:** Кастомная логика валидации моделей без использования тяжелых фреймворков.

### 🛠 Технологический стек
- **Язык:** Go 1.22+
- **Роутер:** [go-chi/chi](https://github.com/go-chi/chi)
- **База данных:** PostgreSQL + [jackc/pgx](https://github.com/jackc/pgx)
- **Контейнеризация:** Docker & Docker Compose

### 📂 Структура проекта
```text
.
├── cmd/
│   └── main.go          # Точка входа и инициализация зависимостей
├── internal/
│   ├── handler/         # Транспортный слой (HTTP обработчики)
│   ├── service/         # Слой бизнес-логики
│   ├── repository/      # Слой работы с БД
│   └── models/          # Доменные сущности и валидация
├── docker-compose.yaml
└── go.mod
```

### 🏁 Быстрый запуск
- **Клонирование и запуск:**

```Bash
git clone [https://github.com/Sarahttheory/meal-plan.git](https://github.com/Sarahttheory/meal-plan.git)
cd meal-plan
docker-compose up --build
```
- **API:**
Сервис будет доступен по адресу http://localhost:8080

---
