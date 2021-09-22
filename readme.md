# Vip-Management-Api

This is my first golang project. This application was created to solve a fictitious case of VIP arrival management.

## How to run ?
<!-- markdownlint-disable MD029 -->

I don't actually know what the best practice for this, but for me

1. Clone the repo. I'm not sure if it has to be inside `$GOPATH`, but just to make sure it works clone there.
2. Create Postgres database

```sql
CREATE DATABASE your_db_name
```

3. Create database table

```sql
CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    name TEXT,
    country_of_origin TEXT,
    photo TEXT,
    arrived BOOL,
    attributes [] TEXT,
);
```

4. CD into project directory.

5. Download all dependencies

```bash
go mod download
```

6. Copy `.env.example` to `.env`
7. Change the contents according to your need.
8. Start the app using `go run` command

```bash
go run main.go
```

9. Check the app using rest client like postman

## APIs

### Create VIP

`POST` : `localhost:4000/api/vips`

Sample request

```json
{
    "name": "Go Gopher G",
    "country_of_origin": "Indonesia",
    "eta": "2021-09-20T11:25:27.194Z",
    "photo": "https://golang.org/lib/godoc/images/home-gopher.png",
    "arrived": false,
    "attributes":  [
        "black flag",
        "blue hat",
        "fire in the heart"
    ]
}
```

Sample response

```json
{
    "message": "Successfully create data",
    "ok": true
}
```

### Get VIPs

Get all VIPs list, returning array

`GET` : `localhost:4000/api/vips`

Sample response

```json
{
    "data": [
        {
            "id": 17,
            "name": "Go Gopher G",
            "country_of_origin": "Indonesia",
            "eta": "2021-09-20T00:00:00Z",
            "photo": "https://golang.org/lib/godoc/images/home-gopher.png",
            "arrived": false,
            "attributes": [
                "black flag",
                "blue hat",
                "fire in the heart"
            ]
        }
    ],
    "message": "Successfully retrieve data",
    "ok": true
}
```

### Get VIP

Get one VIP by id

`GET` : `localhost:4000/api/vips/:id`

Sample response

```json
{
    "data": {
        "id": 17,
        "name": "Go Gopher G",
        "country_of_origin": "Indonesia",
        "eta": "2021-09-20T00:00:00Z",
        "photo": "https://golang.org/lib/godoc/images/home-gopher.png",
        "arrived": false,
        "attributes": [
            "black flag",
            "blue hat",
            "fire in the heart"
        ]
    },
    "message": "Successfully retrieve data",
    "ok": true
}
```

### Update VIP

Update one vip by id

`PUT` : `localhost:4000/api/vips/:id`

Sample request

```json
{
    "name": "Go Go",
    "country_of_origin": "Russia",
    "eta": "2021-09-20T11:25:27.194Z",
    "photo": "https://golang.org/lib/godoc/images/home-gopher.png",
    "arrived": false,
    "attributes":  [
        "blue hat",
        "fire in the heart"
    ]
}
```

Sample response

```json
{
    "message": "Successfully update data",
    "ok": true
}
```

### Delete VIP

Delete one vip by id

`DELETE` : `localhost:4000/api/vips/:id`

Sample response

```json
{
    "message": "Successfully delete data",
    "ok": true
}
```

### Arrived VIP

Set oen VIP arrived status by id

`PATCH` : `localhost:4000/api/vips/:id`

Sample response

```json
{
    "message": "Successfully update data",
    "ok": true
}
```

## Whats Next ?

This is only the start, hopefully this repo could be my playground on golang. I got some plan to do ahead, some may include but not limited to:

- Authentication
- Implement clean code architecture
