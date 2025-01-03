// This is the API for a TribClient that we have written for you as
// an example. DO NOT MODIFY!

package tribclient

import "github.com/dylanfeehan/tribbler/api/tribrpc"

// TribClient defines the set of methods for one possible Tribbler
// client implementation.
type TribHandle interface {
	CreateUser(userID string) (int, error)
	GetSubscriptions(userID string) ([]string, int, error)
	AddSubscription(userID, targetUser string) (int, error)
	RemoveSubscription(userID, targetUser string) (int, error)
	GetTribbles(userID string) ([]tribrpc.Tribble, int, error)
	GetTribblesBySubscription(userID string) ([]tribrpc.Tribble, int, error)
	PostTribble(userID, contents string) (int, error)
	Close() error
}
