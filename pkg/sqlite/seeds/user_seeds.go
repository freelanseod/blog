package seeds

const UserSeed = `
	INSERT INTO user (email, password, phone, secret_question, secret_answer, created_at, updated_at) VALUES
	('admin@example.com', 'adminpassword', '+79030001100', 'what is your favourite name', 'testname', '2019-09-21 13:15:58.584604', '2019-09-21 13:15:58.58478');
`
