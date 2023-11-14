// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetActivePlayerSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetActivePlayerSubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetActivePlayerSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.GetActivePlayerSubscriptionData(nil, packet, callID)
}
