CREATE TABLE user(
 id serial primary key auto_increment,
 username varchar(30) not null unique,
 email varchar(50) not null unique,
 password varchar(30) not null,
 gender varchar(6) not null,
 age int not null,
 address varchar(50) not null,
 created_at timestamp DEFAULT current_timestamp not null,
 updated_at timestamp
)engine = innodb;