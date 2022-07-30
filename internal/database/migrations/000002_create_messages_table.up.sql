CREATE TABLE IF NOT EXISTS messages(
  id serial PRIMARY KEY,
  message TEXT,

  user_id serial,
  CONSTRAINT fk_user
   FOREIGN KEY(user_id)
   REFERENCES users(id)
   ON DELETE NO ACTION
   ON UPDATE NO ACTION
)
