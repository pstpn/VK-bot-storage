DROP TABLE IF EXISTS users_data;

CREATE TABLE users_data
(
    id           serial NOT NULL,
    chat_id      integer NOT NULL,
    service_name text NOT NULL,
    login        text NOT NULL,
    password     text NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (chat_id, service_name)
);