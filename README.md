# microservice-test

This project consists of 1 grpc service (in `movieservice` directory) and the API gateway (in `gateway` directory)

## Run Application

### *Prerequisite*

```
go version >= 1.10
```

### *Run movieservice*

```
cd movieservice
go run cmd/main.go
```

### *Run gateway*

```
cd gateway
go run cmd/main.go
```

See the .env file in both directories if you want to adjust their envs

## Endpoints

### *GET /search/{searchkeyword}/{page}*

Request Example

```
curl -X GET \
  http://localhost:8080/search/nice/5 \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache'
```

Response Example

```
{
    "Search": [
        {
            "Title": "Batman Begins",
            "Year": "2005",
            "imdbID": "tt0372784",
            "Type": "movie",
            "Poster": "https://m.media-amazon.com/images/M/MV5BZmUwNGU2ZmItMmRiNC00MjhlLTg5YWUtODMyNzkxODYzMmZlXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_SX300.jpg"
        },
        {
            "Title": "Batman v Superman: Dawn of Justice",
            "Year": "2016",
            "imdbID": "tt2975590",
            "Type": "movie",
            "Poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
        },
        {
            "Title": "Batman",
            "Year": "1989",
            "imdbID": "tt0096895",
            "Type": "movie",
            "Poster": "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"
        },
        ...
    ],
    "totalResults": "379",
    "Response": "True"
}
```
