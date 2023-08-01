// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetContent sets the GetContent handler function
func (protocol *Protocol) GetContent(handler func(err error, client *nex.Client, callID uint32, param *subscriber_types.SubscriberGetContentParam)) {
	protocol.getContentHandler = handler
}

func (protocol *Protocol) handleGetContent(packet nex.PacketInterface) {
	if protocol.getContentHandler == nil {
		globals.Logger.Warning("Subscriber::GetContent not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(subscriber_types.NewSubscriberGetContentParam())
	if err != nil {
		go protocol.getContentHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getContentHandler(nil, client, callID, param.(*subscriber_types.SubscriberGetContentParam))
}