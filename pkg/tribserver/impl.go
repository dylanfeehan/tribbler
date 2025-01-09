package tribserver

import (
	"context"
	"fmt"
	"time"

	"github.com/dylanfeehan/tribbler/api/tribrpc"
)

func (ts tribServer) AddSubscription(cc context.Context, args *tribrpc.SubscriptionArgs) (*tribrpc.SubscriptionReply, error) {
	return nil, nil
}

func (ts tribServer) CreateUser(cc context.Context, args *tribrpc.CreateUserArgs) (*tribrpc.CreateUserReply, error) {
	return nil, nil
}

func (ts tribServer) GetSubscriptions(cc context.Context, args *tribrpc.GetSubscriptionsArgs) (*tribrpc.GetSubscriptionsReply, error) {
	return nil, nil
}

func (ts tribServer) GetTribbles(cc context.Context, args *tribrpc.GetTribblesArgs) (*tribrpc.GetTribblesReply, error) {
	reply := new(tribrpc.GetTribblesReply)
	reply.Tribbles = ts.Tribs[*args.UserID]
	var status int32 = 0
	reply.Status = &status
	return reply, nil
}

func (tribServer) GetTribblesBySubscription(cc context.Context, args *tribrpc.GetTribblesArgs) (*tribrpc.GetTribblesReply, error) {
	return nil, nil
}

func (ts tribServer) PostTribble(cc context.Context, args *tribrpc.PostTribbleArgs) (*tribrpc.PostTribbleReply, error) {
	now := time.Now().Unix()
	fmt.Println(now)
	tribble := &tribrpc.Tribble{
		UserID:     args.UserID,
		Contents:   args.Contents,
		PostedTime: &now,
	}

	userTribs := ts.Tribs[*args.UserID]
	userTribs = append(userTribs, tribble)
	ts.Tribs[*args.UserID] = userTribs
	var status int32 = 0
	reply := &tribrpc.PostTribbleReply{
		Status: &status,
	}
	return reply, nil
}

func (ts tribServer) RemoveSubscription(cc context.Context, args *tribrpc.SubscriptionArgs) (*tribrpc.SubscriptionReply, error) {
	return nil, nil
}
