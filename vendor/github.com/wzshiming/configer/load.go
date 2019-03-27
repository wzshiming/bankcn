package configer

func Load(v interface{}, cfgpaths ...string) error {
	for _, path := range cfgpaths {
		d, err := Read(path)
		if err != nil {
			return err
		}
		err = Parse(d, path, v)
		if err != nil {
			return err
		}
	}
	return ProcessTags(v)
}
