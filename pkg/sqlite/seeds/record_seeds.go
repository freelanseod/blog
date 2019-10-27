package seeds

const RecordSeed = `
	INSERT INTO record (text, security, subject, user_id, created_at, updated_at, tag_id) VALUES
	('Next year, I’ll be entering my 10th year of being formally employed to write code. Ten years! And besides actual employment, for nearly 2⁄3 of my life, I’ve been building things on the web. I can barely remember a time in my life where I didn’t know HTML, which is kind of weird when you think about it. Some kids learn to play an instrument or dance ballet, but instead I was creating magical worlds with code in my childhood bedroom.', 1, '7 absolute truths I unlearned as junior developer', 2, '2019-10-11 13:15:58.584604', '2019-10-11 13:15:58.58478', 3),
	('Currently I am doing a large V-model waterfall project with a three month testing period and 500+ requirements. To track the progress I want to dust off my old “test progress tracking” method that I matured and described in 2011 and 2014.', 0, 'A Track down History', 2, '2019-10-26 13:15:58.584604', '2019-10-26 13:15:58.58478', 1);
`
