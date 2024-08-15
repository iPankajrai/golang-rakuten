- create a directy 
- cd to the directory 
- go mod init demo

- to run

- compiles + links + creates binary + executes 

```
go run main.go
```

- to build 

```
go build main.go

```

- to get the list of go os and go arch

```
go tool dist list 
```

- cross compilation
```
GOARCH=amd64 GOOS=windows go build -o hello.exe main.go 
```
# Keywords

- package, func,for,import,if,else,switch, return, default, break, continue, var, const,case, map, range, type, struct, fallthrough, defer ,go,goto,interface,chan,select

- 

# builtin functions

- print , println, len,cap, append, copy, make, delete, new , recover ,panic.close

Tasks: Go to the tasks/ directory to find tasks
https://github.com/JitenPalaparthi/golang-training-questglobal

Create and deploy a package
https://github.com/JitenPalaparthi/rakutenshapes


# ===================== June 2024 onwards
- Go provides sort of OOPs (Not completely oops)
    - If we declare a variable or function name starting with a capital letter then it is available outside of package also 

- Receivers: structs with functions
    - we can tie a function with a struct by providing the struct pointer as a reciever to that function.
