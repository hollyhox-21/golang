package mysql

import (
	"database/sql"
	"errors"
	"github.com/hollyhox-21/notpad/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *SnippetModel) Delete(id int) error {
	stmt := `DELETE FROM snippets WHERE id = ?`

	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	//_, err := result.LastInsertId()
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *SnippetModel) Get(id int) (*models.Snipped, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := s.DB.QueryRow(stmt, id)

	sn := &models.Snipped{}
	err := row.Scan(&sn.ID, &sn.Title, &sn.Content, &sn.Created, &sn.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrorNoRecord
		} else {
			return nil, err
		}
	}
	return sn, nil
}

func (s *SnippetModel) Latest() ([]*models.Snipped, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*models.Snipped

	for rows.Next() {
		sn := &models.Snipped{}
		err := rows.Scan(&sn.ID, &sn.Title, &sn.Content, &sn.Created, &sn.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, sn)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}