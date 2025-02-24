# Rest API in Go

This is an API that is capable of managing events, thus being able to create, remove and update events. It is also possible to register users for events and remove a user's registration from an event.

## Routes

```http
POST {host}/events      # creating a new event
```
```http
GET {host}/events       # searching for all events
```
```http
GET {host}/events/id    # searching for an event by id
```
```http
PUT {host}/events/id    # updating a specific event
```
```http
DELETE {host}/events/id   # deleting a specific event
```
```http
POST {host}/signup      # registering a new user
```
```http
POST {host}/login       # user login
```
```http
POST {host}/events/id/register      # registering a user in a specific event
```
```http
DELETE {host}/events/id/register    # deleting a user's record of a given event
```

#### Examples of how routes are used are in the folder `/api-test`
