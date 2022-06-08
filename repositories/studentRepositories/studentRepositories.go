package studentRepositories

//6256000792 11/1/2534 รูปแบบ วว/ดด/ปปปป
func (r studentRepositories) AuthenticateBachelor(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoriesFromDB, error) {

	student := StudentAuthenticationRepositoriesFromDB{}

	query := ""
	err := r.db.Get(&student, query, std_code, birth_date)
	if err != nil {
		return nil, err
	}

	return &student, nil

}

//6014832050 17/10/2500 รูปแบบ วว/ดด/ปปปป
func (r studentRepositories) AuthenticateMaster(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoriesFromDB, error) {

	student := StudentAuthenticationRepositoriesFromDB{}

	query := ""
	err := r.db.Get(&student, query, std_code, birth_date)
	if err != nil {
		return nil, err
	}

	return &student, nil

}


//6014832050 17/10/2500 รูปแบบ วว/ดด/ปปปป
func (r studentRepositories) AuthenticatePhd(std_code string, birth_date string, lev_id string) (*StudentAuthenticationRepositoriesFromDB, error) {

	student := StudentAuthenticationRepositoriesFromDB{}

	query := ""
	err := r.db.Get(&student, query, std_code, birth_date)
	if err != nil {
		return nil, err
	}

	return &student, nil

}
