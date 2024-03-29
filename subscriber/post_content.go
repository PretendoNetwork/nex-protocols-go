// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

// PostContent sets the PostContent handler function
func (protocol *Protocol) PostContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) uint32) {
	protocol.postContentHandler = handler
}

func (protocol *Protocol) handlePostContent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.postContentHandler == nil {
		globals.Logger.Warning("Subscriber::PostContent not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(subscriber_types.NewSubscriberPostContentParam())
	if err != nil {
		errorCode = protocol.postContentHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.postContentHandler(nil, packet, callID, param.(*subscriber_types.SubscriberPostContentParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
