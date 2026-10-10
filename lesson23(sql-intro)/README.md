# Урок 23 — Введение в SQL и подключение к БД

Шпаргалка для лайв-кодинга. Полная теория, тайминг и практика — в `Urok_22_Vvedenie_v_SQL.md`.

## Главные аналогии (сказать в начале)
- База данных — **склад с идеальным кладовщиком**: «дай все ноутбуки дешевле 1000, отсортируй по цене» — приносит за миллисекунды. Файл — коробка на чердаке: чтобы что-то найти, надо перебрать всё вручную.
- `sql.DB` — это **пул соединений**, а не одно соединение. Как стойка с телефонами: берёшь свободный, поговорил — вернул на место.
- Плейсхолдер `$1` — **бланк с пустой графой**: текст запроса печатается заранее, а данные вписываются отдельно, и «превратиться в команду» они уже не могут.

## Подготовка (до урока)
1. Запустить PostgreSQL (вариант с Docker):
   ```bash
   docker run --name pg -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=coursedb \
     -p 5432:5432 -d postgres:16
   ```
2. Проверить подключение: `psql -h localhost -U postgres -d coursedb` (пароль `secret`).
3. В папке урока уже лежат `go.mod` и `go.sum` с драйвером `lib/pq`. Если хотите показать студентам установку вживую — временно удалите их и выполните:
   ```bash
   go mod init dbdemo
   go get github.com/lib/pq
   ```
4. Сброс базы к началу (если что-то сломали или повторяете урок): файл `reset.sql`
   ```bash
   psql -h localhost -U postgres -d coursedb -f reset.sql
   # или в Docker:
   docker exec -i pg psql -U postgres -d coursedb < reset.sql
   ```

## Как запускать
Шаги 1–2 — SQL-файлы, команды копируем в `psql` по одной. Шаги 3–11 — Go-программы:
```bash
cd "lesson23(sql-intro)"
go run ./05-query-mnogo-strok
# или из папки шага: cd 05-query-mnogo-strok && go run main.go
```
Строка подключения во всех шагах одна: `host=localhost port=5432 user=postgres password=secret dbname=coursedb sslmode=disable`.

## Порядок шагов
| # | Папка | Что показываем |
|---|---|---|
| 1 | `01-sozdaem-tablicu` | `CREATE TABLE`, `INSERT`. Ограничения: UNIQUE и CHECK не пускают плохие данные. |
| 2 | `02-vybiraem-dannye` | `SELECT`: `WHERE`, `ORDER BY`, `LIMIT`, `COUNT/AVG`, `LIKE`, `BETWEEN`. |
| 3 | `03-podklyuchenie-k-baze` | `sql.Open` + `Ping`. Драйвер с `_`. Что будет при неверном пароле. |
| 4 | `04-nastroika-pula` | `SetMaxOpenConns` и другие, `db.Stats()`. `Open` ничего не открывает. |
| 5 | `05-query-mnogo-strok` | `db.Query` — шаблон из 5 пунктов: `Query → defer Close → Next → Scan → Err`. |
| 6 | `06-queryrow-odna-stroka` | `db.QueryRow` и `sql.ErrNoRows` («не нашлось» — это не поломка). |
| 7 | `07-parametry-i-inekcii` | `$1` против склейки строк. Живая SQL-инъекция `' OR '1'='1`. |
| 8 | `08-agregaty-i-null` | `COUNT/AVG/MIN/MAX`, `sql.NullFloat64`, `.Valid`. |
| 9 | `09-dinamicheskii-filtr` | Фильтр на указателях: запрос собирается из частей, значения — только через `args`. |
| 10 | `10-konsolnyi-prosmotrshchik` | Мини-проект: меню + база. Всё вместе. |
| 11 | `11-tipichnye-oshibki` | Ломаем нарочно: порядок в `Scan`, NULL в `float64`, забытый `rows.Close()`. |

Что здесь **нарочно сломано** — папка `11-tipichnye-oshibki` (и «плохая» функция `findByNameUnsafe` в шаге 7). Не исправляйте.

## Частые ошибки новичков
| Ошибка | Симптом | Решение |
|---|---|---|
| Нет `_` перед драйвером | `unknown driver "postgres"` | `_ "github.com/lib/pq"` |
| Забыт `rows.Close()` | Утечка соединений, зависание | `defer rows.Close()` сразу после `Query` |
| Порядок в `Scan` | `Scan error ... converting ...` | Строго как в `SELECT` |
| NULL в обычный тип | `converting NULL to float64 is unsupported` | `sql.NullXxx` или указатель `*T` |
| Склейка SQL строками | SQL-инъекция | Плейсхолдеры `$1` |
| Нет `rows.Err()` | Тихая потеря данных | Проверять после цикла |
| `sql.Open` в каждой функции | Куча пулов | Один `*sql.DB` на приложение |
| `sslmode` не указан | `connection requires SSL` | `sslmode=disable` локально |

## Вопросы группе в конце
1. Почему БД лучше файла для поиска?
2. Что делает `sql.Open` и почему нужен `Ping`?
3. Зачем `_` перед импортом драйвера?
4. Что такое SQL-инъекция и как Go от неё защищает?
5. Чем `Query` отличается от `QueryRow` и `Exec`?
