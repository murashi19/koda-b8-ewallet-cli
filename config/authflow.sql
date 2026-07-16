DROP TABLE users;
-- DATA USER
CREATE TABLE users(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    name VARCHAR(25) NOT NULL,
    email VARCHAR(25) UNIQUE NOT NULL,
    password VARCHAR(25) NOT NULL,
    status BOOLEAN DEFAULT FALSE
);

SELECT * FROM users;

CREATE TABLE roles(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    role_name VARCHAR(15) NOT NULL DEFAULT user
);

SELECT * FROM roles;

CREATE TABLE user_roles(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    user_id INT NOT NULL REFERENCES users(id),
    role_id INT NOT NULL REFERENCES roles(id)
);

SELECT * FROM user_roles;

ALTER TABLE users DROP COLUMN name;

ALTER TABLE users ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP AND
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

SELECT * FROM users;

CREATE TABLE sessions(
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id BIGINT NOT NULL UNIQUE
    REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expired_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
SELECT * FROM sessions;

CREATE TABLE profiles (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY ,
    user_id BIGINT NOT NULL UNIQUE
    REFERENCES users(id) ON DELETE CASCADE,
    fullname VARCHAR(25) NOT NULL,
    phone_number VARCHAR(13),
    birthday DATE,
    address TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

SELECT * FROM  profiles;


-- Register Flow
INSERT INTO users (email, password, status) VALUES
( 'admin@mail.com', 'admin123', FALSE),
( 'rafli@mail.com', 'rafli123', FALSE),
( 'hani@mail.com', 'hani123', FALSE);

INSERT INTO roles (role_name) VALUES
('Admin'), ('User');

INSERT INTO user_roles (user_id, role_id) VALUES
(1,1), (2,1), (3,2);

INSERT INTO profiles (user_id, fullname, phone_number, birthday, address)
VALUES
(1,'Administrator','081111111111','1999-01-01','Bandung'),
(2,'Muhamad Rafli','082222222222','2003-05-15','Cimahi'),
(3,'Hani','083333333333','2002-10-20','Jakarta');

-- Data User
SELECT p.fullname, u.email, r.role_name FROM users u
JOIN user_roles ur ON u.id = ur.user_id
JOIN roles r ON ur.role_id = r.id
JOIN profiles p ON u.id = p.user_id;

-- Login Flow

SELECT *
FROM users
WHERE email = 'admin@mail.com'
AND password = 'admin123'
AND status = TRUE;
