package memory

type User struct {
	Id       uint32
	Login    string
	Password string
}

func (s *Storage) GetAllUseres() ([]User, error) {
	users := []User{}
	query := `select * from users`
	rows, err := s.ConnectionDB.QueryContext(s.Ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}

		if err := rows.Scan(user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
