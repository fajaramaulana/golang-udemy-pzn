package repository

import (
	"belajarrestfullapi/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, categoryId int) error
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	sql := "select id, name, createdAt, updatedAt from categories order by id desc"

	rows, err := tx.QueryContext(ctx, sql)

	if err != nil {
		return nil, err
	}
	categories := []domain.Category{}

	defer rows.Close()

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "select id, name, createdAt, updatedAt from categories where id = ?"

	rows, err := tx.QueryContext(ctx, sql, categoryId)
	if err != nil {
		log.Print(err.Error())
		return domain.Category{}, err
	}

	defer rows.Close()

	category := domain.Category{}
	if rows.Next() == false {
		return category, errors.New("no data found")
	}

	err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Print(err.Error())
		return domain.Category{}, err
	}

	category.CreatedAt = category.CreatedAt.Local()
	category.UpdatedAt = category.UpdatedAt.Local()

	return category, nil
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	sql := "insert into categories(name, createdAt, updatedAt) values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return category, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return category, err
	}

	category.Id = int(id)
	return category, nil

}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	sql := "UPDATE categories SET name = ?, updatedAt = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.UpdatedAt, category.Id)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) error {
	sql := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, categoryId)
	if err != nil {
		return err
	}

	return nil
}
