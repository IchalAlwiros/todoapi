# CRUD API CASE TODO APP
In this project I tried to use the Golang programming language to create a client API service to perform CRUD operations
This API service is supported by using a **_`Supabase`_** database and deploying it using **_`Railway`_**

> [!NOTE]
> + **`GET`** : https://todoapi-production.up.railway.app/todo
> + **`GET`** : https://todoapi-production.up.railway.app/todo/:id
> + **`POST`** : https://todoapi-production.up.railway.app/todo
> + **`PATCH`** : https://todoapi-production.up.railway.app/todo/:id
> + **`PUT`** : https://todoapi-production.up.railway.app/todo/:id
> + **`DELETE`** : https://todoapi-production.up.railway.app/todo/:id


**`GET`** Get All item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo`
> [!NOTE]
> + you can use pagination with query params
> +`https://todoapi-production.up.railway.app/todo?page=1&page_limit=2`

```javascript
    Content-Type: application/json
    {
        "data": [
          {
              "Id": 1,
              "Title": "Title Todo",
              "Note": "I Learn Golang Today",
              "CreatedAt": "2024-03-27T03:19:01.966964Z",
              "UpdatedAt": "2024-03-27T03:19:01.966964Z"
          }
        ],
        "success": true
   }
```

**`POST`** Create new item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo`<br>
Request Body:
```javascript
    Content-Type: application/json
    {
      "title" : "Judul 2",
      "note" : "Hari ini saya ingin membuat API"
    }
```
Response Body - HTTP Response Code: **200**
```javascript
    Content-Type: application/json
    {
      "data": {
        "Id": 4,
        "Title": "Judul 2",
        "Note": "Hari ini saya ingin membuat API",
        "CreatedAt": "2024-03-29T05:57:45.259327828Z",
        "UpdatedAt": "2024-03-29T05:57:45.259327828Z"
      },
      "success": true
    }
```


**`GET`** Get One item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo/:id`
> [!NOTE]
> + take only 1 data from a certain ID
> +`https://todoapi-production.up.railway.app/todo/2`

```javascript
    Content-Type: application/json
    {
      "data": {
        "Id": 2,
        "Title": "Judul 2",
        "Note": "Hari ini saya ingin membuat API",
        "CreatedAt": "2024-03-27T03:22:48.82648Z",
        "UpdatedAt": "2024-03-29T06:06:56.600493703Z"
      },
      "success": true
    }
```


**`PATCH`** Save edit item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo/:id`<br>
> [!NOTE]
> + can be used to edit certain fields, for example here I only want to change the title
> +`https://todoapi-production.up.railway.app/todo/4`
Request Body:
```javascript
    Content-Type: application/json
    {
      "title" : "Judul Baru",
    }
```
Response Body - HTTP Response Code: **200**
```javascript
    Content-Type: application/json
    {
      "data": {
        "Id": 4,
        "Title": "Judul Baru",
        "Note": "Hari ini saya ingin membuat API",
        "CreatedAt": "2024-03-29T05:57:45.259327828Z",
        "UpdatedAt": "2024-03-29T05:57:45.259327828Z"
      },
      "success": true
    }
```

**`PATCH`** Save edit item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo/:id`<br>
> [!NOTE]
> + put involves all fields in the request body to change one piece of data
> +`https://todoapi-production.up.railway.app/todo/4`
Request Body:
```javascript
    Content-Type: application/json
    {
      "title" : "Judul Baru",
      "Note": "Hari ini saya ingin membuat API Baru",
    }
```
Response Body - HTTP Response Code: **200**
```javascript
    Content-Type: application/json
    {
      "data": {
        "Id": 4,
        "Title": "Judul Baru",
        "Note": "Hari ini saya ingin membuat API Baru",
        "CreatedAt": "2024-03-29T05:57:45.259327828Z",
        "UpdatedAt": "2024-03-29T05:57:45.259327828Z"
      },
      "success": true
    }
```


**`DELETE`** Get One item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo/:id`
> [!NOTE]
> + take only 1 data from a certain ID
> +`https://todoapi-production.up.railway.app/todo/2`

```javascript
    Content-Type: application/json
    {
      "data": {
        "Id": 2,
        "Title": "Judul 2",
        "Note": "Hari ini saya ingin membuat API",
        "CreatedAt": "2024-03-27T03:22:48.82648Z",
        "UpdatedAt": "2024-03-29T06:06:56.600493703Z"
      },
      "success": true
    }
```

if error
```javascript
     Content-Type: application/json
    {
      "message": "Record not found!",
      "success": false
    }
```



