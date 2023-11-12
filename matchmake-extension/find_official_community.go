// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindOfficialCommunity sets the FindOfficialCommunity handler function
func (protocol *Protocol) FindOfficialCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange) uint32) {
	protocol.findOfficialCommunityHandler = handler
}

func (protocol *Protocol) handleFindOfficialCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findOfficialCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::FindOfficialCommunity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	isAvailableOnly, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.findOfficialCommunityHandler(fmt.Errorf("Failed to read isAvailableOnly from parameters. %s", err.Error()), packet, callID, false, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findOfficialCommunityHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, false, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findOfficialCommunityHandler(nil, packet, callID, isAvailableOnly, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
