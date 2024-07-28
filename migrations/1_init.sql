-- +goose Up
CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY ,
    passport_num int NOT NULL,
    passport_series int NOT NULL,
    surname text,
    name text,
    patronymic text,
    address text,
    created_at timestamp NOT NULL
);

CREATE TABLE tasks(
    id BIGSERIAL PRIMARY KEY ,
    name text NOT NULL,
    status boolean,
    created_at  timestamp,
    finished_at timestamp,
    lead_time int,
    user_id BIGINT  NOT NULL
);


ALTER TABLE tasks ADD CONSTRAINT tasks_user_id_fkey foreign key(user_id) REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE tasks DROP CONSTRAINT tasks_user_id_fkey;
DROP TABLE tasks;
DROP TABLE users;