CREATE TABLE users (
  id integer primary key,
  name varchar(50),
  msg varchar(150),
  age integer,
  img_url text,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null
);

CREATE TABLE categories (
  id integer primary key,
  name varchar(50),
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null
);

CREATE TABLE posts (
  id integer primary key,
  user_id integer,
  category_id integer,
  title varchar(150),
  content text,
  created_at timestamp with time zone not null,
  updated_at timestamp with time zone not null
);