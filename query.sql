/** creating database **/
create database dbname;
use dbname;
/** initiating database table schema **/
create table users(
	ID        int auto_increment,
	Name      varchar(30),
	Family    varchar(30),
	Password  varchar(30),
	Email     varchar(30),
	created_at datetime,
    primary key(ID)
)
create table products(
	ID        int auto_increment,
	Name      varchar(30),
	SerialNumber  varchar(30),
	created_at datetime,
    primary key(ID)
);
create table orders(
	ID int auto_increment,
    UserId int,
    ProductId int,
	primary key(ID),
    foreign key(UserId) references users(ID),
    foreign key(ProductId) references products(ID)
);
/*alter table users rename column CreatedAt to created_At;*/
