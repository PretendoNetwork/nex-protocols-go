// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrivacyLevels sets the GetPrivacyLevels handler function
func (protocol *SubscriptionProtocol) GetPrivacyLevels(handler func(err error, packet nex.PacketInterface, callID uint32)) {
	protocol.getPrivacyLevelsHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetPrivacyLevels(packet nex.PacketInterface) {
	if protocol.getPrivacyLevelsHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetPrivacyLevels not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getPrivacyLevelsHandler(nil, packet, callID)
}
