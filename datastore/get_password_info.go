// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPasswordInfo sets the GetPasswordInfo handler function
func (protocol *Protocol) GetPasswordInfo(handler func(err error, packet nex.PacketInterface, callID uint32, dataID uint64) uint32) {
	protocol.getPasswordInfoHandler = handler
}

func (protocol *Protocol) handleGetPasswordInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPasswordInfoHandler == nil {
		globals.Logger.Warning("DataStore::GetPasswordInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getPasswordInfoHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPasswordInfoHandler(nil, packet, callID, dataID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
