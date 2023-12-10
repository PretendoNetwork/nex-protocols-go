// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestMigration(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.RequestMigration == nil {
		globals.Logger.Warning("DataStorePokemonBank::RequestMigration not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	oneTimePassword, err := parametersStream.ReadString()
	if err != nil {
		_, errorCode = protocol.RequestMigration(fmt.Errorf("Failed to read oneTimePassword from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	boxes, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		_, errorCode = protocol.RequestMigration(fmt.Errorf("Failed to read boxes from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestMigration(nil, packet, callID, oneTimePassword, boxes)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
