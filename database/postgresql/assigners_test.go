package postgresql_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/coffemanfp/yofiotest/assigners"
	"github.com/coffemanfp/yofiotest/database/postgresql"
	"github.com/stretchr/testify/assert"
)

var fakeAssigner = assigners.Assigner{
	ID:         1,
	Success:    true,
	Investment: 3000,
}

func TestNewAssignersDB(t *testing.T) {
	aDB := postgresql.NewAssignersDB(&postgresql.PostgreSQL{})
	assert.NotNil(t, aDB)
	assert.Empty(t, aDB)
}

func TestCreate(t *testing.T) {
	assignersDB, mock := newMockDB(t)
	defer assignersDB.DB.Close()

	query := "INSERT INTO (.+) RETURNING id, investment, success"

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(
		fakeAssigner.Investment, fakeAssigner.Success,
	).WillReturnRows(
		sqlmock.NewRows([]string{"id", "investment", "success"}).AddRow(
			fakeAssigner.ID, fakeAssigner.Investment, fakeAssigner.Success,
		),
	)

	newA, err := assignersDB.Create(fakeAssigner)
	assert.Nil(t, err)
	assert.Equal(t, fakeAssigner, newA)
}

func TestFailCreate(t *testing.T) {
	t.Run("error in prepare", func(t *testing.T) {
		assignersDB, mock := newMockDB(t)
		defer assignersDB.DB.Close()

		query := "invalid query"

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery()

		newA, err := assignersDB.Create(fakeAssigner)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to prepare")
		assert.Empty(t, newA)

	})

	t.Run("error in query", func(t *testing.T) {
		assignersDB, mock := newMockDB(t)
		defer assignersDB.DB.Close()

		query := "INSERT INTO (.+) RETURNING id, investment"

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery().WithArgs(
			fakeAssigner.Investment, fakeAssigner.Success,
		).WillReturnRows(
			sqlmock.NewRows([]string{"id", "investment"}).AddRow(
				fakeAssigner.ID, fakeAssigner.Investment,
			),
		)

		newA, err := assignersDB.Create(fakeAssigner)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to query")
		assert.Empty(t, newA)
	})
}

func TestGetStats(t *testing.T) {
	assignersDB, mock := newMockDB(t)
	defer assignersDB.DB.Close()

	totalAsgmtsDone := `SELECT count(id) FROM assignments`
	totalAsgmtsSuccess := `SELECT count(id) FROM assignments WHERE success = TRUE`
	totalAsgmtsFail := `SELECT count(id) FROM assignments WHERE success = FALSE`
	averageInvtSuccess := `SELECT avg(investment)::numeric(12, 1) FROM assignments WHERE success = TRUE`
	averageInvtFail := `SELECT avg(investment)::numeric(12, 1) FROM assignments WHERE success = FALSE`

	rows := mock.NewRows(
		[]string{totalAsgmtsDone, totalAsgmtsSuccess, totalAsgmtsFail, averageInvtSuccess, averageInvtFail},
	).AddRow(1, 1, 0, 3000, 0)

	prep := mock.ExpectPrepare("SELECT (.+)")
	prep.ExpectQuery().WillReturnRows(rows)

	_, err := assignersDB.GetStats()
	assert.Nil(t, err)
}

func TestFailGetStats(t *testing.T) {
	t.Run("error in prepare", func(t *testing.T) {
		assignersDB, mock := newMockDB(t)
		defer assignersDB.DB.Close()

		query := "invalid query"

		prep := mock.ExpectPrepare(query)
		prep.ExpectQuery()

		stats, err := assignersDB.GetStats()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to prepare")
		assert.Empty(t, stats)

	})

	t.Run("error in query", func(t *testing.T) {
		assignersDB, mock := newMockDB(t)
		defer assignersDB.DB.Close()

		totalAsgmtsDone := `SELECT count(id) FROM assignments`

		rows := mock.NewRows(
			[]string{totalAsgmtsDone},
		).AddRow(1)

		prep := mock.ExpectPrepare("SELECT (.+)")
		prep.ExpectQuery().WillReturnRows(rows)

		stats, err := assignersDB.GetStats()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to query")
		assert.Empty(t, stats)
	})

}

func newMockDB(t *testing.T) (aDB *postgresql.AssignersDB, mock sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	aDB = &postgresql.AssignersDB{DB: db}
	return
}
