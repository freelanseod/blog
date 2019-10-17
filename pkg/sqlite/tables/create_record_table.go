package tables

const CreateRecordTable = `
	DROP TABLE IF EXISTS record;
	CREATE TABLE record (
		id integer PRIMARY KEY AUTOINCREMENT,
		text varchar NOT NULL,
		security integer NOT NULL,
		subject varchar NOT NULL,
		user_id integer NOT NULL,
		created_at varchar NOT NULL,
		updated_at varchar NOT NULL,
		tag_id integer
	);
`
