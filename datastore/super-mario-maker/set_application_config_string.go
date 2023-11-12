// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetApplicationConfigString sets the SetApplicationConfigString handler function
func (protocol *Protocol) SetApplicationConfigString(handler func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32, key uint32, value string) uint32) {
	protocol.setApplicationConfigStringHandler = handler
}

func (protocol *Protocol) handleSetApplicationConfigString(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.setApplicationConfigStringHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetApplicationConfigString not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setApplicationConfigStringHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	key, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setApplicationConfigStringHandler(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	value, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.setApplicationConfigStringHandler(fmt.Errorf("Failed to read value from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.setApplicationConfigStringHandler(nil, packet, callID, applicationID, key, value)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
