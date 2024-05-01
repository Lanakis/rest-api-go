package repository

import (
	"authorization/src/modules/user-module/entity"
	"authorization/src/utils"
	"authorization/src/utils/filter"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type UserRepository struct {
	Db *sql.DB
}

/*
func (b *UserRepository) FindByToken(ctx context.Context, token string) {
	tx, _ := b.Db.Begin()
	defer utils.CommitOrRollback(tx)

	SQL := "SELECT id, refresh_token FROM users WHERE id = $1"
	result, _ := tx.QueryContext(ctx, SQL, userId)
}*/

func NewUserRepository(Db *sql.DB) entity.IUserRepository {
	return &UserRepository{Db: Db}
}

func (b *UserRepository) Create(ctx context.Context, user entity.UserEntity) (int, error) {
	tx, err := b.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer utils.CommitOrRollback(tx)

	SQL := "INSERT INTO users (username,password, role) VALUES ($1, $2,$3) RETURNING id"
	var id int
	err1 := tx.QueryRowContext(ctx, SQL, user.Username, user.Password, user.Role).Scan(&id)
	if err1 != nil {
		return 0, err1
	}

	return id, nil
}
func (b *UserRepository) FindAll(ctx context.Context, option []filter.Field) ([]entity.UserEntity, error) {
	tx, err := b.Db.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	// Создаем начальный SQL-запрос без условий
	SQL := "SELECT u.id, u.created_at, u.updated_at, u.username, u.role, p.first_name, p.middle_name, p.last_name, p.age, p.head FROM users u LEFT JOIN profiles p ON u.id = p.user_id WHERE 1=1"
	var args []interface{}

	for _, field := range option {
		switch field.Name {
		case "role":
			operator := "="
			var value string
			if field.Operator == filter.OperatorLike {
				operator = "LIKE"
				fmt.Printf("%s \n", field.Value)

				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				fmt.Printf("%s \n", value)
				args = append(args, `'`+value+`'`)
			} else {
				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				args = append(args, value)
			}
			SQL += fmt.Sprintf(" AND %s %s %s", field.Name, operator, `'`+value+`'`)
		case "username":
			operator := "="
			var value string
			if field.Operator == filter.OperatorLike {
				operator = "LIKE"
				fmt.Printf("%s \n", field.Value)

				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				fmt.Printf("%s \n", value)
				args = append(args, `'`+value+`'`)
			} else {
				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				args = append(args, value)
			}
			SQL += fmt.Sprintf(" AND %s %s %s", field.Name, operator, `'`+value+`'`)
		case "first_name":
			operator := "="
			var value string
			if field.Operator == filter.OperatorLike {
				operator = "LIKE"
				fmt.Printf("%s \n", field.Value)

				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				fmt.Printf("%s \n", value)
				args = append(args, `'`+value+`'`)
			} else {
				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				args = append(args, value)
			}
			SQL += fmt.Sprintf(" AND %s %s %s", field.Name, operator, `'`+value+`'`)
		case "middle_name":
			operator := "="
			var value string
			if field.Operator == filter.OperatorLike {
				operator = "LIKE"
				fmt.Printf("%s \n", field.Value)

				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				fmt.Printf("%s \n", value)
				args = append(args, `'`+value+`'`)
			} else {
				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				args = append(args, value)
			}
			SQL += fmt.Sprintf(" AND %s %s %s", field.Name, operator, `'`+value+`'`)
		case "last_name":
			operator := "="
			var value string
			if field.Operator == filter.OperatorLike {
				operator = "LIKE"
				fmt.Printf("%s \n", field.Value)

				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				fmt.Printf("%s \n", value)
				args = append(args, `'`+value+`'`)
			} else {
				// Заменяем специальные символы в значении поля
				value = utils.EscapeSpecialChars(field.Value)
				args = append(args, value)
			}
			SQL += fmt.Sprintf(" AND %s %s %s", field.Name, operator, `'`+value+`'`)
		case "created_at", "updated_at":
			if field.Operator == filter.OperatorBetween {
				dates := strings.Split(field.Value, ":")
				if len(dates) == 2 {
					startDate := utils.EscapeSpecialChars(dates[0])
					endDate := utils.EscapeSpecialChars(dates[1])
					args = append(args, startDate, endDate)
					SQL += fmt.Sprintf(" AND %s BETWEEN ? AND ?", field.Name)
				}
			} else {
				// Заменяем специальные символы в значении поля
				value := utils.EscapeSpecialChars(field.Value)
				SQL += fmt.Sprintf(" AND %s %s ?", field.Name, field.Operator)
				args = append(args, value)
			}
		}
	}
	fmt.Printf("%s \n", SQL)
	result, errQuery := tx.QueryContext(ctx, SQL)
	if errQuery != nil {
		fmt.Println("Error:", errQuery)
		return nil, errQuery
	}

	fmt.Println(errQuery)

	utils.NewAppError(errQuery, "Can't get query result", "user repository", "")
	defer result.Close()

	var users []entity.UserEntity

	for result.Next() {
		user := entity.UserEntity{BaseEntity: &utils.BaseEntity{}}
		err := result.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Role, &user.Profile.FirstName, &user.Profile.MiddleName, &user.Profile.LastName, &user.Profile.Age, &user.Profile.Head)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (b *UserRepository) FindOne(ctx context.Context, userId int) (entity.UserEntity, error) {
	tx, _ := b.Db.Begin()
	defer utils.CommitOrRollback(tx)

	SQL := "SELECT u.id, u.created_at, u.updated_at, u.username, u.role, p.first_name, p.middle_name, p.last_name, p.age, p.head FROM users u LEFT JOIN profiles p ON u.id = p.user_id WHERE u.id = $1"
	result, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		fmt.Println("Error:", err)

	}

	user := entity.UserEntity{BaseEntity: &utils.BaseEntity{}}
	for result.Next() {
		err := result.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Role, &user.Profile.FirstName, &user.Profile.MiddleName, &user.Profile.LastName, &user.Profile.Age, &user.Profile.Head)
		utils.PanicIfError(err)

	}
	return user, nil
}

func (b *UserRepository) Update(ctx context.Context, user entity.UserEntity, userId int) {
	tx, err := b.Db.Begin()
	defer utils.CommitOrRollback(tx)

	SQL := "UPDATE users SET updated_at = CURRENT_TIMESTAMP, username = $2, password = $3, role = $4 WHERE id = $1"

	_, err = tx.ExecContext(ctx, SQL, userId, user.Username, user.Password, user.Role)
	utils.PanicIfError(err)

}

func (b *UserRepository) Delete(ctx context.Context, userId int) {
	tx, err := b.Db.Begin()
	defer utils.CommitOrRollback(tx)
	SQL := "DELETE FROM users  WHERE id = $1"
	_, err = tx.ExecContext(ctx, SQL, userId)
	utils.PanicIfError(err)
}

func (b *UserRepository) FindByUsername(ctx context.Context, username string) (entity.UserEntity, error) {
	tx, _ := b.Db.Begin()
	defer utils.CommitOrRollback(tx)
	SQL := "SELECT id, username, password FROM users WHERE username = $1"
	result, _ := tx.QueryContext(ctx, SQL, username)
	el := entity.UserEntity{BaseEntity: &utils.BaseEntity{}}
	for result.Next() {
		err := result.Scan(&el.Id, &el.Username, &el.Password)
		utils.PanicIfError(err)

	}

	return el, nil
}
