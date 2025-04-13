package controller

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/minio/minio-go/v7"

	"github.com/woxQAQ/upload-server/internal/types"
	"github.com/woxQAQ/upload-server/pkg/constants"
)

func (u *uploadController) Import(
	ctx context.Context,
	mc *minio.Client,
	progressId uint64,
) error {
	pgs, err := u.ps.GetProgress(ctx, progressId)
	if err != nil {
		return err
	}
	bytes, err := pgs.TaskDetail.MarshalJSON()
	if err != nil {
		return err
	}
	var detail types.ImportTaskDetail
	err = json.Unmarshal(bytes, &detail)
	if err != nil {
		return err
	}

	obj, err := mc.GetObject(
		ctx,
		constants.BucketName,
		detail.FileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return err
	}
	defer obj.Close()
	extractor, err := newExtractor(detail.FileType, obj)
	if err != nil {
		return err
	}

	defer extractor.Close()

	var dt_buf = NewDtQueue[[]string](1024, 1024*1024*1024)
	defer dt_buf.Close()
	sinker := newSinker(ctx,
		detail.DatabaseType,
		detail.DSN,
		dt_buf,
	)
	if sinker == nil {
		return errors.New("db sinker not implement")
	}

	done := make(chan struct{})
	errChan := make(chan error, 1)

	u.pool.Submit(func() {
		defer func() {
			done <- struct{}{}
		}()
		err = sinker.Import(ctx, detail.Database, detail.Table)
		if err != nil {
			errChan <- err
		}
	})

	for extractor.Next() {
		select {
		case err := <-errChan:
			return err
		default:
		}
		row, err := extractor.Columns()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		dt_buf.Push(row)
	}

	<-done

	return nil
}
