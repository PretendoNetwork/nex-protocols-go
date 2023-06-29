package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPersistentInfo sets the GetFriendPersistentInfo handler function
func (protocol *Friends3DSProtocol) GetFriendPersistentInfo(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.GetFriendPersistentInfoHandler = handler
}

func (protocol *Friends3DSProtocol) handleGetFriendPersistentInfo(packet nex.PacketInterface) {
	if protocol.GetFriendPersistentInfoHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPersistentInfo not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	PidList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.GetFriendPersistentInfoHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetFriendPersistentInfoHandler(nil, client, callID, PidList)
}
