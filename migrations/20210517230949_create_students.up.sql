CREATE TABLE IF NOT EXISTS students
(
    id          bigserial not null primary key,
    fullname    varchar   not null,
    faculty     varchar   not null,
    course      smallint  not null
);

