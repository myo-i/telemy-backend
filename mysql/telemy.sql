CREATE TABLE `accounts` (
  `id` bigserial PRIMARY KEY,
  `nickname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT "now()"
);

CREATE TABLE `outputs` (
  `id` bigserial PRIMARY KEY,
  `account_id` bigint NOT NULL,
  `output_content` varchar(255) NOT NULL,
  `generated_question` varchar(255) NOT NULL,
  `created_at` timestamptz NOT NULL DEFAULT "now()"
);

CREATE INDEX `accounts_index_0` ON `accounts` (`nickname`);

CREATE INDEX `outputs_index_1` ON `outputs` (`account_id`);

ALTER TABLE `outputs` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);
