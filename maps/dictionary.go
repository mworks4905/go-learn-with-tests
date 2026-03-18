package dictionary

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound          = DictionaryErr("word not found")
	ErrNotFoundForUpdate = DictionaryErr("cannot update, word not found")
	ErrNotFoundForDelete = DictionaryErr("cannot delete, word not found")
	ErrAlreadyExists     = DictionaryErr("word already exists")
)	

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if !ok {
		d[word] = definition
		return nil
	}
	return ErrAlreadyExists
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFoundForUpdate
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFoundForDelete
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}