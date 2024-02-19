CREATE TABLE IF NOT EXISTS groups (
    group_id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    group_members INT DEFAULT 0 
);


CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    Username VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    group_id INT REFERENCES groups(group_id),
    is_admin BOOLEAN DEFAULT false,
    is_investor BOOLEAN DEFAULT false,
    is_student BOOLEAN DEFAULT false
);
