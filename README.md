# Gothpool
 
Gothpool is a Golang library for running fire-and-forget jobs using a fixed size goroutine pool and processing queue. 

## Install

``` sh
go get github.com/radutanasa/gothpool
```

Or, using dep:

``` sh
dep ensure -add github.com/radutanasa/gothpool
```

## Use

Simply instantiate the gothpool (executor pool) with your desired level of parallelism and queue size 
and start sending it functions.

``` go
package main

import (
    "github.com/radutanasa/gothpool"
)

func main() {	
    exec := gothpool.New(4, 1e5)
    exec.Start()
    defer exec.Stop()

    exec.Run(func() {
        // perform operations
    })	
}
```

It's important to note that the executor pool will finish its current job queue, but won't receive any new jobs. 

A `Run()` call on a stopped pool will return an `ExecPoolStoppedErr` error.

The pool can be restarted by calling the `Start()` function.

**IMPORTANT:** Variables sent to the executor should not change!

This code results in an unexpected behavior, as `i` references different values:
``` go
for i:=0; i<10; i++ {
    exec.Run(func() {
        println(i)
    })
} 
```
This results in the correct behavior:
``` go
for i:=0; i<10; i++ {
    var j = i
    exec.Run(func() {
        println(j)
    })
} 
```