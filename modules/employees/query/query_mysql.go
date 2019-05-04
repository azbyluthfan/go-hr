package query

import (
	"errors"
	"github.com/azbyluthfan/go-hr/helper/hash"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type employeeQueryMysql struct {
	db *sqlx.DB
}

// Creates a employeeQueryMysql instance
func NewEmployeeQueryMysql(db *sqlx.DB) *employeeQueryMysql {
	return &employeeQueryMysql{
		db: db,
	}
}

// Get employee by companyId and employeeNo
// Returns error when selected employee is not found
func (q *employeeQueryMysql) GetEmployeeByEmployeeNo(companyId, employeeNo string) (*model.Employee, error) {
	employee := model.Employee{}

	query := `SELECT * FROM employees WHERE company_id=? AND employee_no=? LIMIT 1`

	err := q.db.Get(&employee, query, companyId, employeeNo)

	if err != nil {
		return nil, errors.New("Employee not found")
	}
	return &employee, nil
}

// Record failed login and count its occurence
// Jail for 15 minutes after failed to login 3 times
func (q *employeeQueryMysql) FailedLogin(employeeId string, failedLogin int) error {

	query := `UPDATE employees SET failed_login_count=?`
	if failedLogin == 0 {
		query += `, failed_login_time=NULL, jail_time=NULL`
	} else if failedLogin == 1 {
		query += `, failed_login_time=NOW()`
	} else if failedLogin >= 3 {
		query += `, jail_time=DATE_ADD(NOW(), INTERVAL 15 MINUTE)`
	}
	query += ` WHERE id=?`

	result := q.db.MustExec(query, failedLogin, employeeId)

	_, err := result.RowsAffected()
	if err != nil {
		return errors.New("Failed in updating employee.")
	}

	return nil
}

// Verify password for specified employee
// Check if employee account is locked (jailed until jail time expires) and returns error if it does
// Return Employee object when found and password is correct
func (q *employeeQueryMysql) VerifyPassword(companyId, employeeNo, password string) (*model.Employee, error) {
	employee, err := q.GetEmployeeByEmployeeNo(companyId, employeeNo)
	if err != nil {
		return nil, err
	}

	// check if jailed due to excessive login attempt
	if time.Now().Before(employee.JailTime.Time) {
		return nil, errors.New("Account has been locked due to excessive login. Please try again after " + employee.JailTime.Time.Format(time.RFC3339))
	}

	// get hashed password and salt from encoded password
	hashedPassword, salt, err := hash.DecodeHashedPassword(employee.Password)
	if err != nil {
		return nil, err
	}

	// compare input password encrypted with salt vs hashed password from database
	if hash.HashPasswordWithSalt(password, salt) != hashedPassword {
		err = q.FailedLogin(employee.ID, employee.FailedLogin + 1)
		errorMsg := "Password does not match."
		if err != nil {
			errorMsg += " " + err.Error()
		}
		return nil, errors.New(errorMsg)
	} else {
		err = q.FailedLogin(employee.ID, 0)
		if err != nil {
			return nil, err
		}
	}

	return employee, nil
}

// Check if notice is available for selected time period for certain employee
func (q * employeeQueryMysql) NoticeAvailable(employeeID string, periodStart, periodEnd time.Time) (bool, error) {

	var count int
	err := q.db.Get(&count, `SELECT COUNT(id) FROM notices
		WHERE notices.employee_id=? AND 
		(
			(notices.period_start <= ? AND notices.period_end >= ?) OR 
			(notices.period_start >= ? AND notices.period_end <= ?) OR 
			(notices.period_start <= ? AND notices.period_end >= ?)
		)`, employeeID, periodStart, periodStart, periodStart, periodEnd, periodEnd, periodEnd)

	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}
	return true, nil
}

// Create notice for an employee.
// Can not create notice if there is an overlapping notice for selected period.
func (q *employeeQueryMysql) CreateNotice(
	companyId, employeeNo string,
	noticeType enum.NoticeType,
	visibility enum.NoticeVisibility,
	periodStart, periodEnd time.Time) error {

	employee, err := q.GetEmployeeByEmployeeNo(companyId, employeeNo)
	if err != nil {
		return err
	}

	available, err := q.NoticeAvailable(employee.ID, periodStart, periodEnd)
	if !available {
		return errors.New("Can not create notice. There is overlapping notice(s) for selected period.")
	}
	if err != nil {
		return err
	}

	// insert notice
	query := `INSERT INTO notices (id, employee_id, type, visibility, period_start, period_end) VALUES (UUID(), ?, ?, ?, ?, ?)`
	result := q.db.MustExec(query, employee.ID, noticeType, visibility, periodStart, periodEnd)

	rows, err := result.RowsAffected()
	if rows == 0 || err != nil {
		return errors.New("Failed in creating notice.")
	}

	return nil
}

// Get list of Notice for selected employee
func (q *employeeQueryMysql) GetNotice(employeeId string) ([]*model.Notice, error) {
	notices := []*model.Notice{}
	err := q.db.Select(&notices, `SELECT * FROM notices WHERE employee_id=? ORDER BY period_end DESC`, employeeId)

	if err != nil {
		return nil, err
	}
	return notices, nil
}

// Get list of Notice for employees from selected company
// Filter by visibility
func (q *employeeQueryMysql) GetCompanyNotice(companyId, visibility string) ([]*model.Notice, error) {
	notices := []*model.Notice{}
	query := `
		SELECT notices.* FROM notices
		JOIN employees ON employees.id = notices.employee_id
		WHERE employees.company_id=?`

	if visibility != "" {
		query += ` AND notices.visibility="` + visibility + `"`
	}

	query += ` ORDER BY period_end DESC`

	err := q.db.Select(&notices, query, companyId)

	if err != nil {
		return nil, err
	}
	return notices, nil
}