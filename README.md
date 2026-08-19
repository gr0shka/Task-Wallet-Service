# Task-Wallet-Service

Микросервис для пополнения баланса пользователя, реализованный по слоям (domain, repository, usecase/service, delivery/http).

## Задание

Тема: **Сервис баланса пользователя (Wallet Service)**.

Нужно реализовать микросервис с одним действием: **пополнение баланса пользователя**.

### Требования к бизнес-логике

- Пополнять можно только на сумму больше 0.
- Если пользователя с таким ID не существует, возвращать кастомную ошибку `ErrUserNotFound`.
- Если всё успешно, увеличить баланс пользователя, сохранить в хранилище и вернуть обновлённый баланс.

### Архитектурные ограничения

1. **Пакет `domain`:**
   - Структура `User` (`ID int64`, `Balance int64`).
   - Ошибка `var ErrUserNotFound = errors.New("user not found")`.
   - Интерфейс `UserRepository` (методы для юзкейса: получить по ID, обновить баланс).
   - Интерфейс `WalletUsecase` (метод `Deposit(ctx context.Context, userID int64, amount int64) (domain.User, error)`).

2. **Пакет `repository/memory`:**
   - Реализация `InMemoryUserRepository` на `map[int64]domain.User` и `sync.RWMutex`.
   - Репозиторий должен удовлетворять интерфейсу `domain.UserRepository`.

3. **Пакет `usecase`:**
   - Структура `walletUsecase`, принимающая `domain.UserRepository`.
   - Метод `Deposit(...)`, реализующий бизнес-правила (проверка `amount > 0`, обработка ошибок).

4. **Пакет `delivery/http`:**
   - HTTP-хэндлер на `net/http` (или `chi`) для `POST /wallet/deposit`.
   - Входящий JSON: `{"user_id": 1, "amount": 100}`.
   - Если `amount <= 0` → `400 Bad Request`.
   - Если пользователь не найден (`ErrUserNotFound`) → `404 Not Found`.
   - Если успех → `200 OK` и JSON с обновлённым пользователем.

5. **`main.go`:**
   - Создать in-memory репозиторий с дефолтным пользователем (`ID: 1`, `Balance: 1000`).
   - Создать usecase, передав туда репозиторий.
   - Создать HTTP-хэндлер, передав туда usecase.
   - Запустить сервер на порту `:8080`.
