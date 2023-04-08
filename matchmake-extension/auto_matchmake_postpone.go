package matchmake_extension

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

// AutoMatchmake_Postpone sets the AutoMatchmake_Postpone handler function
func (protocol *MatchmakeExtensionProtocol) AutoMatchmake_Postpone(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *match_making.MatchmakeSession, message string)) {
	protocol.AutoMatchmake_PostponeHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleAutoMatchmake_Postpone(packet nex.PacketInterface) {
	if protocol.AutoMatchmake_PostponeHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AutoMatchmake_PostponeHandler not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		go protocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtension::AutoMatchmake_Postpone] Data holder name does not match")
		go protocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtension::AutoMatchmake_Postpone] Data holder missing lengths")
		go protocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go protocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, protocol.Server)

	matchmakeSessionStructureInterface, err := dataHolderContentStream.ReadStructure(match_making.NewMatchmakeSession())
	if err != nil {
		globals.Logger.Error(err.Error())
		go protocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*match_making.MatchmakeSession)

	message, err := parametersStream.ReadString()
	if err != nil {
		go protocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	go protocol.AutoMatchmake_PostponeHandler(nil, client, callID, matchmakeSession, message)
}
