// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPrivacyLevels(packet nex.PacketInterface) {
	if protocol.GetPrivacyLevels == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetPrivacyLevels not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.GetPrivacyLevels(nil, packet, callID)
}
