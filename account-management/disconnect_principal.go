// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DisconnectPrincipal sets the DisconnectPrincipal handler function
func (protocol *Protocol) DisconnectPrincipal(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32)) {
	protocol.disconnectPrincipalHandler = handler
}

func (protocol *Protocol) handleDisconnectPrincipal(packet nex.PacketInterface) {
	if protocol.disconnectPrincipalHandler == nil {
		globals.Logger.Warning("AccountManagement::DisconnectPrincipal not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.disconnectPrincipalHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.disconnectPrincipalHandler(nil, client, callID, idPrincipal)
}
