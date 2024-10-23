package sql

const DropUsersTable = `
drop table if exists users
`

const CreateUsersTable = `
create table if not exists users (
	id integer primary key,
	name text
)`

const InsertUsers = `
insert into users (name) values ('Matt'), ('Laura'), ('Ross'), ('Eva')
`

const SelectUsers = `
select id, name from users limit 4
`
