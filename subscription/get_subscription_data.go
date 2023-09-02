// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSubscriptionData sets the GetSubscriptionData handler function
func (protocol *SubscriptionProtocol) GetSubscriptionData(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	protocol.GetSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetSubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.GetSubscriptionDataHandler(nil, client, callID, nil)
	}

	go protocol.GetSubscriptionDataHandler(nil, client, callID, pids)
}