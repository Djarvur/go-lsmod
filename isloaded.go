package lsmod

func IsLoaded(name string) (bool, error) {
	mods, err := LsMod()
	if err != nil {
		return false, err
	}

	_, ok := mods[name]

	return ok, nil
}
