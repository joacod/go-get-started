# go-get-started
<p align="center">
  <img height="200" src="https://pilsniak.com/wp-content/uploads/2017/04/golang.jpg">
</p>
Go playground project, created to play a little bit with this awesome language and learn in the process.

## Prerequisites
- [Install Go](https://golang.org/)
- [VS Code Extension for Go](https://github.com/microsoft/vscode-go)

## How to

### Need Help?
From the terminal: `go run playground --help`

### Run it
From the terminal: `go run playground --artist Beatles --song "Come Together"`

### Build it
From the terminal: `go build playground`

Once it's finished, you will find an executable (specific to your OS) in the root folder.

## What you will find so far

### Current playground :heavy_check_mark:
- [x] Run a console application to get song Lyrics, from an Artist and a Song name

#### Notes
Behind the scenes the program is using the following public api https://lyricsovh.docs.apiary.io/#

### Deprecated
Despite they are deprecated, you can find the corresponding methods in the code.
- [x] Console print of message "Hello Go!"
- [x] Console print of a struct
- [x] Run a basic and simplified WebServer (only supporting GET)

Running the WebServer
- Modify the code to use method named "runUsersWebService"
- From the terminal: `go run playground`
- Once it's running, go to a browser
- Hit the following url (Port 3000 is harcoded) http://localhost:3000/users
