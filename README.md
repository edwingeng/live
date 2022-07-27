# Overview
`live.Data` is a handy general-purpose data wrapper, which solves the problem that the same data type in different versions of a plugin are actually not the same at runtime. It is designed to work with [hotswap](https://github.com/edwingeng/hotswap), which provides a solution for reloading your go code without restarting your server, interrupting or blocking any ongoing procedure.

# Getting Started
```
go get -u github.com/edwingeng/live
```

# Usage
```go
type questInfo1 struct {
    ID   int64
    Name string
    Done bool
}

q1 := questInfo1{
    ID:   5848,
    Name: "Of Love and Family",
}

ld := live.WrapObject(&q1)

// 2000 years later...

type questInfo2 struct {
    ID   int64
    Name string
    Desc string
    Done bool
}

var q2 questInfo2
q2.Desc = "<>"
ld.UnwrapObject(&q2)
fmt.Printf("ID: %v\n", q2.ID)
fmt.Printf("Name: %v\n", q2.Name)
fmt.Printf("Desc: %v\n", q2.Desc)
fmt.Printf("Done: %v\n", q2.Done)

// Output:
// ID: 5848
// Name: Of Love and Family
// Desc: <>
// Done: false
```