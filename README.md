# IT Purple Hack, Кейс: Платформа ценообразования Авито 
## Команда *ОПЗ (Объединение программистов здравомыслящих)*

## Что мы используем для решения?
- система контроля изменений для сохранения изменений цен 
- сервис отправки цен пользователям
- хранение данных в PostgreSQL
- Golang, использование библиотеки database/sql, каналов и горутин
- ReactJS для фронтенда

## Функционал приложения
1. Возможность добавления цены в матрицу
2. Создание матрицы
3. Редактирование множества столбцов
4. Получение цены

# Документация API

Это API предоставляет конечные точки для управления данными о ценах и таблицами. Оно позволяет добавлять, обновлять и извлекать информацию о ценах, создавать новые таблицы и получать имена таблиц.

## Конечные точки

- `POST /user/`: Получает информацию о цене для указанного пользователя(внешняя ручка).
- `POST /add`: Добавляет новые данные о цене.
- `PUT /update`: Обновляет существующие данные о цене.
- `PUT /update/many`: Обновляет несколько записей о ценах на основе процентов или фиксированной цены.
- `POST /create`: Создает новую таблицу.
- `GET /`: Получает имена всех доступных таблиц.

## Примечание

- Все конечные точки отвечают данными в формате JSON.
- Ответы на ошибки содержат соответствующие коды статуса и сообщения об ошибке.
- Убедитесь, что для каждой конечной точки используется правильный метод (POST, PUT, GET).
- Проверьте формат запроса и ответа для каждой конечной точки, чтобы обеспечить правильное использование.
- При запуске backend нужно создать папку в cmd\config и туда добавить .yaml файл со след структурой
- env: "local"
http_server:
  address: "localhost:8082"
  timeout: 2s
  idle_timeout: 60s
database:
  host: "localhost"
  port: DB_PORT
  name: DB_NAME
  username: DB_USERNAME
  password: DB_PASSWORD
