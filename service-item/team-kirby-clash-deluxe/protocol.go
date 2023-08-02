// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item "github.com/PretendoNetwork/nex-protocols-go/service-item"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Service Item (Team Kirby Clash Deluxe) protocol
	ProtocolID = 0x77

	// MethodGetEnvironment is the method ID for the GetEnvironment method
	MethodGetEnvironment = 0x1

	// MethodHTTPGetRequest is the method ID for the HTTPGetRequest method
	MethodHTTPGetRequest = 0x2

	// MethodHTTPGetResponse is the method ID for the HTTPGetResponse method
	MethodHTTPGetResponse = 0x3

	// MethodPurchaseServiceItemRequest is the method ID for the PurchaseServiceItemRequest method
	MethodPurchaseServiceItemRequest = 0x4

	// MethodPurchaseServiceItemResponse is the method ID for the PurchaseServiceItemResponse method
	MethodPurchaseServiceItemResponse = 0x5

	// MethodListServiceItemRequest is the method ID for the ListServiceItemRequest method
	MethodListServiceItemRequest = 0x6

	// MethodListServiceItemResponse is the method ID for the ListServiceItemResponse method
	MethodListServiceItemResponse = 0x7

	// MethodGetBalanceRequest is the method ID for the GetBalanceRequest method
	MethodGetBalanceRequest = 0x8

	// MethodGetBalanceResponse is the method ID for the GetBalanceResponse method
	MethodGetBalanceResponse = 0x9

	// MethodGetPrepurchaseInfoRequest is the method ID for the GetPrepurchaseInfoRequest method
	MethodGetPrepurchaseInfoRequest = 0xA

	// MethodGetPrepurchaseInfoResponse is the method ID for the GetPrepurchaseInfoResponse method
	MethodGetPrepurchaseInfoResponse = 0xB

	// MethodGetServiceItemRightRequest is the method ID for the GetServiceItemRightRequest method
	MethodGetServiceItemRightRequest = 0xC

	// MethodGetServiceItemRightResponse is the method ID for the GetServiceItemRightResponse method
	MethodGetServiceItemRightResponse = 0xD

	// MethodGetPurchaseHistoryRequest is the method ID for the GetPurchaseHistoryRequest method
	MethodGetPurchaseHistoryRequest = 0xE

	// MethodGetPurchaseHistoryResponse is the method ID for the GetPurchaseHistoryResponse method
	MethodGetPurchaseHistoryResponse = 0xF

	// MethodPostRightBinaryByAccount is the method ID for the PostRightBinaryByAccount method
	MethodPostRightBinaryByAccount = 0x10

	// MethodUseServiceItemByAccountRequest is the method ID for the UseServiceItemByAccountRequest method
	MethodUseServiceItemByAccountRequest = 0x11

	// MethodUseServiceItemByAccountResponse is the method ID for the UseServiceItemByAccountResponse method
	MethodUseServiceItemByAccountResponse = 0x12

	// MethodAcquireServiceItemByAccount is the method ID for the AcquireServiceItemByAccount method
	MethodAcquireServiceItemByAccount = 0x13

	// MethodGetSupportID is the method ID for the GetSupportID method
	MethodGetSupportID = 0x14

	// MethodGetLawMessageRequest is the method ID for the GetLawMessageRequest method
	MethodGetLawMessageRequest = 0x15

	// MethodGetLawMessageResponse is the method ID for the GetLawMessageResponse method
	MethodGetLawMessageResponse = 0x16
)

var patchedMethods = []uint32{}

type serviceItemProtocol = service_item.Protocol

