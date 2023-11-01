// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFromBlockList sets the RemoveFromBlockList handler function
func (protocol *Protocol) RemoveFromBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID []uint32) uint32) {
	protocol.removeFromBlockListHandler = handler
}

func (protocol *Protocol) handleRemoveFromBlockList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.removeFromBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::RemoveFromBlockList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipalID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.removeFromBlockListHandler(fmt.Errorf("Failed to read lstPrincipalID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.removeFromBlockListHandler(nil, packet, callID, lstPrincipalID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
