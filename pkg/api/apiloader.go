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
	return LoadContainerService(contents)
}

// LoadContainerService loads an ACS Cluster API Model, validates it, and returns the unversioned representation
func LoadContainerService(contents []byte) (*ContainerService, string, error) {
	m := &TypeMeta{}
	if err := json.Unmarshal(contents, &m); err != nil {
		return nil, "", err
	}

	version := m.APIVersion

	switch version {
	case v20160330.APIVersion:
		containerService := &v20160330.ContainerService{}
		if e := json.Unmarshal(contents, &containerService); e != nil {
			return nil, version, e
		}

		if e := containerService.Properties.Validate(); e != nil {
			return nil, version, e
		}
		return ConvertV20160330ContainerService(containerService), version, nil

	case vlabs.APIVersion:
		containerService := &vlabs.ContainerService{}
		if e := json.Unmarshal(contents, &containerService); e != nil {
			return nil, version, e
		}

		if e := containerService.Properties.Validate(); e != nil {
			return nil, version, e
		}
		return ConvertVLabsContainerService(containerService), version, nil

	default:
		return nil, version, fmt.Errorf("unrecognized APIVersion '%s'", m.APIVersion)
	}
}
