// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateUserStatus sets the UpdateUserStatus handler function
func (protocol *Protocol) UpdateUserStatus(handler func(err error, client *nex.Client, callID uint32, unknown1 []*subscriber_types.Unknown, unknown2 []uint8)) {
	protocol.updateUserStatusHandler = handler
}

func (protocol *Protocol) handleUpdateUserStatus(packet nex.PacketInterface) {
	if protocol.updateUserStatusHandler == nil {
		globals.Logger.Warning("Subscriber::UpdateUserStatus not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown1, err := parametersStream.ReadListStructure(subscriber_types.NewUnknown())
	if err != nil {
		go protocol.updateUserStatusHandler(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	unknown2, err := parametersStream.ReadListUInt8()
	if err != nil {
		go protocol.updateUserStatusHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.updateUserStatusHandler(nil, client, callID, unknown1.([]*subscriber_types.Unknown), unknown2)
}