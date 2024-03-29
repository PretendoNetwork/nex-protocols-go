// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescriptionRegex sets the FindByDescriptionRegex handler function
func (protocol *Protocol) FindByDescriptionRegex(handler func(err error, packet nex.PacketInterface, callID uint32, strDescriptionRegex string, resultRange *nex.ResultRange) uint32) {
	protocol.findByDescriptionRegexHandler = handler
}

func (protocol *Protocol) handleFindByDescriptionRegex(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByDescriptionRegexHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescriptionRegex not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescriptionRegex, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findByDescriptionRegexHandler(fmt.Errorf("Failed to read strDescriptionRegex from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByDescriptionRegexHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByDescriptionRegexHandler(nil, packet, callID, strDescriptionRegex, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
