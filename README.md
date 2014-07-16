# Simple JSON API server in Golang

This is a small sample project I used to help me learn go. It will expand and change as I try to add more functionality

- First iteration uses a text file to store json
- Next steps:
    - Implement Tests
    - Refactor to multiple files
    - implement Postgresql
    
Working in conjunction with 
- Scott Skender
- Mike Gehard

## Get and Post tasks

our little server handles POST commands to `/task/new` with in the following format:

```
{"Name":"Wash Dishes","Description":"Is the name not clear enough?"}

```

you can also get all tasks with `/tasks`

## Take a look online!

With the help of Mike Gehard, we were able to deploy our server below!
http://107.170.31.184:8080/tasks

## Get it running!

I would write a little intro here to get things up and running on your machine... but the Golang docs doa  good job at that:

http://golang.org/doc/