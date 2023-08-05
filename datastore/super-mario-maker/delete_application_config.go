// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteApplicationConfig sets the DeleteApplicationConfig handler function
func (protocol *Protocol) DeleteApplicationConfig(handler func(err error, client *nex.Client, callID uint32, applicationID uint32, key uint32) uint32) {
	protocol.deleteApplicationConfigHandler = handler
}

func (protocol *Protocol) handleDeleteApplicationConfig(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteApplicationConfigHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteApplicationConfig not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.deleteApplicationConfigHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	key, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.deleteApplicationConfigHandler(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), client, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteApplicationConfigHandler(nil, client, callID, applicationID, key)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
