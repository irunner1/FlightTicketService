# Flight ticket service

---

## Docker

```bash
# Сборка
docker build -t airticketfinder .

# Запуск
docker run -it -p 8080:8080 --env-file=.env airticketfinder
docker run -d -p 8080:8080 airticketfinder

```

Для запуска необходимо подготовить файл переменных окружения (.env) c примерным содержанием:

```cmd
HOST=0.0.0.0
PORT=8080
```

+ HOST - хост запускаемой программы
+ PORT - порт запускаемой программы

---
