[![integration-tests](https://github.com/belamov/ypgo-url-shortener/actions/workflows/shortenertest.yml/badge.svg)](https://github.com/belamov/ypgo-url-shortener/actions/workflows/shortenertest.yml)
[![unit tests](https://github.com/belamov/ypgo-url-shortener/actions/workflows/tests.yml/badge.svg)](https://github.com/belamov/ypgo-url-shortener/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/belamov/ypgo-url-shortener/badge.svg?branch=main)](https://coveralls.io/github/belamov/ypgo-url-shortener?branch=main)
[![lint](https://github.com/belamov/ypgo-url-shortener/actions/workflows/lint.yml/badge.svg)](https://github.com/belamov/ypgo-url-shortener/actions/workflows/golangci-lint.yml)
[![go vet test](https://github.com/belamov/ypgo-url-shortener/actions/workflows/statictest.yml/badge.svg)](https://github.com/belamov/ypgo-url-shortener/actions/workflows/statictest.yml)

## Практические задания на курсе [продвинутый go-разработчик](https://practicum.yandex.ru/go-advanced/)

практика была оформлена в виде инкрементов, каждый инкремент открывался последовательно

## Список инкрементов

<details>
  <summary>Инкремент 1</summary>

Напишите сервис для сокращения длинных URL. Требования:
* Сервер должен быть доступен по адресу: `http://localhost:8080`.
* Сервер должен предоставлять два эндпоинта: `POST /` и `GET /{id}`.
* Эндпоинт `POST /` принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом `201` и сокращённым URL в виде текстовой строки в теле.
* Эндпоинт `GET /{id}` принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом `307` и оригинальным URL в HTTP-заголовке Location.
* Нужно учесть некорректные запросы и возвращать для них ответ с кодом `400`.

</details>

<details>
  <summary>Инкремент 2</summary>

 Покройте сервис юнит-тестами. Сконцентрируйтесь на покрытии тестами эндпоинтов, чтобы защитить API сервиса от случайных изменений.

</details>

<details>
  <summary>Инкремент 3</summary>

 Вы написали приложение с помощью стандартной библиотеки `net/http`. Используя любой пакет (роутер или фреймворк), совместимый с net/http, перепишите ваш код.
Задача направлена на рефакторинг приложения с помощью готовой библиотеки.

Обратите внимание, что необязательно запускать приложение вручную: тесты, которые вы написали до этого, помогут вам в рефакторинге.

</details>

<details>
  <summary>Инкремент 4</summary>

Добавьте в сервер новый эндпоинт `POST /api/shorten`, принимающий в теле запроса JSON-объект `{"url":"<some_url>"}` и возвращающий в ответ объект `{"result":"<shorten_url>"}`.

Не забудьте добавить тесты на новый эндпоинт, как и на предыдущие.

Помните про HTTP content negotiation, проставляйте правильные значения в заголовок `Content-Type`.

</details>

<details>
  <summary>Инкремент 5</summary>

Добавьте возможность конфигурировать сервис с помощью переменных окружения:
* адрес запуска HTTP-сервера с помощью переменной `SERVER_ADDRESS`.
* базовый адрес результирующего сокращённого URL с помощью переменной `BASE_URL`.
</details>

<details>
  <summary>Инкремент 6</summary>

Сохраняйте все сокращённые URL на диск в виде файла. При перезапуске приложения все URL должны быть восстановлены.

Путь до файла должен передаваться в переменной окружения `FILE_STORAGE_PATH`.

При отсутствии переменной окружения или при её пустом значении вернитесь к хранению сокращённых URL в памяти.
</details>

<details>
  <summary>Инкремент 7</summary>

Поддержите конфигурирование сервиса с помощью флагов командной строки наравне с уже имеющимися переменными окружения:
* флаг `-a`, отвечающий за адрес запуска HTTP-сервера (переменная `SERVER_ADDRESS`);
* флаг `-b`, отвечающий за базовый адрес результирующего сокращённого URL (переменная `BASE_URL`);
* флаг `-f`, отвечающий за путь до файла с сокращёнными URL (переменная `FILE_STORAGE_PATH`).

Во всех случаях должны быть:
* значения по умолчанию,
* приоритет значений, полученных через ENV, перед значениями, задаваемыми посредством флагов.
</details>

<details>
  <summary>Инкремент 8</summary>

Добавьте поддержку `gzip` в ваш сервис. Научите его:
* принимать запросы в сжатом формате (HTTP-заголовок `Content-Encoding`);
* отдавать сжатый ответ клиенту, который поддерживает обработку сжатых ответов (HTTP-заголовок `Accept-Encoding`).

Вспомните middleware из урока про HTTP-сервер, это может вам помочь.
</details>

<details>
  <summary>Инкремент 9</summary>

Добавьте в сервис функциональность аутентификации пользователя.

Сервис должен:
* Выдавать пользователю симметрично подписанную куку, содержащую уникальный идентификатор пользователя, если такой куки не существует или она не проходит проверку подлинности.
* Иметь хендлер `GET /api/user/urls`, который сможет вернуть пользователю все когда-либо сокращённые им URL в формате:
```
[
    {
        "short_url": "http://...",
        "original_url": "http://..."
    },
    ...
]
```
При отсутствии сокращённых пользователем URL хендлер должен отдавать HTTP-статус `204 No Content`.
Получить куки запроса можно из поля `(*http.Request).Cookie`, а установить — методом `http.SetCookie`.
</details>

<details>
  <summary>Инкремент 10</summary>

* Добавьте в сервис функциональность подключения к базе данных. В качестве СУБД используйте PostgreSQL не ниже 10 версии.
* Добавьте в сервис хендлер `GET /ping`, который при запросе проверяет соединение с базой данных. При успешной проверке хендлер должен вернуть HTTP-статус `200 OK`, при неуспешной — `500 Internal Server Error`.
* Строка с адресом подключения к БД должна получаться из переменной окружения `DATABASE_DSN` или флага командной строки `-d`.
</details>

<details>
  <summary>Инкремент 11</summary>

Перепишите сервис так, чтобы СУБД PostgreSQL стала хранилищем сокращённых URL вместо текущей реализации.

Сервису нужно самостоятельно создать все необходимые таблицы в базе данных. Схема и формат хранения остаются на ваше усмотрение.

При отсутствии переменной окружения `DATABASE_DSN` или флага командной строки `-d` или при их пустых значениях вернитесь последовательно к:
* хранению сокращённых URL в файле при наличии соответствующей переменной окружения или флага командной строки;
* хранению сокращённых URL в памяти.
</details>

<details>
  <summary>Инкремент 12</summary>

Добавьте новый хендлер `POST /api/shorten/batch`, принимающий в теле запроса множество URL для сокращения в формате:
```
[
    {
        "correlation_id": "<строковый идентификатор>",
        "original_url": "<URL для сокращения>"
    },
    ...
]
```

В качестве ответа хендлер должен возвращать данные в формате:

```
[
    {
        "correlation_id": "<строковый идентификатор из объекта запроса>",
        "short_url": "<результирующий сокращённый URL>"
    },
    ...
]
```
</details>

<details>
  <summary>Инкремент 13</summary>

Сделайте в таблице базы данных с сокращёнными URL уникальный индекс для поля с исходным URL. Это позволит избавиться от дублирующих записей в базе данных.

При попытке пользователя сократить уже имеющийся в базе URL через хендлеры `POST /` и POST `/api/shorten` сервис должен вернуть HTTP-статус `409 Conflict`, а в теле ответа — уже имеющийся сокращённый URL в правильном для хендлера формате.

Стратегии реализации:
1. Чтобы не проверять наличие оригинального URL в базе данных отдельным запросом, можно воспользоваться конструкцией `INSERT ... ON CONFLICT` в PostgreSQL. Однако в этом случае придётся самостоятельно возвращать и проверять собственную ошибку.
2. Чтобы определить тип ошибки PostgreSQL, с которой завершился запрос, можно воспользоваться библиотекой `github.com/jackc/pgerrcode`, в частности `pgerrcode.UniqueViolation`. В таком случае нужно будет сделать дополнительный запрос к хранилищу, чтобы определить сокращённый вариант URL.
</details>

<details>
  <summary>Инкремент 14</summary>

Сделайте в таблице базы данных с сокращёнными URL дополнительное поле с флагом, указывающим на то, что URL должен считаться удалённым.

Далее добавьте в сервис новый асинхронный хендлер `DELETE /api/user/urls`, который принимает список идентификаторов сокращённых URL для удаления в формате:

```
[ "a", "b", "c", "d", ...]
```

В случае успешного приёма запроса хендлер должен возвращать HTTP-статус `202 Accepted`. Фактический результат удаления может происходить позже — каким-либо образом оповещать пользователя об успешности или неуспешности не нужно.

Успешно удалить URL может пользователь, его создавший. При запросе удалённого URL с помощью хендлера `GET /{id}` нужно вернуть статус `410 Gone`.

Совет:
* Для эффективного проставления флага удаления в базе данных используйте множественное обновление (batch update).
* Используйте паттерн fanIn для максимального наполнения буфера объектов обновления.
</details>

<details>
  <summary>Инкремент 15</summary>

Добавьте в свой проект бенчмарки, измеряющие скорость выполнения важнейших компонентов вашей системы.

Проведите анализ использования памяти вашим проектом, определите и исправьте неэффективные части кода по следующему алгоритму:
1. Используя профилировщик `pprof`, сохраните профиль потребления памяти вашим проектом в директорию `profiles` с именем `base.pprof`.
2. Изучите полученный профиль, определите и исправьте неэффективные части вашего кода.
3. Повторите пункт 1 и сохраните новый профиль потребления памяти в директорию `profiles` с именем `result.pprof`.
4. 
Проверьте результат внесённых изменений командой:
```
pprof -top -diff_base=profiles/base.pprof profiles/result.pprof
```

В случае успешной оптимизации вы увидите в выводе командной строки результаты с отрицательными значениями, означающими уменьшение потребления ресурсов.

Внимание: к концу текущего спринта покрытие вашего кода автотестами должно быть не менее 40%.
</details>

<details>
  <summary>Инкремент 16</summary>

Отформатируйте свой проект с помощью `gofmt` или `goimports`. Убедитесь, что все файлы проекта прошли форматирование.

Внимание: к концу текущего спринта покрытие вашего кода автотестами должно быть не менее 40%.
</details>

<details>
  <summary>Инкремент 17</summary>

Добавьте к основным экспортированным методам и переменным (хендлерам, публичным структурам и интерфейсам) в вашем проекте документацию в формате godoc.

Добавьте примеры работы с эндпоинтами практического трека в формате `example_test.go`.

Покрытие вашего кода тестами на данный момент должно быть не менее 40%.
</details>

<details>
  <summary>Инкремент 18</summary>

Создайте свой `multichecker`, состоящий из:
* стандартных статических анализаторов пакета `golang.org/x/tools/go/analysis/passes`;
* всех анализаторов класса `SA` пакета `staticcheck.io`;
* не менее одного анализатора остальных классов пакета `staticcheck.io`;
* двух или более любых публичных анализаторов на ваш выбор.

Напишите и добавьте в `multichecker` собственный анализатор, запрещающий использовать прямой вызов `os.Exit` в функции `main` пакета `main`. При необходимости перепишите код своего проекта так, чтобы он удовлетворял данному анализатору.

Поместите анализатор в директорию `cmd/staticlint` вашего проекта. Добавьте документацию в формате `godoc`, подробно опишите в ней механизм запуска `multichecker`, а также каждый анализатор и его назначение.

Исходный код вашего проекта должен проходить статический анализ созданного `multichecker`.

Покрытие вашего кода автотестами к концу спринта должно быть не менее 55%.
</details>

<details>
  <summary>Инкремент 19</summary>

Добавьте в пакет `cmd/shortener` глобальные переменные:
* `var buildVersion string`,
* `var buildDate string`,
* `var buildCommit string`.

При старте приложения выводите в `stdout` сообщение в следующем формате:

```
Build version: <buildVersion> (или "N/A" при отсутствии значения)
Build date: <buildDate> (или "N/A" при отсутствии значения)
Build commit: <buildCommit> (или "N/A" при отсутствии значения)
```

Покрытие вашего кода автотестами на данный момент должно быть не менее 55%.
</details>

<details>
  <summary>Инкремент 20</summary>

Добавьте в свой код возможность включения `HTTPS` в веб-сервере.

При передаче флага `-s` или переменной окружения `ENABLE_HTTPS` запускайте сервер с помощью метода `http.ListenAndServeTLS` или `tls.Listen`.
</details>

<details>
  <summary>Инкремент 21</summary>

Добавьте возможность конфигурации приложения с помощью файла в формате `JSON`. Нужно поддержать все действующие опции приложения. Имя файла конфигурации должно задаваться через флаг `-c/-config` или переменную окружения `CONFIG`. Значения из файла конфигурации должны иметь меньший приоритет, чем флаги или переменные окружения.

Формат файла:
```
{
    "server_address": "localhost:8080", // аналог переменной окружения SERVER_ADDRESS или флага -a
    "base_url": "http://localhost", // аналог переменной окружения BASE_URL или флага -b
    "file_storage_path": "/path/to/file.db", // аналог переменной окружения FILE_STORAGE_PATH или флага -f
    "database_dsn": "", // аналог переменной окружения DATABASE_DSN или флага -d
    "enable_https": true // аналог переменной окружения ENABLE_HTTPS или флага -s
}
```
</details>

<details>
  <summary>Инкремент 22</summary>

Сервер должен штатно завершаться по сигналам: `syscall.SIGTERM`, `syscall.SIGINT`, `syscall.SIGQUIT`.

Все необработанные запросы должны быть обработаны до конца, все несохранённые данные должны быть сохранены в хранилище.
</details>

<details>
  <summary>Инкремент 23</summary>

Добавьте в сервер новый эндпоинт `GET /api/internal/stats`, возвращающий в ответ объект:
```
{
  "urls": <int>, // количество сокращённых URL в сервисе
  "users": <int> // количество пользователей в сервисе
}
```

Добавьте в конфигурационный JSON-файл HTTP-сервера поле `trusted_subnet` (тип `string`, переменная окружения `TRUSTED_SUBNET`, флаг `-t`), в которое можно передать строковое представление бесклассовой адресации (CIDR).

При запросе эндпоинта `/api/internal/stats` нужно проверять, что переданный в заголовке запроса `X-Real-IP` IP-адрес клиента входит в доверенную подсеть, в противном случае возвращать статус ответа `403 Forbidden`.

При пустом значении переменной `trusted_subnet` доступ к эндпоинту должен быть запрещён для любого входящего запроса.
</details>

<details>
  <summary>Инкремент 24</summary>

Добавьте возможность делать запросы к вашему серверу по протоколу gRPC. Все хендлеры должны быть доступны по gRPC и функционировать настолько идентично уже имеющимся, насколько это возможно в рамках протокола. Поддержите все возможности конфигурации gRPC-сервера аналогично конфигурациям HTTP-сервера.

Совет: попробуйте сделать HTTP- и gRPC-хендлеры фасадами к общему коду с бизнес-логикой.
</details>
