// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCommonData sets the DeleteCommonData handler function
func (protocol *Protocol) DeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID uint64) uint32) {
	protocol.deleteCommonDataHandler = handler
}

func (protocol *Protocol) handleDeleteCommonData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteCommonDataHandler == nil {
		globals.Logger.Warning("Ranking2::DeleteCommonData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.deleteCommonDataHandler(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteCommonDataHandler(nil, packet, callID, nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
