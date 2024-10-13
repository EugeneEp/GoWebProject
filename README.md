# GoWebProject

### Введение:
**GoWebProject** - Небольшой Rest Api Token Based проект, написанный на Golang, с использованием фреймворка Gin [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)


### Описание
Проект представляет из себя веб сервер с базой данных, с возможностью выполнения CRUD операций над сущностью User.  
А так же с аутентификацией пользователя посредством передачи Bearer токена через HTTP заголовки.

### Архитектура проекта
При разработке была использована "чистая архитектура"

### Реализованный функционал:

#### API:

### Auth

**[POST]**

*/api/v1/auth* - получить JWT токен **<API_TOKEN>**
JSON 
```
{
  `username`:"admin",
  `password`:"admin"
}
```

### Users
Headers
Authorization: **Bearer <API_TOKEN> (required)**

**[GET]**

*/api/v1/users* - получить всех пользователей
*/api/v1/users/{id}* - получить пользователя по id  

**[POST]**

*/api/v1/users* - создать пользователя  
JSON 
```
{
  `username`:"test123",
  `password`:"test123",
  `email`:"test123@mail.ru"
}
```

**[PATCH]**

*/api/v1/users/{id}* - обновить данные пользователя по id
JSON 
```
{
  `username`:"test321",
  `password`:"test321"
}
```
-----

### Docker Compose
> sudo docker-compose up --build
