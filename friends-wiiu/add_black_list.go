// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddBlackList sets the AddBlackList handler function
func (protocol *Protocol) AddBlackList(handler func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal)) {
	protocol.addBlackListHandler = handler
}

func (protocol *Protocol) handleAddBlackList(packet nex.PacketInterface) {
	if protocol.addBlackListHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AddBlackList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	blacklistedPrincipal, err := parametersStream.ReadStructure(friends_wiiu_types.NewBlacklistedPrincipal())
	if err != nil {
		go protocol.addBlackListHandler(fmt.Errorf("Failed to read blacklistedPrincipal from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.addBlackListHandler(nil, client, callID, blacklistedPrincipal.(*friends_wiiu_types.BlacklistedPrincipal))
}
