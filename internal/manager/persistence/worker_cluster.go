package persistence

// SPDX-License-Identifier: GPL-3.0-or-later

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type WorkerCluster struct {
	Model

	UUID        string `gorm:"type:char(36);default:'';unique;index"`
	Name        string `gorm:"type:varchar(64);default:'';unique"`
	Description string `gorm:"type:varchar(255);default:''"`

	Workers []*Worker `gorm:"many2many:worker_cluster_membership;constraint:OnDelete:CASCADE"`
}

func (db *DB) CreateWorkerCluster(ctx context.Context, wc *WorkerCluster) error {
	if err := db.gormDB.WithContext(ctx).Create(wc).Error; err != nil {
		return fmt.Errorf("creating new worker cluster: %w", err)
	}
	return nil
}

// HasWorkerClusters returns whether there are any clusters defined at all.
func (db *DB) HasWorkerClusters(ctx context.Context) (bool, error) {
	var count int64
	tx := db.gormDB.WithContext(ctx).
		Model(&WorkerCluster{}).
		Count(&count)
	if err := tx.Error; err != nil {
		return false, workerClusterError(err, "counting worker clusters")
	}
	return count > 0, nil
}

func (db *DB) FetchWorkerCluster(ctx context.Context, uuid string) (*WorkerCluster, error) {
	tx := db.gormDB.WithContext(ctx)
	return fetchWorkerCluster(tx, uuid)
}

// fetchWorkerCluster fetches the worker cluster using the given database instance.
func fetchWorkerCluster(gormDB *gorm.DB, uuid string) (*WorkerCluster, error) {
	w := WorkerCluster{}
	tx := gormDB.First(&w, "uuid = ?", uuid)
	if tx.Error != nil {
		return nil, workerClusterError(tx.Error, "fetching worker cluster")
	}
	return &w, nil
}

func (db *DB) SaveWorkerCluster(ctx context.Context, cluster *WorkerCluster) error {
	if err := db.gormDB.WithContext(ctx).Save(cluster).Error; err != nil {
		return workerClusterError(err, "saving worker cluster")
	}
	return nil
}

// DeleteWorkerCluster deletes the given cluster, after unassigning all workers from it.
func (db *DB) DeleteWorkerCluster(ctx context.Context, uuid string) error {
	tx := db.gormDB.WithContext(ctx).
		Where("uuid = ?", uuid).
		Delete(&WorkerCluster{})
	if tx.Error != nil {
		return workerClusterError(tx.Error, "deleting worker cluster")
	}
	if tx.RowsAffected == 0 {
		return ErrWorkerClusterNotFound
	}
	return nil
}

func (db *DB) FetchWorkerClusters(ctx context.Context) ([]*WorkerCluster, error) {
	clusters := make([]*WorkerCluster, 0)
	tx := db.gormDB.WithContext(ctx).Model(&WorkerCluster{}).Scan(&clusters)
	if tx.Error != nil {
		return nil, workerClusterError(tx.Error, "fetching all worker clusters")
	}
	return clusters, nil
}

func (db *DB) fetchWorkerClustersWithUUID(ctx context.Context, clusterUUIDs []string) ([]*WorkerCluster, error) {
	clusters := make([]*WorkerCluster, 0)
	tx := db.gormDB.WithContext(ctx).
		Model(&WorkerCluster{}).
		Where("uuid in ?", clusterUUIDs).
		Scan(&clusters)
	if tx.Error != nil {
		return nil, workerClusterError(tx.Error, "fetching all worker clusters")
	}
	return clusters, nil
}

func (db *DB) WorkerSetClusters(ctx context.Context, worker *Worker, clusterUUIDs []string) error {
	clusters, err := db.fetchWorkerClustersWithUUID(ctx, clusterUUIDs)
	if err != nil {
		return workerClusterError(err, "fetching worker clusters")
	}

	err = db.gormDB.WithContext(ctx).
		Model(worker).
		Association("Clusters").
		Replace(clusters)
	if err != nil {
		return workerClusterError(err, "updating worker clusters")
	}
	return nil
}
