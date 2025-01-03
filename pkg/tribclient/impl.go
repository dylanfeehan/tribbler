package tribclient

import "github.com/dylanfeehan/tribbler/api/tribrpc"

func (th *tribHandle) CreateUser(userID string) (int, error) {
	args := &tribrpc.CreateUserArgs{UserID: &userID}
	reply, err := th.client.CreateUser(th.ctx, args, th.opts)
	if err != nil {
		return 0, err
	}
	return int(reply.GetStatus()), nil
}

func (th *tribHandle) GetSubscriptions(userID string) ([]string, int, error) {
	args := &tribrpc.GetSubscriptionsArgs{UserID: &userID}
	reply, err := th.client.GetSubscriptions(th.ctx, args, th.opts)
	if err != nil {
		return nil, 0, err
	}
	return reply.GetUserIDs(), int(reply.GetStatus()), nil
}

func (th *tribHandle) AddSubscription(userID, targetUserID string) (int, error) {
	args := &tribrpc.SubscriptionArgs{UserID: &userID, TargetUserID: &targetUserID}
	reply, err := th.client.AddSubscription(th.ctx, args, th.opts)
	if err != nil {
		return 0, err
	}
	return int(reply.GetStatus()), nil
}

func (th *tribHandle) RemoveSubscription(userID, targetUserID string) (int, error) {
	args := &tribrpc.SubscriptionArgs{UserID: &userID, TargetUserID: &targetUserID}
	reply, err := th.client.RemoveSubscription(th.ctx, args, th.opts)
	if err != nil {
		return 0, err
	}
	return int(reply.GetStatus()), nil
}

// requires updating client since return is slice to pointer to tribbles
func (th *tribHandle) GetTribbles(userID string) ([]*tribrpc.Tribble, int, error) {
	args := &tribrpc.GetTribblesArgs{UserID: &userID}
	reply, err := th.client.GetTribbles(th.ctx, args, th.opts)
	if err != nil {
		return nil, 0, err
	}
	return reply.GetTribbles(), int(reply.GetStatus()), nil
}

func (th *tribHandle) GetTribblesBySubscription(userID string) ([]*tribrpc.Tribble, int, error) {
	args := &tribrpc.GetTribblesArgs{UserID: &userID}
	reply, err := th.client.GetTribblesBySubscription(th.ctx, args, th.opts)
	if err != nil {
		return nil, 0, err
	}
	return reply.GetTribbles(), int(reply.GetStatus()), nil
}

func (th *tribHandle) PostTribble(userID, contents string) (int, error) {
	args := &tribrpc.PostTribbleArgs{UserID: &userID, Contents: &contents}
	reply, err := th.client.PostTribble(th.ctx, args, th.opts)
	if err != nil {
		return 0, err
	}
	return int(reply.GetStatus()), nil
}
