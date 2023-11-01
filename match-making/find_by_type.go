// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByType sets the FindByType handler function
func (protocol *Protocol) FindByType(handler func(err error, packet nex.PacketInterface, callID uint32, strType string, resultRange *nex.ResultRange) uint32) {
	protocol.findByTypeHandler = handler
}

func (protocol *Protocol) handleFindByType(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByTypeHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByType not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strType, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findByTypeHandler(fmt.Errorf("Failed to read strType from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByTypeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByTypeHandler(nil, packet, callID, strType, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
