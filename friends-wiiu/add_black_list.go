// Package friends_wiiu implements the Friends WiiU NEX protocol
package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddBlackList sets the AddBlackList handler function
func (protocol *FriendsWiiUProtocol) AddBlackList(handler func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *friends_wiiu_types.BlacklistedPrincipal)) {
	protocol.AddBlackListHandler = handler
}

func (protocol *FriendsWiiUProtocol) handleAddBlackList(packet nex.PacketInterface) {
	if protocol.AddBlackListHandler == nil {
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
		go protocol.AddBlackListHandler(fmt.Errorf("Failed to read blacklistedPrincipal from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.AddBlackListHandler(nil, client, callID, blacklistedPrincipal.(*friends_wiiu_types.BlacklistedPrincipal))
}
