-- This is the initial Goose migration for Flamenco Worker. It recreates the
-- schema that might exist already, hence the "IF NOT EXISTS" clauses.
--
-- WARNING: the 'Down' step will erase the entire database.
--
-- +goose Up
CREATE TABLE IF NOT EXISTS `task_updates` (
  `id` integer,
  `created_at` datetime,
  `task_id` varchar(36) DEFAULT "",
  `payload` BLOB,
  PRIMARY KEY (`id`)
);

-- +goose Down
DROP TABLE `task_updates`;
