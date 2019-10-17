package seeds

const UserSeed = `
	INSERT INTO user (email, password, phone, secret_question, secret_answer, created_at, updated_at) VALUES
	('admin@example.com', 'adminpassword', '+79030001100', 'what is your favourite name', 'testname', '2019-09-21 13:15:58.584604', '2019-09-21 13:15:58.58478'),
	('user@gmail.com', 'testuser', '+79870993212', 'what does cardigans lose', 'favourite game', '2019-10-10 13:15:58.584604', '2019-10-10 13:15:58.58478');
`
