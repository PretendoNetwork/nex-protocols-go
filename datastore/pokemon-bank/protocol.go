// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore "github.com/PretendoNetwork/nex-protocols-go/datastore"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Pokemon Bank) protocol
	ProtocolID = 0x73

	// MethodUploadPokemon is the method ID for the UploadPokemon method
	MethodUploadPokemon = 0x29

	// MethodSearchPokemon is the method ID for the SearchPokemon method
	MethodSearchPokemon = 0x2A

	// MethodPrepareTradePokemon is the method ID for the PrepareTradePokemon method
	MethodPrepareTradePokemon = 0x2B

	// MethodTradePokemon is the method ID for the TradePokemon method
	MethodTradePokemon = 0x2C

	// MethodDownloadOtherPokemon is the method ID for the DownloadOtherPokemon method
	MethodDownloadOtherPokemon = 0x2D

	// MethodDownloadMyPokemon is the method ID for the DownloadMyPokemon method
	MethodDownloadMyPokemon = 0x2E

	// MethodDeletePokemon is the method ID for the DeletePokemon method
	MethodDeletePokemon = 0x2F

	// MethodGetTransactionParam is the method ID for the GetTransactionParam method
	MethodGetTransactionParam = 0x30

	// MethodPreparePostBankObject is the method ID for the PreparePostBankObject method
	MethodPreparePostBankObject = 0x31

	// MethodCompletePostBankObject is the method ID for the CompletePostBankObject method
	MethodCompletePostBankObject = 0x32

	// MethodPrepareGetBankObject is the method ID for the PrepareGetBankObject method
	MethodPrepareGetBankObject = 0x33

	// MethodPrepareUpdateBankObject is the method ID for the PrepareUpdateBankObject method
	MethodPrepareUpdateBankObject = 0x34

	// MethodCompleteUpdateBankObject is the method ID for the CompleteUpdateBankObject method
	MethodCompleteUpdateBankObject = 0x35

	// MethodRollbackBankObject is the method ID for the RollbackBankObject method
	MethodRollbackBankObject = 0x36

	// MethodGetUnlockKey is the method ID for the GetUnlockKey method
	MethodGetUnlockKey = 0x37

	// MethodRequestMigration is the method ID for the RequestMigration method
	MethodRequestMigration = 0x38
)

var patchedMethods = []uint32{
	MethodUploadPokemon,
	MethodSearchPokemon,
	MethodPrepareTradePokemon,
	MethodTradePokemon,
	MethodDownloadOtherPokemon,
	MethodDownloadMyPokemon,
	MethodDeletePokemon,
	MethodGetTransactionParam,
	MethodPreparePostBankObject,
	MethodCompletePostBankObject,
	MethodPrepareGetBankObject,
	MethodPrepareUpdateBankObject,
	MethodCompleteUpdateBankObject,
	MethodRollbackBankObject,
	MethodGetUnlockKey,
	MethodRequestMigration,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Pokemon Bank) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	Server *nex.Server
	dataStoreProtocol
	uploadPokemonHandler            func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam)
	searchPokemonHandler            func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationSearchPokemonParam)
	prepareTradePokemonHandler      func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationPrepareTradePokemonParam)
	tradePokemonHandler             func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationTradePokemonParam)
	downloadOtherPokemonHandler     func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadOtherPokemonParam)
	downloadMyPokemonHandler        func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadMyPokemonParam)
	deletePokemonHandler            func(err error, client *nex.Client, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDeletePokemonParam)
	getTransactionParamHandler      func(err error, client *nex.Client, callID uint32, slotID uint16)
	preparePostBankObjectHandler    func(err error, client *nex.Client, callID uint32, slotID uint16, size uint32)
	completePostBankObjectHandler   func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParam)
	prepareGetBankObjectHandler     func(err error, client *nex.Client, callID uint32, slotID uint16, applicationID uint16)
	prepareUpdateBankObjectHandler  func(err error, client *nex.Client, callID uint32, transactionParam *datastore_pokemon_bank_types.BankTransactionParam)
	completeUpdateBankObjectHandler func(err error, client *nex.Client, callID uint32, slotID uint16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce bool)
	rollbackBankObjectHandler       func(err error, client *nex.Client, callID uint32, slotID uint16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce bool)
	getUnlockKeyHandler             func(err error, client *nex.Client, callID uint32, challengeValue uint32)
	requestMigrationHandler         func(err error, client *nex.Client, callID uint32, oneTimePassword string, boxes []uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.dataStoreProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodUploadPokemon:
		go protocol.handleUploadPokemon(packet)
	case MethodSearchPokemon:
		go protocol.handleSearchPokemon(packet)
	case MethodPrepareTradePokemon:
		go protocol.handlePrepareTradePokemon(packet)
	case MethodTradePokemon:
		go protocol.handleTradePokemon(packet)
	case MethodDownloadOtherPokemon:
		go protocol.handleDownloadOtherPokemon(packet)
	case MethodDownloadMyPokemon:
		go protocol.handleDownloadMyPokemon(packet)
	case MethodDeletePokemon:
		go protocol.handleDeletePokemon(packet)
	case MethodGetTransactionParam:
		go protocol.handleGetTransactionParam(packet)
	case MethodPreparePostBankObject:
		go protocol.handlePreparePostBankObject(packet)
	case MethodCompletePostBankObject:
		go protocol.handleCompletePostBankObject(packet)
	case MethodPrepareGetBankObject:
		go protocol.handlePrepareGetBankObject(packet)
	case MethodPrepareUpdateBankObject:
		go protocol.handlePrepareUpdateBankObject(packet)
	case MethodCompleteUpdateBankObject:
		go protocol.handleCompleteUpdateBankObject(packet)
	case MethodRollbackBankObject:
		go protocol.handleRollbackBankObject(packet)
	case MethodGetUnlockKey:
		go protocol.handleGetUnlockKey(packet)
	case MethodRequestMigration:
		go protocol.handleRequestMigration(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStore (Pokemon Bank) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStorePokemonBank protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.dataStoreProtocol.Server = server

	protocol.Setup()

	return protocol
}
