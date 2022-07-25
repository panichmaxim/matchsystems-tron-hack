package cfg

import "gitlab.com/falaleev-golang/config"

func Load() (*DistributedConfig, error) {
	sharedConfig := &DistributedConfig{}
	if err := config.LoadDefault(sharedConfig); err != nil {
		return nil, err
	}

	return sharedConfig, nil
}
