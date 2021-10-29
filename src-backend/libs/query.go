package libs

func InsertUserForm(openId, name, mail string) error {
	query := `INSERT INTO users (user_open_id, 
		name,
		mail) VALUES ($1,
			$2,
			$3)`
	err := Conn.Exec(query, openId, name, mail)
	if err != nil {
		return err
	}
	return nil
}