# How to install (mac)
$ brew install go

# How to install (linux)
$ sudo apt install golang-go


# Check version / Installation
$ go version 

# Creating the program
$ mkdir program

$ cd program

$ mkdir hello greetings

$ cd hello

$ go mod init test.com/hello

$ touch hello.go

$ cd ..

$ cd greetings

$ go  mod init test.com/greetings

$ touch greetings.go

$ cd ..

$ cd hello

$ go mod edit --replace test.com/greetings=../greetings

$ go mod tidy

$ go run .


# Making build
$ go list -f '{{.Target}}' ## (in your hello directory) ## Suppose -- /xyz/abc/go/bin/hello

$ export PATH=$PATH:/xyz/abc/go/bin/hello

$ go env -w GOBIN=/xyz/abc/go/bin

$ go install

$ ./hello

$ go build 

# Reopen terminal
$ export PATH=$PATH:/xyz/abc/go/bin/hello

$ hello

