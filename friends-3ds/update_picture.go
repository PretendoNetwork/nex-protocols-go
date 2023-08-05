// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePicture sets the UpdatePicture handler function
func (protocol *Protocol) UpdatePicture(handler func(err error, client *nex.Client, callID uint32, unknown uint32, picture []byte) uint32) {
	protocol.updatePictureHandler = handler
}

func (protocol *Protocol) handleUpdatePicture(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updatePictureHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePicture not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.updatePictureHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	picture, err := parametersStream.ReadBuffer()
	if err != nil {
		errorCode = protocol.updatePictureHandler(fmt.Errorf("Failed to read picture from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updatePictureHandler(nil, client, callID, unknown, picture)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
