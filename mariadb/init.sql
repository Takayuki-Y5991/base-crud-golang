CREATE TABLE IF NOT EXISTS departments (
    department_id INT PRIMARY KEY AUTO_INCREMENT,
    department_name varchar(255) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS users (
    user_id int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    user_name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    department_id int NOT NULL,
    FOREIGN KEY(department_id)
    REFERENCES departments(department_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO departments(department_name) VALUES('Accounting');
INSERT INTO departments(department_name) VALUES('Finance');
INSERT INTO departments(department_name) VALUES('Research');
INSERT INTO departments(department_name) VALUES('Planning');

INSERT INTO users(user_name, password,email, department_id) VALUES('user-1','user-1','user-1@example.com','1');
INSERT INTO users(user_name, password,email, department_id) VALUES('user-2','user-2','user-2@example.com','2');
INSERT INTO users(user_name, password,email, department_id) VALUES('user-3','user-3','user-3@example.com','3');
INSERT INTO users(user_name, password,email, department_id) VALUES('user-4','user-4','user-4@example.com','4');
INSERT INTO users(user_name, password,email, department_id) VALUES('user-5','user-5','user-5@example.com','1');

