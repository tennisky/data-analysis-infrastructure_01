// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package model

import "database/sql"

func ScanUser(r *sql.Row) (User, error) {
	var s User
	if err := r.Scan(
		&s.ID,
		&s.Name,
		&s.Email,
		&s.Salt,
		&s.Salted,
		&s.Created,
		&s.Updated,
	); err != nil {
		return User{}, err
	}
	return s, nil
}

func ScanUsers(rs *sql.Rows) ([]User, error) {
	structs := make([]User, 0, 16)
	var err error
	for rs.Next() {
		var s User
		if err = rs.Scan(
			&s.ID,
			&s.Name,
			&s.Email,
			&s.Salt,
			&s.Salted,
			&s.Created,
			&s.Updated,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanArticle(r *sql.Row) (Article, error) {
	var s Article
	if err := r.Scan(
		&s.ID,
		&s.Title,
		&s.Body,
		&s.UserID,
		&s.Created,
		&s.Updated,
	); err != nil {
		return Article{}, err
	}
	return s, nil
}

func ScanArticles(rs *sql.Rows) ([]Article, error) {
	structs := make([]Article, 0, 16)
	var err error
	for rs.Next() {
		var s Article
		if err = rs.Scan(
			&s.ID,
			&s.Title,
			&s.Body,
			&s.UserID,
			&s.Created,
			&s.Updated,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanArticleUser(r *sql.Row) (ArticleUser, error) {
	var s ArticleUser
	if err := r.Scan(
		&s.ID,
		&s.Title,
		&s.Body,
		&s.UserID,
		&s.Created,
		&s.Updated,
		&s.UserName,
	); err != nil {
		return ArticleUser{}, err
	}
	return s, nil
}

func ScanArticleUsers(rs *sql.Rows) ([]ArticleUser, error) {
	structs := make([]ArticleUser, 0, 16)
	var err error
	for rs.Next() {
		var s ArticleUser
		if err = rs.Scan(
			&s.ID,
			&s.Title,
			&s.Body,
			&s.UserID,
			&s.Created,
			&s.Updated,
			&s.UserName,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

