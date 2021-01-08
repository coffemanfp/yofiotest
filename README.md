# YoFioTest
This is a test project for the YoFio Dev Golang vacancy.

The cloud version is in [Heroku](http://yofiotest.herokuapp.com/)

## Install

```sh
git clone https://github.com/coffemanfp/yofiotest.git
```


## Environment

### Server
**PORT**: Port server.                   (default: **8080**).


### PostgreSQL 
**DB_NAME**: Database name.         (default: **yofiotest**).
**DB_USER**: Database user.        (default: **yofiotest**).
**DB_PASSWORD**: Database password. (default: **yofiotest**).
**DB_HOST**: Database host.           (default: **localhost**).
**DB_PORT**: Database port.         (default: **5432**).
**DB_SSLMODE**: Database SSL mode.     (default: **disable**).
**DB_URL**: Database URL. When used, it overrides all previous values.


## Migrations
To run migrations you need to have an accessible database and the appropriate settings.

By default the following script executes the migrations and data examples on them.

```sh
    ./migrations.sh
```

To manually build the migrations, use:
(from work dir)

```sh
    cd migrations/postgresql
    go build -o ../../bin/postgresql
    cd -
```

And to run them::
(from work dir)

```sh
    cd migrations/postgresql
    ../../bin/postgresql
    cd -
```


## Run

Once the settings have been provided (or used by default) and the migrations have been run, the following script will build and run the server:

(from work dir)

```sh
    ./run.sh
```


## Test

To run the unit tests, along with coverage in the default browser, run:

(from work dir)

```sh
    ./test_cover.sh
```

To run the tests manually, run:

(from work dir)

```sh
    go test ./...
```


## End Points

### POST /create-assignment

Creates an assignment based on the investment provided. The assignment is saved even if it fails.

### Request

**Headers**

| Name         | Value            |
| ------------ | ---------------- |
| Content-Type | application/json |

**Body**

```json
{
    "investment": 3000
}
```


### Response

**Headers**

| Name           | Value            |
| -------------- | ---------------- |
| Content-Type   | application/json |
| Content-Length | (Body length)    |

**Body**

```json
// For a successed assignment
{
    "id": 1,
    "investment": 3000,
    "success": true,
    "credit_type_300": 2,
    "credit_type_500": 2,
    "credit_type_700": 2
}

// For a failed assigment
{
    "id": 1,
    "investment": 3000,
    "success": false,
    "credit_type_300": 0,
    "credit_type_500": 0,
    "credit_type_700": 0
}

// For error a internal server error.
{
    "error": "Oops!"
}
```


### POST /statistics

Gets the statistics of the assignments.

### Response

**Headers**

| Name           | Value            |
| -------------- | ---------------- |
| Content-Type   | application/json |
| Content-Length | (Body length)    |

**Body**

```json
// For a great response
{
    "total_asignments_done": 1,
    "total_asignments_success": 1,
    "total_asignments_fail": 0,
    "average_investment_successful": 3000,
    "average_investment_fail": 0
}

// For internal server error
{
    "error": "Oops!"
}
```

