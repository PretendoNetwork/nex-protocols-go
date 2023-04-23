package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetProfiles sets the GetProfiles handler function
func (protocol *DataStoreSuperSmashBros4Protocol) GetProfiles(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	protocol.GetProfilesHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleGetProfiles(packet nex.PacketInterface) {
	if protocol.GetProfilesHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::GetProfiles not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidList := parametersStream.ReadListUInt32LE()

	go protocol.GetProfilesHandler(nil, client, callID, pidList)
}