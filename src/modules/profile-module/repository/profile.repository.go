package repository

import (
	"authorization/src/modules/profile-module/entity"
	"authorization/src/utils"
	"authorization/src/utils/filter"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Repository struct {
	Db *sql.DB
}

func NewProfileRepository(Db *sql.DB) entity.IProfileRepository {
	return &Repository{Db: Db}
}

func (b *Repository) Create(ctx context.Context, profile entity.ProfileEntity, userId int) {

	tx, err := b.Db.Begin()
	if err != nil {

	}
	defer tx.Rollback()

	SQL := "INSERT INTO profiles (first_name, middle_name, last_name, age, head,user_id) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = tx.ExecContext(ctx, SQL, profile.FirstName, profile.MiddleName, profile.LastName, profile.Age, profile.Head, userId)
	if err != nil {
		fmt.Printf("%s \n", err)
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("%s \n", err)
	}

}

func (b *Repository) FindAll(ctx context.Context, option []filter.Field) ([]entity.ProfileEntity, error) {
	tx, err := b.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	SQL := "SELECT id, created_at, updated_at, first_name, middle_name, last_name, age, head FROM profiles WHERE 1=1"
	// Применяем фильтры к запросу
	for _, field := range option {
		fmt.Println("2")
		switch field.Name {
		case "name", "secondName", "patronymic":
			if field.Operator == filter.OperatorLike {
				SQL += fmt.Sprintf(" AND %s %s '%%%s%%'", field.Name, field.Operator, field.Value)
				fmt.Printf("%+v\n IF", SQL)
			} else {
				SQL += fmt.Sprintf(" AND %s %s '%s'", field.Name, field.Operator, field.Value)
				fmt.Printf("%+v\n ELSE", SQL)
			}
		case "age":
			SQL += fmt.Sprintf(" AND %s %s %s", field.Name, field.Operator, field.Value)
		case "created_at", "updated_at":
			if field.Operator == filter.OperatorBetween {
				dates := strings.Split(field.Value, ":")
				if len(dates) == 2 {
					startDate := dates[0]
					endDate := dates[1]
					SQL += fmt.Sprintf(" AND %s BETWEEN '%s' AND '%s'", field.Name, startDate, endDate)
				}
			} else {
				SQL += fmt.Sprintf(" AND %s %s '%s'", field.Name, field.Operator, field.Value)
			}
		case "head":
			head, err := strconv.ParseBool(field.Value)
			if err == nil { // Добавлено условие на отсутствие ошибки
				SQL += fmt.Sprintf(" AND %s = %t", field.Name, head)
			}

		}
	}

	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var profiles []entity.ProfileEntity
	for result.Next() {
		var p entity.ProfileEntity
		err := result.Scan(&p.Id, &p.CreatedAt, &p.UpdatedAt, &p.FirstName, &p.MiddleName, &p.LastName, &p.Age, &p.Head)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}

	return profiles, nil
}

func (b *Repository) FindOne(ctx context.Context, profileId int) (entity.ProfileEntity, error) {
	tx, _ := b.Db.Begin()
	defer utils.CommitOrRollback(tx)
	SQL := "SELECT id, created_at, updated_at, first_name, middle_name, last_name,age,head FROM profiles WHERE id = $1"
	result, _ := tx.QueryContext(ctx, SQL, profileId)

	el := entity.ProfileEntity{BaseEntity: &utils.BaseEntity{}}
	for result.Next() {
		err := result.Scan(&el.Id, &el.CreatedAt, &el.UpdatedAt, &el.FirstName, &el.MiddleName, &el.LastName, &el.Age, &el.Head)
		utils.PanicIfError(err)

	}
	return el, nil
}

func (b *Repository) Update(ctx context.Context, profileId int, profile entity.ProfileEntity) {
	tx, err := b.Db.Begin()
	defer utils.CommitOrRollback(tx)

	SQL := "UPDATE profiles SET updated_at = CURRENT_TIMESTAMP, first_name = $2, middle_name = $3, last_name = $4, age =$5, head =$6 WHERE id = $1"

	_, err = tx.ExecContext(ctx, SQL, profileId, profile.FirstName, profile.MiddleName, profile.LastName, profile.Age, profile.Head)
	utils.PanicIfError(err)

}

func (b *Repository) Delete(ctx context.Context, profileId int) {
	tx, err := b.Db.Begin()
	defer utils.CommitOrRollback(tx)
	SQL := "DELETE FROM profiles  WHERE id = $1"
	_, err = tx.ExecContext(ctx, SQL, profileId)
	utils.PanicIfError(err)
}
