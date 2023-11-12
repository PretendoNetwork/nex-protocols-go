// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetList sets the GetList handler function
func (protocol *Protocol) GetList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship uint8, bReversed bool) uint32) {
	protocol.getListHandler = handler
}

func (protocol *Protocol) handleGetList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getListHandler == nil {
		globals.Logger.Warning("Friends::GetList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	byRelationship, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.getListHandler(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReversed, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.getListHandler(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getListHandler(nil, packet, callID, byRelationship, bReversed)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
