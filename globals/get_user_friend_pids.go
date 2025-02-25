package globals

import (
	"context"
	"fmt"

	pbfriends "github.com/PretendoNetwork/grpc-go/friends"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"google.golang.org/grpc/metadata"
)

func GetUserFriendPIDs(pid uint32) []uint32 {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCFriendsCommonMetadata)

	response, err := GRPCFriendsClient.GetUserFriendPIDs(ctx, &pbfriends.GetUserFriendPIDsRequest{Pid: pid})
	if err != nil {
		globals.Logger.Error(err.Error())
		return make([]uint32, 0)
	}
	fmt.Println(response.Pids)
	return response.Pids
}