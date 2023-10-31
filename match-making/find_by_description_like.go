// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByDescriptionLike sets the FindByDescriptionLike handler function
func (protocol *Protocol) FindByDescriptionLike(handler func(err error, packet nex.PacketInterface, callID uint32, strDescriptionLike string, resultRange *nex.ResultRange) uint32) {
	protocol.findByDescriptionLikeHandler = handler
}

func (protocol *Protocol) handleFindByDescriptionLike(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByDescriptionLikeHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByDescriptionLike not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strDescriptionLike, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findByDescriptionLikeHandler(fmt.Errorf("Failed to read strDescriptionLike from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByDescriptionLikeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByDescriptionLikeHandler(nil, packet, callID, strDescriptionLike, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
