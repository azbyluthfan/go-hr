package query

import (
	"errors"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
	"strings"
)

func TestEmployeeQueryMysql_GetEmployeeByEmployeeNo(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)

	columns := []string{"id", "name", "company_id", "role", "employee_no", "password", "failed_login_count", "failed_login_time", "jail_time"}
	rows := sqlmock.NewRows(columns)

	t.Run("Returns employee", func(t *testing.T) {

		rows = rows.AddRow("x", "azby", "1", "admin", "2", "password", 0, time.Now(), time.Now())

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

		employee, _ := q.GetEmployeeByEmployeeNo("1", "2")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if employee.ID != "x" {
			t.Errorf("Incorrect return value, got: %s, want: %s", employee.ID, "x")
		}
	})
	t.Run("Returns error", func(t *testing.T) {

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))

		_, qErr := q.GetEmployeeByEmployeeNo("1", "2")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		if qErr.Error() != "Employee not found" {
			t.Errorf("Incorrect return value, got: %s, want: %s", qErr.Error(), "Employee not found")
		}
	})
}

func TestEmployeeQueryMysql_FailedLogin(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)

	t.Run("failed count reset", func(t *testing.T) {
		mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
		q.FailedLogin("1", 0)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("failed time is set", func(t *testing.T) {
		mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
		q.FailedLogin("1", 1)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("jail time is set", func(t *testing.T) {
		mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
		q.FailedLogin("1", 3)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestEmployeeQueryMysql_VerifyPassword(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)
	columns := []string{"id", "name", "company_id", "role", "employee_no", "password", "failed_login_count", "failed_login_time", "jail_time"}
	rows := sqlmock.NewRows(columns)

	t.Run("Jailed", func(t *testing.T) {

		rows = rows.AddRow("x", "azby", "1", "admin", "2", "password", 0, time.Now(), time.Now().Add(time.Hour))

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

		_, err := q.VerifyPassword("1", "2", "password")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if err == nil || (err != nil && !strings.Contains(err.Error(), "Account has been locked due to excessive login.")) {
			t.Errorf("Should return account has been locked error, got: %s", err.Error())
		}

	})
	t.Run("Bad password", func(t *testing.T) {

		rows = rows.AddRow("x", "azby", "1", "admin", "2", "password", 0, time.Now(), time.Now())

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

		_, err := q.VerifyPassword("1", "2", "password")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if err == nil || (err != nil && !strings.Contains(err.Error(), "Bad hashed password")) {
			t.Errorf("Should return bad hashed password error, got %s", err.Error())
		}

	})
	t.Run("Password does not match", func(t *testing.T) {

		rows = rows.AddRow("x", "azby", "1", "admin", "2", "YmQzMzBmZGJkMzM3ZTk1OGU3NWNmMWM1M2E0ZDU4ZTNkYWI2NzRiMDQ0NDhlMzVkZmQ2MDdiNTMxYWUwNTQzNC5sMzF3Y25tbWR5PQ==", 0, time.Now(), time.Now())

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)
		mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		_, err := q.VerifyPassword("1", "2", "123456")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if err == nil || (err != nil && !strings.Contains(err.Error(), "Password does not match.")) {
			t.Errorf("Should return bad hashed password error, got %s", err.Error())
		}
	})
	t.Run("Password matches", func(t *testing.T) {

		rows = rows.AddRow("x", "azby", "1", "admin", "2", "ZDMyMjQ4ZWQwNzI3YTJlMmYzMGViNjEwNDliNmFhNzI1MWI5ODVhY2E2NjE1ZTE2MGNlNzI5MDBkMjllNGY0Ny52YXpvNW9xZXM5PQ==", 0, time.Now(), time.Now())

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)
		mock.ExpectExec("UPDATE").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		employee, _ := q.VerifyPassword("1", "2", "123456")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if employee.ID != "x" {
			t.Errorf("Incorrect return value, got: %s, want: %s", employee.ID, "x")
		}

	})
}

func TestEmployeeQueryMysql_NoticeAvailable(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)

	columns := []string{"count"}
	rows := sqlmock.NewRows(columns)

	t.Run("Returns error", func(t *testing.T) {

		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errors.New("error"))

		available, err := q.NoticeAvailable("1", time.Now(), time.Now())
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if available || err == nil {
			t.Errorf("Should receive error, got: %v", err)
		}
	})
	t.Run("Returns error due to not available", func(t *testing.T) {

		rows = rows.AddRow(1)
		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(rows)

		available, _ := q.NoticeAvailable("1", time.Now(), time.Now())
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if available {
			t.Errorf("Should receive not available with no error, got: %v", available)
		}

	})
	t.Run("Returns available", func(t *testing.T) {

		rows = rows.AddRow(0)
		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(rows)

		available, _ := q.NoticeAvailable("1", time.Now(), time.Now())
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if !available {
			t.Errorf("Should receive available, got: %v", available)
		}
	})
}

func TestEmployeeQueryMysql_CreateNotice(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)

	columns := []string{"id", "name", "company_id", "role", "employee_no", "password", "failed_login_count", "failed_login_time", "jail_time"}
	rows := sqlmock.NewRows(columns)
	rows = rows.AddRow("x", "azby", "1", "admin", "2", "password", 0, time.Now(), time.Now())
	mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	columns = []string{"count"}
	rows = sqlmock.NewRows(columns)
	rows = rows.AddRow(0)
	mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)

	mock.ExpectExec("INSERT").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	q.CreateNotice("1", "1", enum.SICK, enum.PUBLIC, time.Now(), time.Now())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEmployeeQueryMysql_GetNotice(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)

	columns := []string{"id", "employee_id", "type", "visibility", "period_start", "period_end"}
	rows := sqlmock.NewRows(columns)
	rows = rows.AddRow("1", "1", "sick", "public", time.Now(), time.Now())
	rows = rows.AddRow("2", "1", "sick", "public", time.Now(), time.Now())
	mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	notices, _ := q.GetNotice("1")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if len(notices) != 2 {
		t.Errorf("Should receive two rows, got: %d", len(notices))
	}
}

func TestEmployeeQueryMysql_GetCompanyNotice(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDb := sqlx.NewDb(db, "sqlmock")
	q := NewEmployeeQueryMysql(sqlxDb)

	columns := []string{"id", "employee_id", "type", "visibility", "period_start", "period_end"}
	rows := sqlmock.NewRows(columns)

	t.Run("Admin receives all notices", func(t *testing.T) {
		rows = sqlmock.NewRows(columns)
		rows = rows.AddRow("1", "1", "sick", "public", time.Now(), time.Now())
		rows = rows.AddRow("2", "1", "sick", "private", time.Now(), time.Now())
		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

		notices, _ := q.GetCompanyNotice("1", "")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if len(notices) != 2 {
			t.Errorf("Should receive two rows, got: %d", len(notices))
		}
	})
	t.Run("Admin receives all notices", func(t *testing.T) {
		rows = sqlmock.NewRows(columns)
		rows = rows.AddRow("1", "1", "sick", "public", time.Now(), time.Now())
		mock.ExpectQuery("SELECT").WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

		notices, _ := q.GetCompanyNotice("1", "public")
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if len(notices) != 1 {
			t.Errorf("Should receive two rows, got: %d", len(notices))
		}
	})

}