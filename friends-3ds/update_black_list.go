// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateBlackList sets the UpdateBlackList handler function
func (protocol *Protocol) UpdateBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, unknown []uint32) uint32) {
	protocol.updateBlackListHandler = handler
}

func (protocol *Protocol) handleUpdateBlackList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateBlackListHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateBlackList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.updateBlackListHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateBlackListHandler(nil, packet, callID, unknown)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
