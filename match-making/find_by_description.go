// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescription sets the FindByDescription handler function
func (protocol *Protocol) FindByDescription(handler func(err error, packet nex.PacketInterface, callID uint32, strDescription string, resultRange *nex.ResultRange) uint32) {
	protocol.findByDescriptionHandler = handler
}

func (protocol *Protocol) handleFindByDescription(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByDescriptionHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescription not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescription, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findByDescriptionHandler(fmt.Errorf("Failed to read strDescription from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByDescriptionHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByDescriptionHandler(nil, packet, callID, strDescription, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
