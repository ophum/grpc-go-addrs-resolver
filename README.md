# grpc-go-addrs-resolver

grpc-goで複数アドレスをカンマ区切りで指定できるresolver

```
package main 

import (
    _ "github.com/ophum/grpc-go-addrs-resolver
)

func main() {
    conn, err := grpc.NewClient("addrs:///localhost:50051,localhost:50052")
}
```