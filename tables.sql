CREATE TABLE categories (
    id string PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255)
);

CREATE TABLE courses (
    id string PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255),
    category_id string,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);