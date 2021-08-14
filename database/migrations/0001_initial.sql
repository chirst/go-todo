CREATE TABLE IF NOT EXISTS public.user(
    id          serial NOT NULL,
    username    varchar(256) UNIQUE NOT NULL,
    password    varchar(256) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS todo(
    id          serial NOT NULL,
    name        varchar(256) NOT NULL,
    completed   timestamp,
    deleted     timestamp,
    user_id     integer REFERENCES public.user NOT NULL,
    PRIMARY KEY (id)
);
