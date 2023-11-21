// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReplaceTargetAndGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.ReplaceTargetAndGetSubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::ReplaceTargetAndGetSubscriptionData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.ReplaceTargetAndGetSubscriptionData(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
