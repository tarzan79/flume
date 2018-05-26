<p align="center">
  <img src="http://90.109.222.6:5333/flume_300ppp.png" width="400"/>
</p>

<br>

```markdown
Flume is a cloud that allows you to :

* Interact with a JSON database
* Create an user, log in the user, disconnect the user
* Access a file storage

With the power of HTTP Requests.

Flume is based on an SQL database and it's written in Golang.
```


## Getting started

### Install flume

* Install [golang](https://golang.org/)
* Build flume `go build github.com/dimensi0n/flume`
* Launch it `./flume -t yourtoken` you have to choose a JSON Authentication Token (the default one is 1234)

### Use flume

You can now access to flume on the port 8080
The details of the following endpoints are in the doc (link of the doc)

```
/db/create
# Create a new db
```
```
/db/read
# Get an entire db
```
```
/db/write
# Update the db
```
```
/auth/signup
# Create a user
```
```
/auth/login
# Login an user
```
```
/auth/disconnect
# Disconnect an user
```
```
/auth/delete
# Delete an user
```
```
/auth/all
# Get the list of users
```
