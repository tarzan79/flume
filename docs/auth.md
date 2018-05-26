# How to use the Auth system

For each http requests you'll have to add a `token` field inside your http request header (If you don't choose a token by using `-t yourtoken` while launching flume,, the default is 1234)

## Create a user
```
POST /auth/signup
```
Parameters to give :

* `username` : The user name
* `password` : The user password
* `email` : The user email

## Login an user
```
GET /auth/login
```
Parameters to give :

* `username` : The user name
* `password` : The user password

## Disconnect an user
```
GET /auth/disconnect
```
Parameters to give :

* `username` : The user name
* `password` : The user password

## Delete an user
```
DELETE /auth/delete
```
Parameters to give :

* `username` : The user name
* `password` : The user password

## Get the list of users
```
GET /auth/all
```
No parameters to give