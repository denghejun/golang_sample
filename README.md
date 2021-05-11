## A multiple modules based on Golang (IDE GoLand)

## Create module based project in GoLand
```
File->New->Project...->Select "Go modules"
```

## Download all dependencies
Go mod vendor will collect all the external packages which are used in this main module, even the remote package/dependency was broken, we still can build the project successfully.
```
go mod vendor
```


## Run

```
go run ./main
```

## Build an executable file
```
cd main
go build
```
# golang_sample
