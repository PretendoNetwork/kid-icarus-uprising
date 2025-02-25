package globals

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/plogger-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pbaccount "github.com/PretendoNetwork/grpc-go/account"
	pbfriends "github.com/PretendoNetwork/grpc-go/friends"
)

var Logger *plogger.Logger
var KerberosPassword = "password" // * Default password
var AuthenticationServer *nex.PRUDPServer
var AuthenticationEndpoint *nex.PRUDPEndPoint
var SecureServer *nex.PRUDPServer
var SecureEndpoint *nex.PRUDPEndPoint
var GRPCAccountClientConnection *grpc.ClientConn
var GRPCAccountClient pbaccount.AccountClient
var GRPCAccountCommonMetadata metadata.MD
var GRPCFriendsCommonMetadata metadata.MD
var GRPCFriendsClientConnection *grpc.ClientConn
var GRPCFriendsClient pbfriends.FriendsClient
