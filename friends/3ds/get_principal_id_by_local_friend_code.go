package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrincipalIDByLocalFriendCode sets the GetPrincipalIDByLocalFriendCode handler function
func (protocol *Friends3DSProtocol) GetPrincipalIDByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64)) {
	protocol.GetPrincipalIDByLocalFriendCodeHandler = handler
}

func (protocol *Friends3DSProtocol) HandleGetPrincipalIDByLocalFriendCode(packet nex.PacketInterface) {
	if protocol.GetPrincipalIDByLocalFriendCodeHandler == nil {
		globals.Logger.Warning("Friends3DS::GetPrincipalIDByLocalFriendCode not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.GetPrincipalIDByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.GetPrincipalIDByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.GetPrincipalIDByLocalFriendCodeHandler(nil, client, callID, lfc, lfcList)
}
