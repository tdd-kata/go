package dictionary

type Dictionary map[string]string

const (
	ErrNotFound         = DictError("could not find the word you were looking for")
	ErrWordExists       = DictError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictError("cannot update word because it does not exist")
)

type DictError string

func (e DictError) Error() string {
	return string(e)
}

func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dict Dictionary) Add(word, definition string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		dict[word] = definition
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
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err

	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// built-in function of map
	delete(d, word)
}
