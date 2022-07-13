package studentRepositories

import "github.com/jmoiron/sqlx"

type (
	StudentAuthenticationRepositoriesFromDB struct {
		// เพิ่มโครงสร้าง database ตรงนี้
	}

	studentRepositories struct {
		db *sqlx.DB
	}

	StudentRepositories interface {
		AuthenticateBachelor(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoriesFromDB, error)
		AuthenticateMaster(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoriesFromDB, error)
		AuthenticatePhd(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoriesFromDB, error)
	}
)

func NewStudentRepositories(db *sqlx.DB) studentRepositories {
	return studentRepositories{db: db}
}
