## База данных

#### Создание:

1. ```createdb -U username db_name```
2. ```psql -U username -d db_name -f internal/db/schema.sql```

Где:
`username` - имя пользователя,
`order_service` - название базы данных

#### Для просмотра:

1. ```psql -U username```

2. ```\c db_name```

3. ```\dt```

**Результат:**

| Схема  | Имя      | Тип     | Владелец |
|--------|----------|---------|----------|
| public | delivery | таблица | postgres |
| public | items    | таблица | postgres |
| public | orders   | таблица | postgres |
| public | orders   | таблица | postgres |
