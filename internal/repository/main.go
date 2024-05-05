package repo

import "github.com/jackc/pgx/v4/pgxpool"

type Repo struct {
	IGps
	IParam
	IPoint
}

func New(pgx *pgxpool.Pool) Repo {
	return Repo{
		IGps:   NewGps(pgx),
		IParam: NewParam(pgx),
		IPoint: NewPoint(pgx),
	}
}
