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

* Install [docker](https://docs.docker.com/install/)
* Run a container with the flume image `docker run -d -e TOKEN=yourtoken -p 8080:8080 --name flume dimensi0n/flume` You have to replace `yourtoken` with a JSON Authentication Token of your choice (the default one is 1234)

### Use flume

You can now access to flume on the port 8080

You can know take a look at the [wiki](https://github.com/dimensi0n/flume/wiki)

Enjoy :)
