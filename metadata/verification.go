//
//   Copyright © 2019 Uncharted Software Inc.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package metadata

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/araddon/dateparse"
	"github.com/pkg/errors"

	"github.com/uncharted-distil/distil-compute/model"
	log "github.com/unchartedsoftware/plog"
)

// VerifyAndUpdate will update the metadata when inconsistentices or errors
// are found.
func VerifyAndUpdate(m *model.Metadata, dataPath string) (bool, error) {
	log.Infof("verifying metadata")
	updated := false
	// read the data
	csvFile, err := os.Open(dataPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to open data file")
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)

	// skip header
	_, err = reader.Read()
	if err != nil {
		return false, errors.Wrap(err, "failed to read header from data file")
	}

	// cycle through the whole dataset
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return false, errors.Wrap(err, "failed to read line from data file")
		}

		updatedType, err := checkTypes(m, line)
		if err != nil {
			return false, errors.Wrap(err, "unable to check data types")
		}
		if updatedType {
			updated = true
		}
	}
	log.Infof("done checking data types")

	// check role consistency
	for _, v := range m.DataResources[0].Variables {
		// if the type is index, then set the role to index as well
		if v.Type == model.IndexType {
			log.Infof("updating %s role to index to match identified type", v.Name)
			v.Role = []string{model.RoleIndex}
			v.SelectedRole = model.RoleIndex
			updated = true
		}
	}

	log.Infof("done verifying metadata")

	return updated, nil
}

func checkTypes(m *model.Metadata, row []string) (bool, error) {
	// cycle through all variables
	updated := false
	for _, v := range m.DataResources[0].Variables {
		// set the type to text if the data doesn't match the metadata
		if !typeMatchesData(v, row) {
			log.Infof("updating %s type to text since the data did not match", v.Name)
			v.Type = model.StringType
			updated = true
		}
	}

	return updated, nil
}

func typeMatchesData(v *model.Variable, row []string) bool {
	val := row[v.Index]
	good := true

	switch v.Type {
	case model.DateTimeType:
		// a date has to be at least 8 characters (yyyymmdd)
		// the library is a bit too permissive
		if len(val) < 8 {
			good = false
		} else {
			_, err := dateparse.ParseAny(val)
			good = err == nil
			if err != nil {
				log.Warnf("error attempting to parse date value '%s': %v", val, err)
			}
		}
	}

	return good
}