-- GORM Automigration created a separate `job_storage_infos` table (because we
-- used it wrong, to be fair), which is actually only used as an embedded struct
-- in the `jobs` table. This means this table can be dropped.
--
-- +goose Up
DROP TABLE `job_storage_infos`;

-- +goose Down
CREATE TABLE `job_storage_infos` (`shaman_checkout_id` varchar(255) DEFAULT "");
