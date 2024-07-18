package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// Repository defines generic CRUD operations for any entity type T.
type Repository[T any] interface {
	Create(entity *T) error
	GetAll() ([]*T, error)
	GetByID(id int) (*T, error)
	Update(entity *T) error
}

// genericRepository implements the Repository interface with generics.
type genericRepository[T any] struct {
	db        *sql.DB
	tableName string
	columns   []string
}

// NewGenericRepository creates a new instance of genericRepository.
func NewGenericRepository[T any](db *sql.DB, tableName string, columns []string) Repository[T] {
	return &genericRepository[T]{db: db, tableName: tableName, columns: columns}
}

func (r *genericRepository[T]) Create(entity *T) error {
	values := make([]interface{}, len(r.columns))
	for i, col := range r.columns {
		field := reflect.ValueOf(entity).Elem().FieldByName(strings.Title(col))
		values[i] = field.Interface()
	}

	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", r.tableName, strings.Join(r.columns, ", "), strings.Join(placeholders, ", "))
	_, err := r.db.Exec(query, values...)
	return err
}

func (r *genericRepository[T]) GetAll() ([]*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(r.columns, ", "), r.tableName)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*T
	for rows.Next() {
		entity := new(T)
		values := make([]interface{}, len(r.columns))
		for i := range values {
			// Check if the field exists in the struct before accessing it
			field := reflect.ValueOf(entity).Elem().Field(i)
			if field.IsValid() && field.CanSet() {
				values[i] = field.Interface()
			} else {
				// Handle the case where the field doesn't exist or cannot be set
				return nil, fmt.Errorf("field %s not found or cannot be set", r.columns[i])
			}
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func (r *genericRepository[T]) GetByID(id int) (*T, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", strings.Join(r.columns, ", "), r.tableName)
	row := r.db.QueryRow(query, id)

	entity := new(T)
	values := make([]interface{}, len(r.columns))
	for i := range values {
		field := reflect.ValueOf(entity).Elem().Field(i).Addr().Interface()
		values[i] = field
	}

	if err := row.Scan(values...); err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *genericRepository[T]) Update(entity *T) error {
	setClauses := make([]string, len(r.columns))
	values := make([]interface{}, len(r.columns))
	for i, col := range r.columns {
		setClauses[i] = fmt.Sprintf("%s = $%d", col, i+1)
		field := reflect.ValueOf(entity).Elem().FieldByName(strings.Title(col))
		values[i] = field.Interface()
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", r.tableName, strings.Join(setClauses, ", "), len(values)+1)
	idField := reflect.ValueOf(entity).Elem().FieldByName("ID").Interface()
	values = append(values, idField)

	_, err := r.db.Exec(query, values...)
	return err
}
