// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrivacyLevels sets the GetPrivacyLevels handler function
func (protocol *SubscriptionProtocol) GetPrivacyLevels(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getPrivacyLevelsHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetPrivacyLevels(packet nex.PacketInterface) {
	if protocol.getPrivacyLevelsHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetPrivacyLevels not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getPrivacyLevelsHandler(nil, client, callID)
}