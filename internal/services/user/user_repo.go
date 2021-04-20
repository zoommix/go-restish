package user

import "context"

func (s *Service) GetAllUsers(limit, skip int64) ([]User, error) {
	users := make([]User, 0)

	rows, err := s.DB.Query(context.Background(), getAllUsersSQL, limit, skip)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}

		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Username); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	rows.Close()

	return users, nil
}
