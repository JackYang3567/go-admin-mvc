drop table if exists posts;
drop table if exists threads;
drop table if exists sessions;
drop table if exists users;


------admin table-----
CREATE TABLE admins (
  id  serial primary key,
  uuid  varchar(64) not null unique,
  username varchar(255) DEFAULT NULL,
  email   varchar(255) not null unique,
  password varchar(255) DEFAULT NULL,
  last_login_time timestamp now()  ,
  last_login_ip varchar(255) DEFAULT '0',
  type BOOLEAN DEFAULT false,
  status BOOLEAN DEFAULT false,
  session_id varchar(255) DEFAULT NULL,
  google_secret varchar(255) DEFAULT NULL ,
  created_at timestamp DEFAULT now() 
);
comment on table admins is '管理员表';
comment on column admins.uid is '我是唯一主键';
comment on column admins.type is '暂时无效';
comment on column admins.google_secret is '谷歌key';


-------------chitchat table---------------------
create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  admin_id    integer references admins(id),
  created_at timestamp not null   
);

create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null       
);

create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null  
);