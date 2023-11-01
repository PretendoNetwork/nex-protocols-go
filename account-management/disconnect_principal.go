// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisconnectPrincipal sets the DisconnectPrincipal handler function
func (protocol *Protocol) DisconnectPrincipal(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal uint32) uint32) {
	protocol.disconnectPrincipalHandler = handler
}

func (protocol *Protocol) handleDisconnectPrincipal(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.disconnectPrincipalHandler == nil {
		globals.Logger.Warning("AccountManagement::DisconnectPrincipal not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.disconnectPrincipalHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.disconnectPrincipalHandler(nil, packet, callID, idPrincipal)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
