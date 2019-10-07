package tables

const CreateUserTable = `
	DROP TABLE IF EXISTS user;
	CREATE TABLE user (
		id integer PRIMARY KEY AUTOINCREMENT,
		email varchar NOT NULL,
		password varchar NOT NULL,
		phone varchar 	NOT NULL,
		secret_question varchar NOT NULL,
		secret_answer varchar 	NOT NULL,
		created_at varchar NOT NULL,
		updated_at varchar NOT NULL
	);
`
