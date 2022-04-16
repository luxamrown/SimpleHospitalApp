package utils

type Utils interface {
	CheckRowsAffected()
}

func CheckRowsAffected(err error) (bool, error) {
	var hasChange bool
	if err != nil {
		hasChange = false
		return hasChange, err
	}
	hasChange = true
	return hasChange, nil
}
