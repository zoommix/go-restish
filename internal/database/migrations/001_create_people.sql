CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email        character varying(255) NOT NULL,
    username     character varying(255) NOT NULL,
    first_name   character varying(255) NOT NULL,
    last_name    character varying(255) NOT NULL
);

CREATE UNIQUE INDEX index_users_on_lowercase_username ON users((lower(username::text)) text_ops);

---- create above / drop below ----

drop table users;
