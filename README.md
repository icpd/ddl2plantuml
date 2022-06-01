# ddl2plantuml

ddl2plantuml is a tool to generate plantuml ER diagram from database ddl.

### Supported database

- [x] mysql

### Quick start

create a sql file, for example:
```sql
create table user
(
  id         int          not null auto_increment comment 'user`s id',
  name       varchar(255) not null comment 'user`s name',
  created_at datetime     not null default current_timestamp,
  updated_at datetime     not null default current_timestamp on update current_timestamp,
  deleted_at datetime     null,
  primary key (id)
) comment 'user table';

create table credit_card
(
  id         int          not null auto_increment comment 'credit card`s id',
  user_id    int          not null comment 'user`s id',
  number     varchar(255) not null comment 'credit card number',
  created_at datetime     not null default current_timestamp,
  updated_at datetime     not null default current_timestamp on update current_timestamp,
  deleted_at datetime     null,
  primary key (id)
) comment 'credit card table';

create table car
(
  id         int          not null auto_increment comment 'car`s id',
  user_id    int          not null comment 'user`s id',
  name       varchar(255) not null comment 'car`s name',
  created_at datetime     not null default current_timestamp,
  updated_at datetime     not null default current_timestamp on update current_timestamp,
  deleted_at datetime     null,
  primary key (id)
)
```

run the command:
```sh
$ ddl2plantuml -f example.sql
```
screenshot  
![Screenshot.png](Screenshot.png)

### Installation
- download the latest release from [Release](https://github.com/whoisix/ddl2plantuml/releases)
- docker run. replace file directory and file name
    ```sh
    $ docker run -v {{ddlpath}}:/var  whoisix/ddl2plantuml -f /var/{{ddlfile}}  -o /var 
    ```

### Usage

```sh
$ ddl2plantuml -h 
NAME:
   ddl2plantuml - Convert DDL to PlantUML

USAGE:
   ddl2plantuml [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

DESCRIPTION:
   ddl2plantuml is a tool to generate plantuml ER diagram from database ddl.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --driver value, -d value    database driver (default: "mysql")
   --template value, -t value  plantuml template file
   --file value, -f value      ddl sql file, required
   --output value, -o value    output directory (default: ".")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```

### Reference
[wangyuheng/ddl2plantuml](https://github.com/wangyuheng/ddl2plantuml)
