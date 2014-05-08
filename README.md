## Go Set

This is a simple data structure implementation which uses `map` internally to store the set.

The value of each item is an empty structure since empty structure uses 0 bytes.

### Usage

The Go Set is Non Thread safe library. So make sure that it is protected with mutex if program involves multiple insertion or deletion.

Installing the go into your Go

```
go get github.com/alinz/go-set
```

One that is done, import the go-set and use the `NewNonThreadSafeSet` method to create a new `Set`.

```
package main

import (
  "fmt"
  "github.com/alinz/go-set"
)

func main() {
  mySet := set.NewNonThreadSafeSet()

  mySet.Add(1)
  mySet.Add(2)

  mySet.Iterate(func (item interface{}){
    number := item.(int)
    fmt.Println(number)
  })
}
```
