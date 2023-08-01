// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostContent sets the PostContent handler function
func (protocol *Protocol) PostContent(handler func(err error, client *nex.Client, callID uint32, param *subscriber_types.SubscriberPostContentParam)) {
	protocol.postContentHandler = handler
}

func (protocol *Protocol) handlePostContent(packet nex.PacketInterface) {
	if protocol.postContentHandler == nil {
		globals.Logger.Warning("Subscriber::PostContent not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(subscriber_types.NewSubscriberPostContentParam())
	if err != nil {
		go protocol.postContentHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.postContentHandler(nil, client, callID, param.(*subscriber_types.SubscriberPostContentParam))
}