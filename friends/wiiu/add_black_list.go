package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddBlackList sets the AddBlackList handler function
func (protocol *FriendsWiiUProtocol) AddBlackList(handler func(err error, client *nex.Client, callID uint32, blacklistedPrincipal *BlacklistedPrincipal)) {
	protocol.AddBlackListHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleAddBlackList(packet nex.PacketInterface) {
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

	blacklistedPrincipal, err := parametersStream.ReadStructure(NewBlacklistedPrincipal())
	if err != nil {
		go protocol.AddBlackListHandler(fmt.Errorf("Failed to read blacklistedPrincipal from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.AddBlackListHandler(nil, client, callID, blacklistedPrincipal.(*BlacklistedPrincipal))
}
