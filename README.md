# go-connect-dbms
connect dbms and execute sql.

my first golang program.

---

# Notes

- localhostにPostgreSQL/MySQLが動いている前提
- PostreSQL
  - user:postgres,pass:postgres,port:5432
  - database:public
  - table名:table1
  - create table table1 ( id serial not null, name varchar(20), primary key(id));
- MySQL
  - user:root,pass:password,port:3306
  - database:database1
  - table:table1
  - create table table1 ( id int auto_increment not null primary key, name varchar(20));

ex)

| id | name |
| -- | ---- |
|  1 | tom  |
|  2 | bob  |

