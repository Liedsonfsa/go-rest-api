# Rest API in Go <img src="images/go-gopher-svgrepo-com.svg" width="30px">

This is an API that is capable of managing events, thus being able to create, remove and update events. It is also possible to register users for events and remove a user's registration from an event.

## Cloning and accessing the repository
```bash
git clone https://github.com/Liedsonfsa/go-rest-api.git
cd go-rest-api
```

## Downloading dependencies
```go
go mod tidy
```

## Environment variables
Create a .env file with the following variables

```bash
SECRET_KEY= # secrete key for signing jwt tokens
DB_NAME= # database name
DB_USER= # database user
DB_PASSWORD= # database password
```

## Running the API
```go
go run main.go
```

## Routes that do not require authentication

```bash
GET {host}/events       # searching for all events
```
```bash
GET {host}/events/id    # searching for an event by id
```

```bash
POST {host}/signup      # registering a new user
```
```bash
POST {host}/login       # user login
```


## Routes that require authentication

```bash
POST {host}/events      # creating a new event
```
```bash
PUT {host}/events/id    # updating a specific event
```
```bash
DELETE {host}/events/id   # deleting a specific event
```
```bash
POST {host}/events/id/register      # registering a user in a specific event
```
```bash
DELETE {host}/events/id/register    # deleting a user's record of a given event
```

#### Examples of how routes are used are in the folder `/api-test`

## Technologies

- Go
- Gin
- MySQL
- JWT