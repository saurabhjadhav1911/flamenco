-- This is the initial Goose migration for Flamenco Manager. It recreates the
-- schema that might exist already, hence the "IF NOT EXISTS" clauses.
--
-- WARNING: the 'Down' step will erase the entire database.
--
-- +goose Up
CREATE TABLE IF NOT EXISTS `job_storage_infos` (`shaman_checkout_id` varchar(255) DEFAULT "");
CREATE TABLE IF NOT EXISTS `last_rendereds` (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `job_id` integer DEFAULT 0,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_last_rendereds_job` FOREIGN KEY (`job_id`) REFERENCES `jobs`(`id`) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `task_dependencies` (
  `task_id` integer,
  `dependency_id` integer,
  PRIMARY KEY (`task_id`, `dependency_id`),
  CONSTRAINT `fk_task_dependencies_task` FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_task_dependencies_dependencies` FOREIGN KEY (`dependency_id`) REFERENCES `tasks`(`id`) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `task_failures` (
  `created_at` datetime,
  `task_id` integer,
  `worker_id` integer,
  PRIMARY KEY (`task_id`, `worker_id`),
  CONSTRAINT `fk_task_failures_task` FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_task_failures_worker` FOREIGN KEY (`worker_id`) REFERENCES `workers`(`id`) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `worker_tag_membership` (
  `worker_tag_id` integer,
  `worker_id` integer,
  PRIMARY KEY (`worker_tag_id`, `worker_id`),
  CONSTRAINT `fk_worker_tag_membership_worker_tag` FOREIGN KEY (`worker_tag_id`) REFERENCES `worker_tags`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_worker_tag_membership_worker` FOREIGN KEY (`worker_id`) REFERENCES `workers`(`id`) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "worker_tags" (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `uuid` char(36) UNIQUE DEFAULT "",
  `name` varchar(64) UNIQUE DEFAULT "",
  `description` varchar(255) DEFAULT "",
  PRIMARY KEY (`id`)
);
CREATE TABLE IF NOT EXISTS "jobs" (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `uuid` char(36) UNIQUE DEFAULT "",
  `name` varchar(64) DEFAULT "",
  `job_type` varchar(32) DEFAULT "",
  `priority` smallint DEFAULT 0,
  `status` varchar(32) DEFAULT "",
  `activity` varchar(255) DEFAULT "",
  `settings` jsonb,
  `metadata` jsonb,
  `delete_requested_at` datetime,
  `storage_shaman_checkout_id` varchar(255) DEFAULT "",
  `worker_tag_id` integer,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_jobs_worker_tag` FOREIGN KEY (`worker_tag_id`) REFERENCES `worker_tags`(`id`) ON DELETE
  SET NULL
);
CREATE TABLE IF NOT EXISTS "workers" (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `uuid` char(36) UNIQUE DEFAULT "",
  `secret` varchar(255) DEFAULT "",
  `name` varchar(64) DEFAULT "",
  `address` varchar(39) DEFAULT "",
  `platform` varchar(16) DEFAULT "",
  `software` varchar(32) DEFAULT "",
  `status` varchar(16) DEFAULT "",
  `last_seen_at` datetime,
  `status_requested` varchar(16) DEFAULT "",
  `lazy_status_request` smallint DEFAULT false,
  `supported_task_types` varchar(255) DEFAULT "",
  `can_restart` smallint DEFAULT false,
  PRIMARY KEY (`id`)
);
CREATE TABLE IF NOT EXISTS "job_blocks" (
  `id` integer,
  `created_at` datetime,
  `job_id` integer DEFAULT 0,
  `worker_id` integer DEFAULT 0,
  `task_type` text,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_job_blocks_job` FOREIGN KEY (`job_id`) REFERENCES `jobs`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_job_blocks_worker` FOREIGN KEY (`worker_id`) REFERENCES `workers`(`id`) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "sleep_schedules" (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `worker_id` integer UNIQUE DEFAULT 0,
  `is_active` numeric DEFAULT false,
  `days_of_week` text DEFAULT "",
  `start_time` text DEFAULT "",
  `end_time` text DEFAULT "",
  `next_check` datetime,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_sleep_schedules_worker` FOREIGN KEY (`worker_id`) REFERENCES `workers`(`id`) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "tasks" (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `uuid` char(36) UNIQUE DEFAULT "",
  `name` varchar(64) DEFAULT "",
  `type` varchar(32) DEFAULT "",
  `job_id` integer DEFAULT 0,
  `priority` smallint DEFAULT 50,
  `status` varchar(16) DEFAULT "",
  `worker_id` integer,
  `last_touched_at` datetime,
  `commands` jsonb,
  `activity` varchar(255) DEFAULT "",
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_tasks_job` FOREIGN KEY (`job_id`) REFERENCES `jobs`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_tasks_worker` FOREIGN KEY (`worker_id`) REFERENCES `workers`(`id`) ON DELETE
  SET NULL
);
CREATE INDEX IF NOT EXISTS `idx_worker_tags_uuid` ON `worker_tags`(`uuid`);
CREATE INDEX IF NOT EXISTS `idx_jobs_uuid` ON `jobs`(`uuid`);
CREATE INDEX IF NOT EXISTS `idx_workers_address` ON `workers`(`address`);
CREATE INDEX IF NOT EXISTS `idx_workers_last_seen_at` ON `workers`(`last_seen_at`);
CREATE INDEX IF NOT EXISTS `idx_workers_deleted_at` ON `workers`(`deleted_at`);
CREATE INDEX IF NOT EXISTS `idx_workers_uuid` ON `workers`(`uuid`);
CREATE UNIQUE INDEX IF NOT EXISTS `job_worker_tasktype` ON `job_blocks`(`job_id`, `worker_id`, `task_type`);
CREATE INDEX IF NOT EXISTS `idx_sleep_schedules_is_active` ON `sleep_schedules`(`is_active`);
CREATE INDEX IF NOT EXISTS `idx_sleep_schedules_worker_id` ON `sleep_schedules`(`worker_id`);
CREATE INDEX IF NOT EXISTS `idx_tasks_uuid` ON `tasks`(`uuid`);
CREATE INDEX IF NOT EXISTS `idx_tasks_last_touched_at` ON `tasks`(`last_touched_at`);
-- +goose Down
DROP TABLE `job_storage_infos`;
DROP TABLE `last_rendereds`;
DROP TABLE `task_dependencies`;
DROP TABLE `task_failures`;
DROP TABLE `worker_tag_membership`;
DROP TABLE `worker_tags`;
DROP TABLE `jobs`;
DROP TABLE `workers`;
DROP TABLE `job_blocks`;
DROP TABLE `sleep_schedules`;
DROP TABLE `tasks`;
