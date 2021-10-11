// Code generated by sqlc. DO NOT EDIT.
// source: indications.sql

package db

import (
	"context"
	"time"
)

const createIndication = `-- name: CreateIndication :exec
INSERT INTO indications (
    indication, controller_serial, sent_at
) VALUES (
             $1, $2, $3
         )
`

type CreateIndicationParams struct {
	Indication       string    `json:"indication"`
	ControllerSerial string    `json:"controllerSerial"`
	SentAt           time.Time `json:"sentAt"`
}

func (q *Queries) CreateIndication(ctx context.Context, arg CreateIndicationParams) error {
	_, err := q.db.ExecContext(ctx, createIndication, arg.Indication, arg.ControllerSerial, arg.SentAt)
	return err
}

const listIndications = `-- name: ListIndications :many
SELECT id, indication, controller_serial, sent_at, created_at FROM indications
ORDER BY sent_at
`

func (q *Queries) ListIndications(ctx context.Context) ([]Indication, error) {
	rows, err := q.db.QueryContext(ctx, listIndications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Indication
	for rows.Next() {
		var i Indication
		if err := rows.Scan(
			&i.ID,
			&i.Indication,
			&i.ControllerSerial,
			&i.SentAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listIndicationsByController = `-- name: ListIndicationsByController :many
SELECT id, indication, controller_serial, sent_at, created_at FROM indications
WHERE controller_serial = $1
ORDER BY sent_at
`

func (q *Queries) ListIndicationsByController(ctx context.Context, controllerSerial string) ([]Indication, error) {
	rows, err := q.db.QueryContext(ctx, listIndicationsByController, controllerSerial)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Indication
	for rows.Next() {
		var i Indication
		if err := rows.Scan(
			&i.ID,
			&i.Indication,
			&i.ControllerSerial,
			&i.SentAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
