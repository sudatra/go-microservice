package account

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*postgresRepository, error) {
	db, err := sql.Open("postgres", url);
	if err != nil {
		return nil, err
	}

	err = db.Ping();
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (pgr *postgresRepository) Close() {
	pgr.db.Close();
}

func (pgr *postgresRepository) Ping() {
	pgr.db.Ping();
}

func (pgr *postgresRepository) PutAccount(ctx context.Context, a Account) error {
	_, err := pgr.db.ExecContext(ctx, "INSERT INTO accounts(id, name) VALUES ($1, $2)", a.ID, a.Name);
	return err;
}

func (pgr *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	row := pgr.db.QueryRowContext(ctx, "SELECT id, name FROM accounts WHERE id = $1", id);
	a := &Account{};

	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}

	return a, nil;
}

func (pgr *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	rows, err := pgr.db.QueryContext(ctx, "SELECT id, name FROM accounts ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take);
	if err != nil {
		return nil, err
	}

	defer rows.Close();
	accounts := []Account{};

	for rows.Next() {
		a := &Account{};
		if err := rows.Scan(&a.ID, &a.Name); err == nil {
			accounts = append(accounts, *a)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}