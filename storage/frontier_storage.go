package storage

type FrontierStorage struct {
	ToBeDownloaded Storer
	Downloaded     Storer
	InProgress     Storer
}

func NewFrontierStorage() *FrontierStorage {
	return &FrontierStorage{
		ToBeDownloaded: newInMemoryStorage(),
		Downloaded:     newInMemoryStorage(),
		InProgress:     newInMemoryStorage(),
	}
}

func (f *FrontierStorage) PutToBeDownloaded(key string) {
	f.ToBeDownloaded.Put(key, true)
}

func (f *FrontierStorage) PutDownloaded(key string) {
	f.Downloaded.Put(key, true)
}

func (f *FrontierStorage) PutInProgress(key string) {
	f.InProgress.Put(key, true)
}

func (f *FrontierStorage) GetToBeDownloaded(key string) bool {
	return f.ToBeDownloaded.Get(key)
}

func (f *FrontierStorage) GetDownloaded(key string) bool {
	return f.Downloaded.Get(key)
}

func (f *FrontierStorage) GetInProgress(key string) bool {
	return f.InProgress.Get(key)
}

func (f *FrontierStorage) GetAndDelURLToBeDownloaded() string {
	return f.ToBeDownloaded.GetAndDelKey()
}

func (f *FrontierStorage) GetAndDelForDownloaded() string {
	return f.Downloaded.GetAndDelKey()
}

func (f *FrontierStorage) GetAndDelForInProgress() string {
	return f.InProgress.GetAndDelKey()
}

func (f *FrontierStorage) DeleteFromToBeDownloaded(key string) {
	f.ToBeDownloaded.Delete(key)
}

func (f *FrontierStorage) DeleteFromDownloaded(key string) {
	f.Downloaded.Delete(key)
}

func (f *FrontierStorage) DeleteFromInProgress(key string) {
	f.InProgress.Delete(key)
}
