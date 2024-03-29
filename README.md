# CRUD API CASE TODO APP
In this project I tried to use the Golang programming language to create a client API service to perform CRUD operations
This API service is supported by using a **_`Supabase`_** database and deploying it using **_`Railway`_**

> [!NOTE]
> + **`GET`** : https://todoapi-production.up.railway.app/todo
> + **`POST`** : https://todoapi-production.up.railway.app/todo
> + **`PATCH`** : https://todoapi-production.up.railway.app/todo
> + **`PUT`** : https://todoapi-production.up.railway.app/todo
> + **`DELETE`** : https://todoapi-production.up.railway.app/todo


1- **`GET`** - Get All item todo - HTTP Response Code: **200** <br>
`https://todoapi-production.up.railway.app/todo`
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

> [!TIP]
> Optional information to help a user be more successful.

> [!IMPORTANT]
> Crucial information necessary for users to succeed.

> [!WARNING]
> Critical content demanding immediate user attention due to potential risks.

> [!CAUTION]
> Negative potential consequences of an action.



