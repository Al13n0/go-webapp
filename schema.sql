create table books(
    id serial,
    title varchar,
    author varchar,
    year varchar
);

insert into books(title,author,year) values ('golang is great','mr great','2021')