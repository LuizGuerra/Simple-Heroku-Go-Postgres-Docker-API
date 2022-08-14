# GoLang API with Postgres database using Docker

This is a really simple project focused into creating a repo that can be runned in Heroku servers.

The database also have requests using the relation method ('home' table contains foreigner key of the 'agent' table)

### Running

docker compose up --build

### RESTful commands

These are simple ```curl``` commands with the intent of making it easier for you to test and use this simple API program in the terminal.

#### Adding new data

``` curl -i -X POST localhost:8080/homes -d @post_req.json ```

#### Getting all data

``` curl -i -X GET localhost:8080/homes ```

#### Updating data

``` curl -i -X POST localhost:8080/homes -d @put_req.json ```

#### Getting data by ID

``` curl -i -X GET localhost:8080/homes/<ID HERE> ```

#### Deleting data by ID

``` curl -i -X DELETE localhost:8080/homes/<ID HERE> ```
