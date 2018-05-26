# How to use the JSON Database

For each http requests you'll have to add a `token` field inside your http request header (If you don't choose a token by using `-t yourtoken` while launching flume,, the default is 1234)

## Create a new db
```
POST /db/create
```
Parameters to give :

* `name` : The name of the db
* `content` : The content of the db


## Get an entire db
```
GET /db/read
```
* `name` : The name of the db


## Update the db
```
PUT /db/write
```
* `name` : The name of the db
* `content` : The content of the db