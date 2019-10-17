package tables

const CreateTagsTable = `
	DROP TABLE IF EXISTS tags;
	CREATE TABLE tags (
		id integer PRIMARY KEY AUTOINCREMENT,
		name varchar NOT NULL
	);
`
