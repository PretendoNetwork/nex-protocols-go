# NEX Protocols Go
## NEX servers with protocol support in Go

[![GoDoc](https://godoc.org/github.com/PretendoNetwork/nex-protocols-go?status.svg)](https://godoc.org/github.com/PretendoNetwork/nex-protocols-go)

### Other NEX libraries
[nex-go](https://github.com/PretendoNetwork/nex-go) - Barebones NEX/PRUDP server implementation

[nex-protocols-common-go](https://github.com/PretendoNetwork/nex-protocols-common-go) - NEX protocols used by many games with premade handlers and a high level API

### Install

`go get github.com/PretendoNetwork/nex-protocols-go`

### Usage

`nex-protocols-go` provides a higher level API than the [NEX Go module](https://github.com/PretendoNetwork/nex-go) to the underlying PRUDP server by providing a set of NEX protocols. This module only provides access to the lower level raw RMC method calls, however, and all method handlers must be defined in full manually. For a higher level API, see the [common NEX method handlers module](https://github.com/PretendoNetwork/nex-protocols-common-go)

### Example, friends (Wii U) authentication server
### For a complete example, see the complete [Friends Authentication Server](https://github.com/PretendoNetwork/friends-authentication), and other game servers

```go
package main

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	nexproto "github.com/PretendoNetwork/nex-protocols-go"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPrudpVersion(0)
	nexServer.SetSignatureVersion(1)
	nexServer.SetKerberosKeySize(16)
	nexServer.SetAccessKey("ridfebb9")

	authenticationServer := nexproto.NewAuthenticationProtocol(nexServer)

	// Handle Login RMC method
	authenticationServer.Login(login)

	// Handle RequestTicket RMC method
	authenticationServer.RequestTicket(requestTicket)

	nexServer.Listen(":60000")
}
```