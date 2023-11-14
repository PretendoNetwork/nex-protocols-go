// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetTargetSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetTargetSubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetTargetSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.GetTargetSubscriptionData(nil, packet, callID)
}
