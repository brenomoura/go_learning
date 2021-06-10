package main

type Dict map[string]string

const (
	ErrNotFoundWord    = ErrDict("it was not possible to find the word")
	ErrExistentWord    = ErrDict("the word to add already exists")
	ErrNonexistentWord = ErrDict("it was not possible to find the word because it does not exist")
)

type ErrDict string

func (e ErrDict) Error() string {
	return string(e)
}

func (dict Dict) Search(word string) (string, error) {
	def, exists := dict[word]
	if !exists {
		return "", ErrNotFoundWord
	}
	return def, nil
}

func (dict Dict) Add(key, value string) error {
	_, err := dict.Search(key)
	switch err {
	case ErrNotFoundWord:
		dict[key] = value
	case nil:
		return ErrExistentWord
	default:
		return err
	}
	return nil
}

func (dict Dict) Update(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFoundWord:
		return ErrNonexistentWord
	case nil:
		dict[key] = value
	default:
		return err
	}
	return nil
}

func (dict Dict) Delete(key string) {
	delete(dict, key)
}
