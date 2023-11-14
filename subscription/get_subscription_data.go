// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetSubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.GetSubscriptionData(nil, packet, callID, nil)
	}

	go protocol.GetSubscriptionData(nil, packet, callID, pids)
}