// Protocol stores all the RMC method handlers for the Service Item (Team Kirby Clash Deluxe) protocol and listens for requests
// Embeds the Service Item protocol
type Protocol struct {
	Server *nex.Server
	serviceItemProtocol
	getEnvironmentHandler                  func(err error, client *nex.Client, callID uint32, uniqueID string)
	httpGetRequestHandler                  func(err error, client *nex.Client, callID uint32, url *service_item_team_kirby_clash_deluxe_types.ServiceItemHTTPGetParam)
	httpGetResponseHandler                 func(err error, client *nex.Client, callID uint32, requestID uint32)
	purchaseServiceItemRequestHandler      func(err error, client *nex.Client, callID uint32, purchaseServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPurchaseServiceItemParam)
	purchaseServiceItemResponseHandler     func(err error, client *nex.Client, callID uint32, requestID uint32)
	listServiceItemRequestHandler          func(err error, client *nex.Client, callID uint32, listServiceItemParam *service_item_team_kirby_clash_deluxe_types.ServiceItemListServiceItemParam)
	listServiceItemResponseHandler         func(err error, client *nex.Client, callID uint32, requestID uint32)
	getBalanceRequestHandler               func(err error, client *nex.Client, callID uint32, getBalanceParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetBalanceParam)
	getBalanceResponseHandler              func(err error, client *nex.Client, callID uint32, requestID uint32)
	getPrepurchaseInfoRequestHandler       func(err error, client *nex.Client, callID uint32, getPrepurchaseInfoParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPrepurchaseInfoParam)
	getPrepurchaseInfoResponseHandler      func(err error, client *nex.Client, callID uint32, requestID uint32)
	getServiceItemRightRequestHandler      func(err error, client *nex.Client, callID uint32, getServiceItemRightParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetServiceItemRightParam)
	getServiceItemRightResponseHandler     func(err error, client *nex.Client, callID uint32, requestID uint32)
	getPurchaseHistoryRequestHandler       func(err error, client *nex.Client, callID uint32, getPurchaseHistoryParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetPurchaseHistoryParam)
	getPurchaseHistoryResponseHandler      func(err error, client *nex.Client, callID uint32, requestID uint32)
	postRightBinaryByAccountHandler        func(err error, client *nex.Client, callID uint32, postRightBinaryByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam)
	useServiceItemByAccountRequestHandler  func(err error, client *nex.Client, callID uint32, useServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemUseServiceItemByAccountParam)
	useServiceItemByAccountResponseHandler func(err error, client *nex.Client, callID uint32, requestID uint32)
	acquireServiceItemByAccountHandler     func(err error, client *nex.Client, callID uint32, acquireServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemAcquireServiceItemByAccountParam)
	getSupportIDHandler                    func(err error, client *nex.Client, callID uint32, getSuppordIDParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam)
	getLawMessageRequestHandler            func(err error, client *nex.Client, callID uint32, getLawMessageParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetLawMessageParam)
	getLawMessageResponseHandler           func(err error, client *nex.Client, callID uint32, requestID uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.serviceItemProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetEnvironment:
		protocol.handleGetEnvironment(packet)
	case MethodHTTPGetRequest:
		protocol.handleHTTPGetRequest(packet)
	case MethodHTTPGetResponse:
		protocol.handleHTTPGetResponse(packet)
	case MethodPurchaseServiceItemRequest:
		protocol.handlePurchaseServiceItemRequest(packet)
	case MethodPurchaseServiceItemResponse:
		protocol.handlePurchaseServiceItemResponse(packet)
	case MethodListServiceItemRequest:
		protocol.handleListServiceItemRequest(packet)
	case MethodListServiceItemResponse:
		protocol.handleListServiceItemResponse(packet)
	case MethodGetBalanceRequest:
		protocol.handleGetBalanceRequest(packet)
	case MethodGetBalanceResponse:
		protocol.handleGetBalanceResponse(packet)
	case MethodGetPrepurchaseInfoRequest:
		protocol.handleGetPrepurchaseInfoRequest(packet)
	case MethodGetPrepurchaseInfoResponse:
		protocol.handleGetPrepurchaseInfoResponse(packet)
	case MethodGetServiceItemRightRequest:
		protocol.handleGetServiceItemRightRequest(packet)
	case MethodGetServiceItemRightResponse:
		protocol.handleGetServiceItemRightResponse(packet)
	case MethodGetPurchaseHistoryRequest:
		protocol.handleGetPurchaseHistoryRequest(packet)
	case MethodGetPurchaseHistoryResponse:
		protocol.handleGetPurchaseHistoryResponse(packet)
	case MethodPostRightBinaryByAccount:
		protocol.handlePostRightBinaryByAccount(packet)
	case MethodUseServiceItemByAccountRequest:
		protocol.handleUseServiceItemByAccountRequest(packet)
	case MethodUseServiceItemByAccountResponse:
		protocol.handleUseServiceItemByAccountResponse(packet)
	case MethodAcquireServiceItemByAccount:
		protocol.handleAcquireServiceItemByAccount(packet)
	case MethodGetSupportID:
		protocol.handleGetSupportID(packet)
	case MethodGetLawMessageRequest:
		protocol.handleGetLawMessageRequest(packet)
	case MethodGetLawMessageResponse:
		protocol.handleGetLawMessageResponse(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Service Item (Team Kirby Clash Deluxe) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new ServiceItemTeamKirbyClashDeluxe protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.serviceItemProtocol.Server = server

	protocol.Setup()

	return protocol
}
