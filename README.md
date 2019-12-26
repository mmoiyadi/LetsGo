# LetsGo

## Test two sided primeness of a number
```
<baseurl>/isNumberTwoSidedPrime/{num}
```
  If num is valid integer, will return true if the integer is two sided prime, false otherwise.
  If num is invalid, will return "invalid integer"

```
<baseurl>/isNumberPrime/{num}
```
  If num is valid integer, will return true if the integer is prime, false otherwise.
  If num is invalid, will return "invalid integer"
  

## Build
The following command will retrieve the dependency package

```
go get
```

## Run
The following command will start API server at port 8080

```
go run prime_check.go
```

## Test
The following command will run the unit test for the application

```
go test
```

