// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByNameLike sets the FindByNameLike handler function
func (protocol *Protocol) FindByNameLike(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange) uint32) {
	protocol.findByNameLikeHandler = handler
}

func (protocol *Protocol) handleFindByNameLike(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByNameLikeHandler == nil {
		globals.Logger.Warning("AccountManagement::FindByNameLike not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findByNameLikeHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strLike, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findByNameLikeHandler(fmt.Errorf("Failed to read strLike from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByNameLikeHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByNameLikeHandler(nil, packet, callID, uiGroups, strLike, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
