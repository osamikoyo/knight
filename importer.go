package main

import "maps"

func MakeManifestWithImports(baseMan *Manifest) (*Manifest, error) {
	manifests := make([]Manifest, len(baseMan.Imports))

	for i, path := range baseMan.Imports {
		man, err := ParseFileWithPath(path)
		if err != nil {
			return nil, err
		}

		manifests[i] = *man
	}

	for _, man := range manifests {
		baseMan.Pipeline = append(baseMan.Pipeline, man.Pipeline...)
		maps.Copy(baseMan.Variables, man.Variables)
	}

	return baseMan, nil
}
