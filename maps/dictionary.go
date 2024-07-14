package maps

const (
	ErrNotFound      = DictionaryErr("couldn't find the word you were looking for")
	ErrWordExists    = DictionaryErr("cannot add word because it already exist")
	ErrWordNotExists = DictionaryErr("cannot add word because it not already exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definiton, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definiton, nil
}

func (d Dictionary) Add(word, definiton string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definiton
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
