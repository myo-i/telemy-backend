CREATE TABLE accounts (
  id BIGINT PRIMARY KEY,
  nickname VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE outputs (
  id BIGINT PRIMARY KEY,
  account_id BIGINT NOT NULL,
  output_content VARCHAR(255) NOT NULL,
  generated_question VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO accounts (id, nickname, email, password, created_at) VALUES (1, 'test', 'test@test.com', 'password', '2020-02-02');
INSERT INTO accounts (id, nickname, email, password, created_at) VALUES (2, 'test2', 'test2@test.com', 'password2', '2020-02-03');
-- CREATE INDEX accounts_index_0 ON accounts (nickname);

-- CREATE INDEX outputs_index_1 ON outputs (account_id);

-- ALTER TABLE outputs ADD FOREIGN KEY (account_id) REFERENCES accounts (id);