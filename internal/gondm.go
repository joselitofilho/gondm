package gondm

// Open initialize db session based on dialector.
func Open(dialector Dialector, config *Config) (*DB, error) {
	if config == nil {
		config = &Config{}
	}

	if dialector != nil {
		config.Dialector = dialector
	}

	db := &DB{Config: config}

	if config.Dialector != nil {
		if err := config.Dialector.Init(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}
