- to do escape analysis

```
go build --gcflags="-m" main.go
```

- There are there things to keep in mind when doing escape analysis

1- does not escape   -> stored in stack
2- escapes to heap   -> stored in heap
3- moved to heap     -> initially stored in stack but later moved to heap;Ultimately stored in heap