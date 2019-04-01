package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/nu7hatch/gouuid"

	"github.com/concourse/concourse/atc"
)

//go:generate counterfeiter . WorkerArtifact
type ArtifactRepository interface {
	CreateWorkerArtifact() (WorkerArtifact, error)
}

type artifactRepository struct {
	conn Conn
}

func NewArtifactRepository(conn Conn) ArtifactRepository {
	return &artifactRepository{
		conn: conn,
	}
}
type WorkerArtifact interface {
	ID() int
	Name() string
	BuildID() int
	CreatedAt() time.Time
	Volume(teamID int) (CreatedVolume, bool, error)
	AssociateWorkerResourceCache(
		workerName string,
		resourceCache UsedResourceCache,
	) error
}

func (af *artifactRepository) CreateWorkerArtifact() (WorkerArtifact, error) {
	guid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	atcArt := atc.WorkerArtifact{
		Name: guid.String(),
	}

	dbArt, err := saveWorkerArtifactWOTX(af.conn, atcArt)
	if err != nil {
		return nil, err
	}

	return dbArt, nil
}

type artifact struct {
	conn Conn

	id        int
	name      string
	buildID   int
	createdAt time.Time
	workerResourceCacheId int
}

func (a *artifact) ID() int              { return a.id }
func (a *artifact) Name() string         { return a.name }
func (a *artifact) BuildID() int         { return a.buildID }
func (a *artifact) CreatedAt() time.Time { return a.createdAt }

func (a *artifact) Volume(teamID int) (CreatedVolume, bool, error) {
	where := map[string]interface{}{
		"v.team_id":            teamID,
		"v.worker_artifact_id": a.id,
	}

	_, created, err := getVolume(a.conn, where)
	if err != nil {
		return nil, false, err
	}

	if created == nil {
		return nil, false, nil
	}

	return created, true, nil
}



func (a *artifact) AssociateWorkerResourceCache(workerName string, resourceCache UsedResourceCache) error {
		tx, err := a.conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	workerResourceCache, err := WorkerResourceCache{
		WorkerName:    workerName,
		ResourceCache: resourceCache,
	}.FindOrCreate(tx)
	if err != nil {
		return err
	}

	rows, err := psql.Update("as").
		Set("worker_resource_cache_id", workerResourceCache.ID).
		Set("team_id", nil).
		Where(sq.Eq{"id": a.id}).
		RunWith(tx).
		Exec()
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == pqUniqueViolationErrCode {
			// another a was 'blessed' as the cache a - leave this one
			// owned by the container so it just expires when the container is GCed
			return nil
		}

		return err
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrVolumeMissing
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	a.workerResourceCacheId = resourceCache.ID()

	return nil
}

func saveWorkerArtifactWOTX(conn Conn, atcArtifact atc.WorkerArtifact) (WorkerArtifact, error) {
	tx, txerr := conn.Begin()
	if txerr != nil {
		return nil, txerr
	}

	defer tx.Rollback()

	var artifactID int

	values := map[string]interface{}{
		"name": atcArtifact.Name,
	}

	if atcArtifact.BuildID != 0 {
		values["build_id"] = atcArtifact.BuildID
	}

	err := psql.Insert("worker_artifacts").
		SetMap(values).
		Suffix("RETURNING id").
		RunWith(tx).
		QueryRow().
		Scan(&artifactID)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	dbArt, found, err := getWorkerArtifact(tx, conn, atcArtifact.ID)
	if err != nil {
		return nil, err
	}
	if !found  {
		return nil, fmt.Errorf("did not find art")
	}
	return dbArt, nil
}

func saveWorkerArtifact(tx Tx, conn Conn, atcArtifact atc.WorkerArtifact) (WorkerArtifact, error) {

	var artifactID int

	values := map[string]interface{}{
		"name": atcArtifact.Name,
	}

	if atcArtifact.BuildID != 0 {
		values["build_id"] = atcArtifact.BuildID
	}

	err := psql.Insert("worker_artifacts").
		SetMap(values).
		Suffix("RETURNING id").
		RunWith(tx).
		QueryRow().
		Scan(&artifactID)

	if err != nil {
		return nil, err
	}

	artifact, found, err := getWorkerArtifact(tx, conn, artifactID)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("Not found")
	}

	return artifact, nil
}

func getWorkerArtifact(tx Tx, conn Conn, id int) (WorkerArtifact, bool, error) {
	var (
		createdAtTime pq.NullTime
		buildID       sql.NullInt64
	)

	artifact := &artifact{conn: conn}

	err := psql.Select("id", "created_at", "name", "build_id").
		From("worker_artifacts").
		Where(sq.Eq{
			"id": id,
		}).
		RunWith(tx).
		QueryRow().
		Scan(&artifact.id, &createdAtTime, &artifact.name, &buildID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	artifact.createdAt = createdAtTime.Time
	artifact.buildID = int(buildID.Int64)

	return artifact, true, nil
}
