// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/frodenas/gcs-resource"
)

type FakeGCSClient struct {
	BucketObjectsStub        func(bucketName string, prefix string) ([]string, error)
	bucketObjectsMutex       sync.RWMutex
	bucketObjectsArgsForCall []struct {
		bucketName string
		prefix     string
	}
	bucketObjectsReturns struct {
		result1 []string
		result2 error
	}
	ObjectGenerationsStub        func(bucketName string, objectPath string) ([]int64, error)
	objectGenerationsMutex       sync.RWMutex
	objectGenerationsArgsForCall []struct {
		bucketName string
		objectPath string
	}
	objectGenerationsReturns struct {
		result1 []int64
		result2 error
	}
	DownloadFileStub        func(bucketName string, objectPath string, generation int64, localPath string) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		bucketName string
		objectPath string
		generation int64
		localPath  string
	}
	downloadFileReturns struct {
		result1 error
	}
	UploadFileStub        func(bucketName string, objectPath string, localPath string, predefinedACL string) (int64, error)
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		bucketName    string
		objectPath    string
		localPath     string
		predefinedACL string
	}
	uploadFileReturns struct {
		result1 int64
		result2 error
	}
	URLStub        func(bucketName string, objectPath string, generation int64) (string, error)
	uRLMutex       sync.RWMutex
	uRLArgsForCall []struct {
		bucketName string
		objectPath string
		generation int64
	}
	uRLReturns struct {
		result1 string
		result2 error
	}
	DeleteObjectStub        func(bucketName string, objectPath string, generation int64) error
	deleteObjectMutex       sync.RWMutex
	deleteObjectArgsForCall []struct {
		bucketName string
		objectPath string
		generation int64
	}
	deleteObjectReturns struct {
		result1 error
	}
}

func (fake *FakeGCSClient) BucketObjects(bucketName string, prefix string) ([]string, error) {
	fake.bucketObjectsMutex.Lock()
	fake.bucketObjectsArgsForCall = append(fake.bucketObjectsArgsForCall, struct {
		bucketName string
		prefix     string
	}{bucketName, prefix})
	fake.bucketObjectsMutex.Unlock()
	if fake.BucketObjectsStub != nil {
		return fake.BucketObjectsStub(bucketName, prefix)
	} else {
		return fake.bucketObjectsReturns.result1, fake.bucketObjectsReturns.result2
	}
}

func (fake *FakeGCSClient) BucketObjectsCallCount() int {
	fake.bucketObjectsMutex.RLock()
	defer fake.bucketObjectsMutex.RUnlock()
	return len(fake.bucketObjectsArgsForCall)
}

func (fake *FakeGCSClient) BucketObjectsArgsForCall(i int) (string, string) {
	fake.bucketObjectsMutex.RLock()
	defer fake.bucketObjectsMutex.RUnlock()
	return fake.bucketObjectsArgsForCall[i].bucketName, fake.bucketObjectsArgsForCall[i].prefix
}

func (fake *FakeGCSClient) BucketObjectsReturns(result1 []string, result2 error) {
	fake.BucketObjectsStub = nil
	fake.bucketObjectsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) ObjectGenerations(bucketName string, objectPath string) ([]int64, error) {
	fake.objectGenerationsMutex.Lock()
	fake.objectGenerationsArgsForCall = append(fake.objectGenerationsArgsForCall, struct {
		bucketName string
		objectPath string
	}{bucketName, objectPath})
	fake.objectGenerationsMutex.Unlock()
	if fake.ObjectGenerationsStub != nil {
		return fake.ObjectGenerationsStub(bucketName, objectPath)
	} else {
		return fake.objectGenerationsReturns.result1, fake.objectGenerationsReturns.result2
	}
}

func (fake *FakeGCSClient) ObjectGenerationsCallCount() int {
	fake.objectGenerationsMutex.RLock()
	defer fake.objectGenerationsMutex.RUnlock()
	return len(fake.objectGenerationsArgsForCall)
}

func (fake *FakeGCSClient) ObjectGenerationsArgsForCall(i int) (string, string) {
	fake.objectGenerationsMutex.RLock()
	defer fake.objectGenerationsMutex.RUnlock()
	return fake.objectGenerationsArgsForCall[i].bucketName, fake.objectGenerationsArgsForCall[i].objectPath
}

