// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDetailedList sets the GetDetailedList handler function
func (protocol *Protocol) GetDetailedList(handler func(err error, client *nex.Client, callID uint32, byRelationship uint8, bReversed bool) uint32) {
	protocol.getDetailedListHandler = handler
}

func (protocol *Protocol) handleGetDetailedList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getDetailedListHandler == nil {
		globals.Logger.Warning("Friends::GetDetailedList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	byRelationship, err := parametersStream.ReadUInt8()
	if err != nil {
		errorCode = protocol.getDetailedListHandler(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), client, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReversed, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.getDetailedListHandler(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), client, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getDetailedListHandler(nil, client, callID, byRelationship, bReversed)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
