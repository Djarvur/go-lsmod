package lsmod

// IsLoaded just check if the specified module loaded
func IsLoaded(name string) (bool, error) {
	mods, err := LsMod()
	if err != nil {
		return false, err
	}

	info, ok := mods[name]

	if !ok {
		return false, nil
	}

	return info.Instances > 0 && info.State == StateLive, nil
}
