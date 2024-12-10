[Как запустить?](https://github.com/DKDemerchyan/todo-list?tab=readme-ov-file#как-запустить-проект)

---
### Информация для ревьюера
Список задач со звездочкой \
`+` в проекте используются переменные из .env файла \
`+` реализованы правила повторения задач еженедельно и ежемесячно \
`+` реализован `Поиск` \
`+` создан `Dockerfile` \
`-` авторизация
---
# Описание API-Service Todo-List
Проект представляет собой планировщик задач, состоящий из web-сервера и базы данных. 

## Стандартный функционал работы с задачами
1. [x] Создание
2. [x] Редактирование
3. [x] Удаление
4. [x] Получение по id
5. [x] Получение всего списка
6. [x] Поиск по substr (регистронезависимый)
7. [x] Поиск по дате

### Паттерны повторения задач
- Если правило не указано, отмеченная выполненной задача будет удаляться из таблицы;
- Через каждые Х дней — задача переносится на указанное число дней;
- Ежегодно: при выполнении задачи дата перенесётся на год вперёд;
- Еженедельно: задачу можно назначить на указанные дни недели;
- Ежемесячно: задача назначается на указанные дни месяца, предпоследний или/и последний, можно также указать на какие
из месяцев может быть назначена задача;



------

## Как запустить проект:
1. Сделать Fork репозитория себе или клонировать этот \
`git clone git@github.com:DKDemerchyan/todo-list.git`


2. В директории проекта создать .env файл \
`touch .env` \
`nano .env` \
Записываем в него \
`TODO_PORT=7540` \
`TODO_DBFILE="scheduler.db"`


3. Запустить докер-контейнер \
`docker build -t todo:latest .` \
`docker run -d -p 7540:7540  todo:latest`

4. Открыть `localhost:7540`