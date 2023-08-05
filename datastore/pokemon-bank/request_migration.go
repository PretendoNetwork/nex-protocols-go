// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestMigration sets the RequestMigration handler function
func (protocol *Protocol) RequestMigration(handler func(err error, client *nex.Client, callID uint32, oneTimePassword string, boxes []uint32) uint32) {
	protocol.requestMigrationHandler = handler
}

func (protocol *Protocol) handleRequestMigration(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.requestMigrationHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::RequestMigration not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oneTimePassword, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.requestMigrationHandler(fmt.Errorf("Failed to read oneTimePassword from parameters. %s", err.Error()), client, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	boxes, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.requestMigrationHandler(fmt.Errorf("Failed to read boxes from parameters. %s", err.Error()), client, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.requestMigrationHandler(nil, client, callID, oneTimePassword, boxes)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
