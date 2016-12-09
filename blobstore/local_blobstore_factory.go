package blobstore

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type LocalBlobstoreFactory interface {
	NewBlobstore() (Blobstore, error)
}

type localBlobstoreFactory struct {
	fs     boshsys.FileSystem
	logger boshlog.Logger
}

func NewLocalBlobstoreFactory(fs boshsys.FileSystem, logger boshlog.Logger) LocalBlobstoreFactory {
	return localBlobstoreFactory{
		fs:     fs,
		logger: logger,
	}
}

func (f localBlobstoreFactory) NewBlobstore() (Blobstore, error) {
	return NewLocalBlobstore(f.fs, f.logger), nil
}
