// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByOwner sets the FindByOwner handler function
func (protocol *Protocol) FindByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id uint32, resultRange *nex.ResultRange) uint32) {
	protocol.findByOwnerHandler = handler
}

func (protocol *Protocol) handleFindByOwner(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByOwnerHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByOwner not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findByOwnerHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByOwnerHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByOwnerHandler(nil, packet, callID, id, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
