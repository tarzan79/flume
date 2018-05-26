<p align="center">
  <!--<img src="https://raw.githubusercontent.com/dimensi0n/flume/master/flume.svg.png"/>-->
  <!-- <img src="https://raw.githubusercontent.com/dimensi0n/flume/master/flume-logo.png"/>-->
  <img src="https://cdn.discordapp.com/attachments/368671564608045058/449851948473253898/flume.png" width="400"/>
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
* Clone the repository (in your GOPATH) `git clone https://github.com/dimensi0n/flume.git`
* Go inside the flume folder
* Build it `go build`
* Launch it

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
