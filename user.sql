DROP TABLE users;

CREATE TABLE users
(
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	username varchar(64),
	first_name varchar(64),
	last_name varchar(64)
);

INSERT INTO users (username, first_name, last_name) values('eduardo_gpg', 'Eduardo Ismael', 'García Pérez');