func (fake *FakeGCSClient) ObjectGenerationsReturns(result1 []int64, result2 error) {
	fake.ObjectGenerationsStub = nil
	fake.objectGenerationsReturns = struct {
		result1 []int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) DownloadFile(bucketName string, objectPath string, generation int64, localPath string) error {
	fake.downloadFileMutex.Lock()
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		bucketName string
		objectPath string
		generation int64
		localPath  string
	}{bucketName, objectPath, generation, localPath})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(bucketName, objectPath, generation, localPath)
	} else {
		return fake.downloadFileReturns.result1
	}
}

func (fake *FakeGCSClient) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakeGCSClient) DownloadFileArgsForCall(i int) (string, string, int64, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.downloadFileArgsForCall[i].bucketName, fake.downloadFileArgsForCall[i].objectPath, fake.downloadFileArgsForCall[i].generation, fake.downloadFileArgsForCall[i].localPath
}

func (fake *FakeGCSClient) DownloadFileReturns(result1 error) {
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGCSClient) UploadFile(bucketName string, objectPath string, localPath string, predefinedACL string) (int64, error) {
	fake.uploadFileMutex.Lock()
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		bucketName    string
		objectPath    string
		localPath     string
		predefinedACL string
	}{bucketName, objectPath, localPath, predefinedACL})
	fake.uploadFileMutex.Unlock()
	if fake.UploadFileStub != nil {
		return fake.UploadFileStub(bucketName, objectPath, localPath, predefinedACL)
	} else {
		return fake.uploadFileReturns.result1, fake.uploadFileReturns.result2
	}
}

func (fake *FakeGCSClient) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *FakeGCSClient) UploadFileArgsForCall(i int) (string, string, string, string) {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return fake.uploadFileArgsForCall[i].bucketName, fake.uploadFileArgsForCall[i].objectPath, fake.uploadFileArgsForCall[i].localPath, fake.uploadFileArgsForCall[i].predefinedACL
}

func (fake *FakeGCSClient) UploadFileReturns(result1 int64, result2 error) {
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) URL(bucketName string, objectPath string, generation int64) (string, error) {
	fake.uRLMutex.Lock()
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct {
		bucketName string
		objectPath string
		generation int64
	}{bucketName, objectPath, generation})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub(bucketName, objectPath, generation)
	} else {
		return fake.uRLReturns.result1, fake.uRLReturns.result2
	}
}

func (fake *FakeGCSClient) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeGCSClient) URLArgsForCall(i int) (string, string, int64) {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return fake.uRLArgsForCall[i].bucketName, fake.uRLArgsForCall[i].objectPath, fake.uRLArgsForCall[i].generation
}

func (fake *FakeGCSClient) URLReturns(result1 string, result2 error) {
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGCSClient) DeleteObject(bucketName string, objectPath string, generation int64) error {
	fake.deleteObjectMutex.Lock()
	fake.deleteObjectArgsForCall = append(fake.deleteObjectArgsForCall, struct {
		bucketName string
		objectPath string
		generation int64
	}{bucketName, objectPath, generation})
	fake.deleteObjectMutex.Unlock()
	if fake.DeleteObjectStub != nil {
		return fake.DeleteObjectStub(bucketName, objectPath, generation)
	} else {
		return fake.deleteObjectReturns.result1
	}
}

func (fake *FakeGCSClient) DeleteObjectCallCount() int {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return len(fake.deleteObjectArgsForCall)
}

func (fake *FakeGCSClient) DeleteObjectArgsForCall(i int) (string, string, int64) {
	fake.deleteObjectMutex.RLock()
	defer fake.deleteObjectMutex.RUnlock()
	return fake.deleteObjectArgsForCall[i].bucketName, fake.deleteObjectArgsForCall[i].objectPath, fake.deleteObjectArgsForCall[i].generation
}

func (fake *FakeGCSClient) DeleteObjectReturns(result1 error) {
	fake.DeleteObjectStub = nil
	fake.deleteObjectReturns = struct {
		result1 error
	}{result1}
}

var _ gcsresource.GCSClient = new(FakeGCSClient)
