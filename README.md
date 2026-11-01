# Exercise `go001`

## Purpose

To implement the simplest possible REST endpoint in Golang.

The goal of the exercise is attained, if the following test succeeds.

### Step 1

Run the tests using `go test`. Expected result :

```
PASS
ok  	go001	...	
```

### Step 2

Build the application using `build.sh` or

```
go build -o bin/barserver .
```

### Step 3

Start the application using `bin/barserver`.

### Step 4

Run

```
curl http://localhost:9000/foo
```

Expected Result: The endpoint returns `"bar"`.
