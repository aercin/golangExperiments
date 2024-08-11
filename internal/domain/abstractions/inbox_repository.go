package abstractions

import (
	"context"
)

type InboxRepository interface {
	Any(ctx context.Context, messageId string) bool
}
