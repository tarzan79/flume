# How to contribute

## 1. Folder Architecture:

```
.github // The github files like CONTRIBUTIONS.md
    _ CONTRIBUTION.md

auth // The auth package folder
    _ auth.go

db // The JSON database package folder
    _ db.go

fs // The file storage package folder
    _ fs.go

main.go // The main file of flume
```

## 2. The Developpement Environment

* You have to install [golang](https://golang.org)
* Fork the repo and clone it into `$GOPATH/src/github.com/dimensi0n/flume` (Replace GOPATH with your GOPATH)
* Go inside `go/src/dimensi0n/flume`
* Make your modifications
* Run `go run main.go` to verify if it works
* Then Commit and make a Pull Request

## 3. Pull Request

If you wan't to add a new feature, you have to open an issue first to ask the community if it's good

The pull request should look like this:

```
title: [new feature|bug fix] The name of your feature or the name of the bug.

message: Why do you add this Feature and what is this new feature / What was the bug and how you fixed it
```

**Thanks to all developpers ;)**

**You are just awesome**
