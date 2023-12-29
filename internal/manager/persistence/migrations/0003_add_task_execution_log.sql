

-- +goose Up
CREATE TABLE IF NOT EXISTS `task_execution_log` (
  `id` integer,
  `created_at` datetime,
  `task_id` integer,
  `from_status` varchar(16) DEFAULT "",
  `to_status` varchar(16) DEFAULT "",
  `worker_id` integer,
  `description` varchar(255) DEFAULT "",
  PRIMARY KEY (`task_id`, `worker_id`),
  CONSTRAINT `fk_task_execution_log_task` FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_task_execution_log_worker` FOREIGN KEY (`worker_id`) REFERENCES `workers`(`id`) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE `task_execution_log`;