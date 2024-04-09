package likes

import "github.com/gofrs/uuid"

const UPDATEINTERACTIONSQUERY = ``

func RecordInteractionEvent(postid uuid.UUID, userid uuid.UUID, interaction bool) error {
	return nil
}