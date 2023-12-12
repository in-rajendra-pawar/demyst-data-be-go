DROP TABLE IF EXISTS tasks CASCADE;

CREATE TABLE users
(
    id           int IDENTITY(1,1)           PRIMARY KEY,
    first_name   VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
    last_name    VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_at     TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks
(
    id           int IDENTITY(1,1)           PRIMARY KEY,
    user         int FOREIGN KEY REFERENCES  users(id),
    title        VARCHAR(32)                 NOT NULL CHECK ( title <> '' ), 
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    updated_by   int FOREIGN KEY REFERENCES  users(id)
);