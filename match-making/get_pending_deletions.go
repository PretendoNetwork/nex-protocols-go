// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPendingDeletions sets the GetPendingDeletions handler function
func (protocol *Protocol) GetPendingDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, uiReason uint32, resultRange *nex.ResultRange) uint32) {
	protocol.getPendingDeletionsHandler = handler
}

func (protocol *Protocol) handleGetPendingDeletions(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPendingDeletionsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetPendingDeletions not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiReason, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getPendingDeletionsHandler(fmt.Errorf("Failed to read uiReason from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.getPendingDeletionsHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPendingDeletionsHandler(nil, packet, callID, uiReason, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
