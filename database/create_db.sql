create table users (
    id integer primary key generated by default as identity,
    login varchar(30),
    password varchar(30),
    role varchar(15)
);

insert into users values(default, 'avalon@jk.com', 'test1234', 'admin');
insert into users values(default, 'chris@jk.com', 'test1234', 'inspector');
insert into users values(default, 'taker@jk.com', 'test1234', 'worker');

create table brigade_log(
    id              integer     primary key generated by default as identity,
    user_id         integer     references users(id),
    execution_date  timestamp,
    long            float8,
    lat             float8,
    point_address   varchar(150),
    number_arc      integer,
    service_type    varchar(50),
    subtype         varchar(50),
    photo_before    varchar(500),
    photo_left      varchar(300),
    photo_right     varchar(300),
    photo_front     varchar(300),
    video           varchar(300),
    comment         varchar(500),


)


