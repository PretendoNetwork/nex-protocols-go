// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRelationships sets the GetRelationships handler function
func (protocol *Protocol) GetRelationships(handler func(err error, client *nex.Client, callID uint32, resultRange *nex.ResultRange) uint32) {
	protocol.getRelationshipsHandler = handler
}

func (protocol *Protocol) handleGetRelationships(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRelationshipsHandler == nil {
		globals.Logger.Warning("Friends::GetRelationships not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.getRelationshipsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRelationshipsHandler(nil, client, callID, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
