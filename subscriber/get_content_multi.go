// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetContentMulti sets the GetContentMulti handler function
func (protocol *Protocol) GetContentMulti(handler func(err error, client *nex.Client, callID uint32, params []*subscriber_types.SubscriberGetContentParam)) {
	protocol.getContentMultiHandler = handler
}

func (protocol *Protocol) handleGetContentMulti(packet nex.PacketInterface) {
	if protocol.getContentMultiHandler == nil {
		globals.Logger.Warning("Subscriber::GetContentMulti not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(subscriber_types.NewSubscriberGetContentParam())
	if err != nil {
		go protocol.getContentMultiHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getContentMultiHandler(nil, client, callID, params.([]*subscriber_types.SubscriberGetContentParam))
}