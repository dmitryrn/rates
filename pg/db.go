package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"time"
)

type BtcUsdtRate struct {
	value decimal.Decimal
	createdAt time.Time
}

type RatesRepository struct {
	conn *pgxpool.Pool
}

func NewRatesRepository(connStr string) (*RatesRepository, error) {
    instance := &RatesRepository{}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	instance.conn = conn
	return instance, nil
}

func (r *RatesRepository) GetBtcUsdtRate() []BtcUsdtRate {
	r.conn.Query()
}

// TODO: add migrations later maybe
func (r *RatesRepository) createSchema(ctx context.Context) error {
	_, err := r.conn.Exec(ctx, `
create table IF NOT EXISTS btc_usd (
    id serial primary key,
    created_at timestamp, -- no timezone
    value numeric -- no precision, it's gonna store only what is needed
)
`)
	if err != nil {
		return err
	}


}