// Package messaging implements the Messaging protocol
package messaging

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeliverMessageMultiTarget sets the DeliverMessageMultiTarget handler function
func (protocol *MessagingProtocol) DeliverMessageMultiTarget(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.deliverMessageMultiTargetHandler = handler
}

func (protocol *MessagingProtocol) handleDeliverMessageMultiTarget(packet nex.PacketInterface) {
	if protocol.deliverMessageMultiTargetHandler == nil {
		globals.Logger.Warning("MessageDelivery::DeliverMessageMultiTarget not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	// TODO - THIS METHOD HAS AN UNKNOWN REQUEST/RESPONSE FORMAT

	go protocol.deliverMessageMultiTargetHandler(nil, client, callID)
}
