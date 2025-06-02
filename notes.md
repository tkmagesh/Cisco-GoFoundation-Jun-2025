# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
|----|----|
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins)
| Wind Up | 5:00 PM |


## Repository
- https://github.com/tkmagesh/cisco-gofoundation-jun-2025
## Methodology
- No powerpoint
- Discuss & Code
- Anytime is good for Q&A

## Software Requirements
- Go Tools (https://go.dev/dl)
    - Verification
    ```shell
    go version
    ```
- Visual Studio Code (Or any other editor)
    - Go Extension for VSCode (https://marketplace.visualstudio.com/items?itemName=golang.Go)


## Why Go?
1. Simplicity
    - ONLY 25 keywords
    - No access modifiers (public, private, protected etc)
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No classes (only structs)
    - No inheritance (only composition)
    - No exceptions (only errors)
    - No "try..catch..finally"
    - No implicit type conversions
2. Performance
    - Close to hardware
        - Compiled to machine code (like C & C++)
        - `Cross compilation` support is enabled in the compiler
    - No need for a VM
        - How is memory managed (where is the GC)?
    - Performance comparable to C++
3. Concurrency Support
    - Concurrent operations are represented as `Goroutines`
    - `Goroutines` are very cheap (2 KB) when compared to OS Threads (~2MB)
    - Support for concurrency is offered `through the language`
        - Language features for concurrency
            - `go` keyword
            - `chan` data type
            - `<-` operator 
            - `range` construct
            - `select-case` construct
    - Standard Library API support
        - `sync` package
        - `sync/atomic` package

## Compilation
```shell
go build [filename.go]
```

## Compile & Execute
```shell
go run [filename.go]
```
