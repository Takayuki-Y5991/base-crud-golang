CREATE TABLE IF NOT EXISTS departments (
    department_id INT PRIMARY KEY AUTOINCREMENT
    department_name varchar(255) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS users (
    user_id int PRIMARY KEY AUTOINCREMENT NOT NULL
    user_name varchar(255) NOT NULL
    email varchar(255) NOT NULL
    department_id int NOT NULL
    FOREIGN KEY(department_id)
    REFERENCES departments(department_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO departments VALUES('Accounting');
INSERT INTO departments VALUES('Finance');
INSERT INTO departments VALUES('Research');
INSERT INTO departments VALUES('Planning');

INSERT INTO users VALUES('user-1','user-1@example.com','1');
INSERT INTO users VALUES('user-2','user-2@example.com','2');
INSERT INTO users VALUES('user-3','user-3@example.com','3');
INSERT INTO users VALUES('user-4','user-4@example.com','4');
INSERT INTO users VALUES('user-5','user-5@example.com','1');

