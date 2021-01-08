package postgresql

import (
	"database/sql"

	"github.com/coffemanfp/yofiotest/assigners"
	"github.com/coffemanfp/yofiotest/database"
)

// NewAssignersDB returns a new assigners database instance.
func NewAssignersDB(p *PostgreSQL) database.AssignersDB {
	return AssignersDB{
		DB: p.GetConn(),
	}
}

// AssignersDB represents the database assigners interactions.
type AssignersDB struct {
	DB *sql.DB
}

// Create creates the assigner register.
func (adb AssignersDB) Create(a assigners.Assigner) (newA assigners.Assigner, err error) {
	query := `
		INSERT INTO
			assignments(investment, success)
		VALUES
			($1, $2)
		RETURNING
			id, investment, success
	`

	stmt, err := adb.DB.Prepare(query)
	if err != nil {
		err = errorPrepare(err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		a.Investment,
		a.Success,
	).Scan(
		&newA.ID,
		&newA.Investment,
		&newA.Success,
	)
	if err != nil {
		err = errorQuery(err)
	}
	return
}

type nullStats struct {
	totalAsgmtsDone             *sql.NullInt64
	totalAsgmtsSuccess          *sql.NullInt64
	totalAsgmtsFail             *sql.NullInt64
	averageInvestmentSuccessful *sql.NullFloat64
	averageInvestmentFail       *sql.NullFloat64
}

func (n nullStats) getValues(s *assigners.Stats) {
	if n.totalAsgmtsDone != nil {
		s.TotalAsgmtsDone = int(n.totalAsgmtsDone.Int64)
	}
	if n.totalAsgmtsSuccess != nil {
		s.TotalAsgmtsSuccess = int(n.totalAsgmtsSuccess.Int64)
	}
	if n.totalAsgmtsFail != nil {
		s.TotalAsgmtsFail = int(n.totalAsgmtsFail.Int64)
	}
	if n.averageInvestmentSuccessful != nil {
		s.AverageInvestmentSuccessful = n.averageInvestmentSuccessful.Float64
	}
	if n.averageInvestmentFail != nil {
		s.AverageInvestmentFail = n.averageInvestmentFail.Float64
	}
}

// GetStats get the assigners stats.
func (adb AssignersDB) GetStats() (s assigners.Stats, err error) {
	query := `
		SELECT
			(
				SELECT
					count(id)
				FROM
					assignments
			),
			(
				SELECT
					count(id)
				FROM
					assignments
				WHERE
					success = TRUE
			),
			(
				SELECT
					count(id)
				FROM
					assignments
				WHERE
					success = FALSE
			),
			(
				SELECT
					avg(investment)::numeric(12, 1)
				FROM
					assignments
				WHERE success = TRUE
			),
			(
				SELECT
					avg(investment)::numeric(12, 1)
				FROM
					assignments
				WHERE
					success = FALSE
			);
	`

	stmt, err := adb.DB.Prepare(query)
	if err != nil {
		err = errorPrepare(err)
		return
	}
	defer stmt.Close()

	var nullValues nullStats

	err = stmt.QueryRow().Scan(
		&nullValues.totalAsgmtsDone,
		&nullValues.totalAsgmtsSuccess,
		&nullValues.totalAsgmtsFail,
		&nullValues.averageInvestmentSuccessful,
		&nullValues.averageInvestmentFail,
	)
	if err != nil {
		err = errorQuery(err)
		return
	}

	nullValues.getValues(&s)
	return
}
