package storage

type BoltStorage struct {
}

func (bs *BoltStorage) Remember(text string) error {
	return nil
}

func (bs *BoltStorage) Recall(needle string) ([]Hit, error) {
	return nil, nil
}
