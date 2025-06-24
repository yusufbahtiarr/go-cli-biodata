CREATE TYPE gender_type AS ENUM ('pria', 'wanita');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL CHECK (username <> ''),
  password VARCHAR(255) NOT NULL CHECK (password <> ''));

CREATE TABLE profiles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL CHECK (name <> ''),
  age INT NOT NULL CHECK (age > 0),  
  gender gender_type NOT NULL,
  created_at TIMESTAMP(0) DEFAULT NOW(),
  id_user INTEGER NOT NULL,
  FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO users (username, password) VALUES
  ('Yusuf', 1111),
  ('Bahtiar', 2222);

INSERT INTO profiles (name, age, gender, id_user) VALUES
  ('Putri Patricia', 30, 'wanita', 1),
  ('Joko Susilo', 28, 'pria', 1),
  ('Dewi Lestari', 22, 'wanita', 1),
  ('Agus Prabowo', 35, 'pria',  1);


  drop table users;
  drop table profiles;
