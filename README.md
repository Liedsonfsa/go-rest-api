# Rest API in Go <img src="images/go-gopher-svgrepo-com.svg" width="20px">

This is an API that is capable of managing events, thus being able to create, remove and update events. It is also possible to register users for events and remove a user's registration from an event.

## Routes

```bash
POST {host}/events      # creating a new event
```
```bash
GET {host}/events       # searching for all events
```
```bash
GET {host}/events/id    # searching for an event by id
```
```bash
PUT {host}/events/id    # updating a specific event
```
```bash
DELETE {host}/events/id   # deleting a specific event
```
```bash
POST {host}/signup      # registering a new user
```
```bash
POST {host}/login       # user login
```
```bash
POST {host}/events/id/register      # registering a user in a specific event
```
```bash
DELETE {host}/events/id/register    # deleting a user's record of a given event
```

#### Examples of how routes are used are in the folder `/api-test`
