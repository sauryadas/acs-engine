package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Azure/acs-engine/pkg/api/v20160330"
	"github.com/Azure/acs-engine/pkg/api/vlabs"
)

// LoadContainerServiceFromFile loads an ACS Cluster API Model from a JSON file
func LoadContainerServiceFromFile(jsonFile string) (*ContainerService, string, error) {
	contents, e := ioutil.ReadFile(jsonFile)
	if e != nil {
		return nil, "", fmt.Errorf("error reading file %s: %s", jsonFile, e.Error())
	}

	m := &TypeMeta{}
	if err := json.Unmarshal(contents, &m); err != nil {
		return nil, "", err
	}
	version := m.APIVersion

	service, err := LoadContainerService(contents, version)

	return service, version, err
}

// LoadContainerService loads an ACS Cluster API Model, validates it, and returns the unversioned representation
func LoadContainerService(contents []byte, version string) (*ContainerService, error) {
	switch version {
	case v20160330.APIVersion:
		containerService := &v20160330.ContainerService{}
		if e := json.Unmarshal(contents, &containerService); e != nil {
			return nil, e
		}

		if e := containerService.Properties.Validate(); e != nil {
			return nil, e
		}
		return ConvertV20160330ContainerService(containerService), nil

	case vlabs.APIVersion:
		containerService := &vlabs.ContainerService{}
		if e := json.Unmarshal(contents, &containerService); e != nil {
			return nil, e
		}

		if e := containerService.Properties.Validate(); e != nil {
			return nil, e
		}
		return ConvertVLabsContainerService(containerService), nil

	default:
		return nil, fmt.Errorf("unrecognized APIVersion '%s'", version)
	}
}
