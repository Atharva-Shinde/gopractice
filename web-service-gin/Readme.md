## A Go RESTful API using Gin Framework
 All the data is stored in the memory (rather than using a database)

## To try this API, execute these command in the root directory of the project
```
go get .
go run .
```
----

 This API has two endpoints POST and GET

- POST:
    Use the following command to add new cover to the existing list
    ``` 
    curl http://localhost:8080/covers \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "5","title": "For the ones to whom neither the past nor the future belongs","artist": "All shall be well","premier": 11.2014}' 
    ```

- GET: 
    - The `curl localhost:8080/covers` fetches data of the covers in JSON format,or you can view the JSON output on your browser through: `localhost:8080/covers`
    - `curl http:localhost:8080/covers/id` fetches the cover with the assigned id.
### Output
```
[
    {
        "id": "1",
        "title": "Interstellar",
        "artist": "Hans Zimmer",
        "premier": 10.2014,
        "time_added": "0001-01-01T00:00:00Z"
    },
    {
        "id": "2",
        "title": "Building Light",
        "artist": "Anne Sophie Versnaeyen",
        "premier": 2018,
        "time_added": "0001-01-01T00:00:00Z"
    },
    {
        "id": "3",
        "title": "Zen",
        "artist": "Torin Borrowdale",
        "premier": 3.2018,
        "time_added": "0001-01-01T00:00:00Z"
    },
    {
        "id": "4",
        "title": "Married Life",
        "artist": "Jason Lyle Black",
        "premier": 10.2015,
        "time_added": "0001-01-01T00:00:00Z"
    },
    {
        "id": "5",
        "title": "For the ones to whom neither the past nor the future belongs",
        "artist": "All shall be well",
        "premier": 11.2014,
        "time_added": "0001-01-01T00:00:00Z"
    }
]
```

#### curl http://localhost:8080/covers/2
```
{
    "id": "2",
    "title": "Building Light",
    "artist": "Anne Sophie Versnaeyen",
    "premier": 2018
}
```