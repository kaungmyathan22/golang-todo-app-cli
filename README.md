# Todo CLI App

A simple command-line application written in Golang to manage your todo list. You can save tasks to a JSON file, mark them as done, and display your todos in a pretty table.

## Features

- Add new todos to your list.
- Mark todos as done.
- View your todos in a formatted table.

## Installation

```bash
$ go get
```
```bash
$ go build ./cmd/todo
```

## Usage

#### Add a todo
```bash
$ ./todo -add "do laundary"
```
#### List all todo
```bash
$ ./todo -list
```
#### List a todo as done
```bash
$ ./todo -complete 1 # where 1 is id of the todo
```
#### Delete a todo
```bash
$ ./todo -del 1 # where 1 is id of the todo
```
#### View available options
```bash
$ ./todo --help
```