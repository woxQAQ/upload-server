package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func MustCheckBucket(
	ctx context.Context,
	cli *minio.Client,
	bucketName string,
) {
	ok, err := cli.BucketExists(ctx, bucketName)
	if err != nil {
		panic(err)
	}
	if !ok {
		err = cli.MakeBucket(
			ctx,
			bucketName,
			minio.MakeBucketOptions{},
		)
		if err != nil {
			panic(err)
		}
	}
}
