// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStatus sets the GetStatus handler function
func (protocol *Protocol) GetStatus(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32) uint32) {
	protocol.getStatusHandler = handler
}

func (protocol *Protocol) handleGetStatus(packet nex.PacketInterface) {
	if protocol.getStatusHandler == nil {
		globals.Logger.Warning("AccountManagement::GetStatus not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getStatusHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getStatusHandler(nil, client, callID, idPrincipal)
}
