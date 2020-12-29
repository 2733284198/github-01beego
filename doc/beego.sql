

create table articles
(
    id bigint auto_increment
        primary key,
    title varchar(100) not null,
    content longtext null
);

create index article_index
	on articles (title);

create table manager
(
    id int auto_increment
        primary key,
    username varchar(255) null,
    password varchar(32) null,
    mobile varchar(11) null,
    email varchar(255) null,
    status tinyint(1) null,
    role_id int null,
    add_time int null,
    is_super tinyint(1) default 0 null
)
    charset=utf8;

create table person
(
    id int auto_increment
        primary key,
    age int null,
    name varchar(255) null
)
    charset=utf8;

create table user2
(
    id int auto_increment,
    username varchar(255) default '' not null,
    password varchar(255) default '' not null,
    age int default 0 null,
    primary key (id, username, password)
);

create table users
(
    id int auto_increment,
    username varchar(255) default '' not null,
    password varchar(255) default '' not null,
    age int default 0 null,
    primary key (id, username, password)
);