// Package authentication implements the Authentication NEX protocol
package authentication

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// LoginWithParam sets the LoginWithParam handler function
func (protocol *AuthenticationProtocol) LoginWithParam(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.LoginWithParamHandler = handler
}

func (protocol *AuthenticationProtocol) handleLoginWithParam(packet nex.PacketInterface) {
	if protocol.LoginWithParamHandler == nil {
		globals.Logger.Warning("Authentication::LoginWithParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	// Unsure what data is sent here, or how to trigger the console to send it
}
