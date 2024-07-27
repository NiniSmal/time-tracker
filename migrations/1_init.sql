-- +goose Up
CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY ,
    passportNumber text,
    passport_num int NOT NULL,
    passport_series int NOT NULL,
    surname text,
    name text,
    address text,
    created_at timestamp NOT NULL
);

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY ,
    name text NOT NULL,
    status boolean,
    created_at  timestamp,
    finished_at timestamp,
    lead_time int,
    user_id BIGINT  NOT NULL
);

-- +goose Down
DROP TABLE tasks;
DROP TABLE users;