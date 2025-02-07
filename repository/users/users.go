package users_repo

import (
	config "api_study/init"
	"api_study/interfaces"
)

func CreateUser(data *interfaces.UserRequest) (interfaces.UserRequest, error) {
	var user interfaces.UserRequest

	config.DB.Statement.Raw(`
		INSERT INTO "user"
		(
			email,
			fullname,
			user_picture,
			password
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) RETURNING
			user_id,
			email,
			fullname,
			user_picture,
			created_at
	`,
		data.Email,
		data.Fullname,
		data.UserPicture,
		data.Password,
	).Scan(&user)

	return user, nil
}

func UpdateUser(data *interfaces.UserRequest) (interfaces.UserRequest, error) {
	var user interfaces.UserRequest

	config.DB.Statement.Raw(`
		UPDATE "user" SET
			email=$1,
			fullname=$2,
			user_picture=$3
		WHERE user_id = $4
		RETURNING
			user_id,
			email,
			fullname,
			user_picture,
			created_at
	`,
		data.Email,
		data.Fullname,
		data.UserPicture,
		data.UserID).Scan(&user)

	return user, nil
}

func RetrieveUserByID(userId int) (interfaces.UserResponse, error) {
	var user interfaces.UserResponse

	config.DB.Statement.Raw(`
		SELECT
			user_id,
			email,
			fullname,
			user_picture,
			created_at,
			updated_at
		FROM "user"
		WHERE user_id = $1 AND deleted_at is null
	`, userId).Scan(&user)

	return user, nil
}

func RetrieveUserByEmail(email string) (interfaces.User, error) {
	var user interfaces.User

	config.DB.Statement.Raw(`
		SELECT
			*
		FROM "user"
		WHERE email = $1 AND deleted_at is null
	`, email).Scan(&user)

	return user, nil
}

func ListUsers() []interfaces.UserResponse {
	var users = []interfaces.UserResponse{}

	config.DB.Statement.Raw(`SELECT * FROM "user" WHERE deleted_at is null`).Scan(&users)

	return users
}

func DeleteUser(userId int) error {
	// Exec - Immediately executes the query
	config.DB.Statement.Exec(`
		UPDATE "user" SET deleted_at = NOW() WHERE user_id = $1 AND deleted_at is null
	`, userId)

	return nil
}
