// Package authentication implements the Authentication NEX protocol
package authentication

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Login sets the Login handler function
func (protocol *AuthenticationProtocol) Login(handler func(err error, client *nex.Client, callID uint32, strUserName string)) {
	protocol.LoginHandler = handler
}

func (protocol *AuthenticationProtocol) handleLogin(packet nex.PacketInterface) {
	if protocol.LoginHandler == nil {
		globals.Logger.Warning("Authentication::Login not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strUserName, err := parametersStream.ReadString()
	if err != nil {
		go protocol.LoginHandler(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.LoginHandler(nil, client, callID, strUserName)
}
