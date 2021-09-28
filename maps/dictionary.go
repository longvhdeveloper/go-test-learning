package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExist        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update/delete word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(keyword string) (string, error) {
	if value, ok := dictionary[keyword]; ok {
		return value, nil
	}
	return "", ErrNotFound
}

func (dictionary Dictionary) Add(key, value string) error {
	if _, ok := dictionary[key]; ok {
		return ErrWordExist
	}
	dictionary[key] = value
	return nil
}

func (dictionary Dictionary) Update(key, newValue string) error {
	if _, ok := dictionary[key]; ok {
		dictionary[key] = newValue
		return nil
	}
	return ErrWordDoesNotExist
}

func (dictionary Dictionary) Delete(key string) error {
	if _, ok := dictionary[key]; ok {
		delete(dictionary, key)
		return nil
	}
	return ErrWordDoesNotExist
}
