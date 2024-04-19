package memory

import "github.com/SashaMelva/tourism/internal/storage/model"

func (s *Storage) GetAllUseres() ([]model.User, error) {
	users := []model.User{}
	query := `select * from users`
	rows, err := s.ConnectionDB.QueryContext(s.Ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}

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

func (s *Storage) CreateUser(user model.User) (uint32, error) {
	var id uint32
	query := `insert into users(login, password) values($1, $2) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, user.Login) // sql.Result
	err := result.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
