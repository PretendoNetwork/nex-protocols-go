// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

// GetContentMulti sets the GetContentMulti handler function
func (protocol *Protocol) GetContentMulti(handler func(err error, packet nex.PacketInterface, callID uint32, params []*subscriber_types.SubscriberGetContentParam) uint32) {
	protocol.getContentMultiHandler = handler
}

func (protocol *Protocol) handleGetContentMulti(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getContentMultiHandler == nil {
		globals.Logger.Warning("Subscriber::GetContentMulti not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(subscriber_types.NewSubscriberGetContentParam())
	if err != nil {
		errorCode = protocol.getContentMultiHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getContentMultiHandler(nil, packet, callID, params.([]*subscriber_types.SubscriberGetContentParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
