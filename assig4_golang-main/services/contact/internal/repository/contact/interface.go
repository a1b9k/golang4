package contact

import (
	"context"

	"github.com/jackc/pgx/v4"

	"Aibek/services/contact/internal/domain/contact"
)

type Contact interface {
	CreateContactTx(ctx context.Context, tx pgx.Tx, contacts ...*contact.Contact) ([]*contact.Contact, error)
}
