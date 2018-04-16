# book_api

[![Build Status](https://travis-ci.org/phanyzewski/book_api.svg?branch=master)](https://travis-ci.org/phanyzewski/book_api)

Book api provides an HTTP API for storing books.

## Usage

```shell
go get -u github.com/phanyzewski/book_api
cd $GOPATH/src/github.com/phanyzewski/book_api
go build
book_api
```

Env variables can also be set with `.env` file in the same
directory as the server. Variables in the `.env` file will override
any variables set in the session. For more information on how to use
an env file see [godotenv](https://github.com/joho/godotenv).

```shell
BOOK_API_URL=localhost:8080
DATABASE_URL=postgres://dev@localhost/book_development?sslmode=disable
PGSSLMODE=disable
```

Database migrations are managed with [pgmgr](https://github.com/rnubel/pgmgr).
```shell
cd $GOPATH/src/github.com/phanyzewski/book_api
pgmgr db create
pgmgr db migrate
```

Dependency management is handled with [dep](https://github.com/golang/dep).
```shell
cd $GOPATH/src/github.com/phanyzewski/book_api
dep ensure
```

## API

* Books

  ```GET /books ```

  ```POST /books ```

  ```GET /book/:book_id ```

  ```PUT /book/:book_id ```

  ```DELETE /book/:book_id ```

* Authors

  ```GET /authors ```

  ```POST /authors ```

  ```GET /author/:author_id ```

  ```PUT /author/:author_id ```

  ```DELETE /author/:author_id ```

* Publishers

  ```GET /publishers ```

  ```POST /publishers ```

  ```GET /publisher/:publisher_id ```

  ```PUT /publisher/:publisher_id ```

  ```DELETE /publisher/:publisher_id ```
