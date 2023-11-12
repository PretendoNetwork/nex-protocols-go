// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

// GetContent sets the GetContent handler function
func (protocol *Protocol) GetContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) uint32) {
	protocol.getContentHandler = handler
}

func (protocol *Protocol) handleGetContent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getContentHandler == nil {
		globals.Logger.Warning("Subscriber::GetContent not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(subscriber_types.NewSubscriberGetContentParam())
	if err != nil {
		errorCode = protocol.getContentHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getContentHandler(nil, packet, callID, param.(*subscriber_types.SubscriberGetContentParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
