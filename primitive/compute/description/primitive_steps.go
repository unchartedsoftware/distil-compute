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

package description

import "github.com/uncharted-distil/distil-compute/pipeline"

// NewSimonStep creates a SIMON data classification step.  It examines an input
// dataframe, and assigns types to the columns based on the exposed metadata.
func NewSimonStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	// since Simon has fit & produce, need to set the params from
	// set_training_data. In this case, outputs is not used.
	if inputs["inputs"] != nil && inputs["outputs"] == nil {
		inputs["outputs"] = inputs["inputs"]
	}

	return NewStepData(
		&pipeline.Primitive{
			Id:         "d2fa8df2-6517-3c26-bafc-87b701c4043a",
			Version:    "1.2.1",
			Name:       "simon",
			PythonPath: "d3m.primitives.data_cleaning.column_type_profiler.Simon",
			Digest:     "6fa0e87f8044df78a99507a60648d3b7336161862aee8994ff2ed57bea5b0f05",
		},
		outputMethods,
		map[string]interface{}{"statistical_classification": true},
		inputs,
	)
}

// NewSlothStep creates a Sloth timeseries clustering step.
func NewSlothStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	// since Sloth has fit & produce, need to set the params from
	// set_training_data. In this case, outputs is not used.
	if inputs["inputs"] != nil && inputs["outputs"] == nil {
		inputs["outputs"] = inputs["inputs"]
	}

	return NewStepData(
		&pipeline.Primitive{
			Id:         "77bf4b92-2faa-3e38-bb7e-804131243a7f",
			Version:    "2.0.3",
			Name:       "Sloth",
			PythonPath: "d3m.primitives.clustering.k_means.Sloth",
			Digest:     "cd237a13cbcc980d3c80ebdbc148c6a29dbb588c2897792e4b12aa65a39ce1f6",
		},
		outputMethods,
		map[string]interface{}{"nclusters": 4},
		inputs,
	)
}

// NewUnicornStep creates a unicorn image clustering step.
func NewUnicornStep(inputs map[string]DataRef, outputMethods []string, targetColumns []string, outputLabels []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "475c26dc-eb2e-43d3-acdb-159b80d9f099",
			Version:    "1.1.0",
			Name:       "unicorn",
			PythonPath: "d3m.primitives.digital_image_processing.unicorn.Unicorn",
			Digest:     "8c1280cb1355115d98de08e9981ea7cb95f6952885d5b190d9db789921664020",
		},
		outputMethods,
		map[string]interface{}{
			"target_columns": targetColumns,
			"output_labels":  outputLabels,
		},
		inputs,
	)
}

// NewPCAFeaturesStep creates a PCA-based feature ranking call that can be added to
// a pipeline.
func NewPCAFeaturesStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "04573880-d64f-4791-8932-52b7c3877639",
			Version:    "3.0.2",
			Name:       "PCA Features",
			PythonPath: "d3m.primitives.feature_selection.pca_features.Pcafeatures",
			Digest:     "48b7c56260320b32b800f628fbdb103741aa923750a4926a3ae9067f117ae119",
		},
		outputMethods,
		map[string]interface{}{},
		inputs,
	)
}

// NewTargetRankingStep creates a target ranking call that can be added to
// a pipeline. Ranking is based on mutual information between features and a selected
// target.  Returns a DataFrame containing (col_idx, col_name, score) tuples for
// each ranked feature. Features that could not be ranked are excluded
// from the returned set.
func NewTargetRankingStep(inputs map[string]DataRef, outputMethods []string, targetCol int) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "a31b0c26-cca8-4d54-95b9-886e23df8886",
			Version:    "0.2.0",
			Name:       "Mutual Information Feature Ranking",
			PythonPath: "d3m.primitives.feature_selection.mi_ranking.DistilMIRanking",
			Digest:     "5302eebf2fb8a80e9f00e7b74888aba9eb448a9c0463d9d26786dab717a62c61",
		},
		outputMethods,
		map[string]interface{}{"target_col_index": targetCol},
		inputs,
	)
}

// NewDukeStep creates a wrapper for the Duke dataset classifier.
func NewDukeStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "46612a42-6120-3559-9db9-3aa9a76eb94f",
			Version:    "1.1.7",
			Name:       "duke",
			PythonPath: "d3m.primitives.data_cleaning.text_summarization.Duke",
			Digest:     "b40cbf3631a19ef0141fb852079330c622b00ef286e54a755e6a90fc85be5963",
		},
		outputMethods,
		map[string]interface{}{},
		inputs,
	)
}

// NewDataCleaningStep creates a wrapper for the Punk data cleaning primitive.
func NewDataCleaningStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "fc6bf33a-f3e0-3496-aa47-9a40289661bc",
			Version:    "3.0.1",
			Name:       "Data cleaning",
			PythonPath: "d3m.primitives.data_cleaning.data_cleaning.Datacleaning",
			Digest:     "4b2ad84f9b1d4c906da37cd2c9ae5fc23c054a3e974be17ce6e75147d257cc3d",
		},
		outputMethods,
		map[string]interface{}{},
		inputs,
	)
}

// NewCrocStep creates a wrapper for the Croc image classifier.
func NewCrocStep(inputs map[string]DataRef, outputMethods []string, targetColumns []string, outputLabels []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "404fae2a-2f0a-4c9b-9ad2-fb1528990561",
			Version:    "1.2.4",
			Name:       "croc",
			PythonPath: "d3m.primitives.digital_image_processing.croc.Croc",
			Digest:     "a0cd922401d96b8ffbfe11f5db188b6a4d410119319392932e417b706ed5ae6",
		},
		outputMethods,
		map[string]interface{}{
			"target_columns": targetColumns,
			"output_labels":  outputLabels,
		},
		inputs,
	)
}

// NewDatasetToDataframeStep creates a primitive call that transforms an input dataset
// into a PANDAS dataframe.
func NewDatasetToDataframeStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "4b42ce1e-9b98-4a25-b68e-fad13311eb65",
			Version:    "0.3.0",
			Name:       "Extract a DataFrame from a Dataset",
			PythonPath: "d3m.primitives.data_transformation.dataset_to_dataframe.Common",
			Digest:     "a141e6821de7ae586968b0986237745a5510850e6940cf946db9d50d3828b030",
		},
		outputMethods,
		map[string]interface{}{},
		inputs,
	)
}

// NewHorizontalConcatStep creates a primitive call that concats two data frames.
func NewHorizontalConcatStep(inputs map[string]DataRef, outputMethods []string, useIndex bool, removeSecondIndex bool) *StepData {

	return NewStepData(
		&pipeline.Primitive{
			Id:         "aff6a77a-faa0-41c5-9595-de2e7f7c4760",
			Version:    "0.2.0",
			Name:       "Concatenate two dataframes",
			PythonPath: "d3m.primitives.data_transformation.horizontal_concat.DataFrameConcat",
			Digest:     "dea9e41d4dc0204d963ee01617416a6c5165c008ac717c09e59703dcee5d6cfd",
		},
		outputMethods,
		map[string]interface{}{
			"use_index":           useIndex,
			"remove_second_index": removeSecondIndex,
		},
		inputs,
	)
}

// NewDatasetToDataframeStepWithResource creates a primitive call that transforms an input dataset
// into a PANDAS dataframe using the specified resource.
func NewDatasetToDataframeStepWithResource(inputs map[string]DataRef, outputMethods []string, resourceName string) *StepData {
	if resourceName == "" {
		resourceName = "learningData"
	}

	return NewStepData(
		&pipeline.Primitive{
			Id:         "4b42ce1e-9b98-4a25-b68e-fad13311eb65",
			Version:    "0.3.0",
			Name:       "Extract a DataFrame from a Dataset",
			PythonPath: "d3m.primitives.data_transformation.dataset_to_dataframe.Common",
			Digest:     "a141e6821de7ae586968b0986237745a5510850e6940cf946db9d50d3828b030",
		},
		outputMethods,
		map[string]interface{}{
			"dataframe_resource": resourceName,
		},
		inputs,
	)
}

// NewDatasetWrapperStep creates a primitive that wraps a dataframe primitive such that it can be
// used as a datset primitive in the pipeline prepend.  The primitive to wrap is indicated using its
// index in the pipeline.    Leaving the resource ID as the empty value allows the primitive to infer
// the main resource from the dataset.
func NewDatasetWrapperStep(inputs map[string]DataRef, outputMethods []string, primitiveIndex int, resourceID string) *StepData {

	hyperparams := map[string]interface{}{
		"primitive": &PrimitiveReference{primitiveIndex},
	}
	if resourceID != "" {
		hyperparams["resource_id"] = resourceID
	}

	return NewStepData(
		&pipeline.Primitive{
			Id:         "5bef5738-1638-48d6-9935-72445f0eecdc",
			Version:    "0.1.0",
			Name:       "Map DataFrame resources to new resources using provided primitive",
			PythonPath: "d3m.primitives.operator.dataset_map.DataFrameCommon",
			Digest:     "b602026372cab83090708ad7f1c8e8e9d48cd03b1841f59b52b59244727a4aa0",
		},
		outputMethods,
		hyperparams,
		inputs,
	)
}

// ColumnUpdate defines a set of column indices to add/remvoe
// a set of semantic types to/from.
type ColumnUpdate struct {
	Indices       []int
	SemanticTypes []string
}

// NewAddSemanticTypeStep adds semantic data values to an input
// dataset.  An add of (1, 2), ("type a", "type b") would result in "type a" and "type b"
// being added to index 1 and 2.
func NewAddSemanticTypeStep(inputs map[string]DataRef, outputMethods []string, add *ColumnUpdate) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "d7e14b12-abeb-42d8-942f-bdb077b4fd37",
			Version:    "0.1.0",
			Name:       "Add semantic types to columns",
			PythonPath: "d3m.primitives.data_transformation.add_semantic_types.DataFrameCommon",
			Digest:     "f165abd067b013c18459729c20c082efe7f450d98775e4b1579716f4fd988e76",
		},
		outputMethods,
		map[string]interface{}{
			"columns":        add.Indices,
			"semantic_types": add.SemanticTypes,
		},
		inputs,
	)
}

// NewRemoveSemanticTypeStep removes semantic data values from an input
// dataset.  A remove of (1, 2), ("type a", "type b") would result in "type a" and "type b"
// being removed from index 1 and 2.
func NewRemoveSemanticTypeStep(inputs map[string]DataRef, outputMethods []string, remove *ColumnUpdate) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "3002bc5b-fa47-4a3d-882e-a8b5f3d756aa",
			Version:    "0.1.0",
			Name:       "Remove semantic types from columns",
			PythonPath: "d3m.primitives.data_transformation.remove_semantic_types.DataFrameCommon",
			Digest:     "ff48930a123697994f8b606b8a353c7e60aaf21738f4fd1a2611d8d1eb4a349a",
		},
		outputMethods,
		map[string]interface{}{
			"columns":        remove.Indices,
			"semantic_types": remove.SemanticTypes,
		},
		inputs,
	)
}

// NewDenormalizeStep denormalize data that is contained in multiple resource files.
func NewDenormalizeStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "f31f8c1f-d1c5-43e5-a4b2-2ae4a761ef2e",
			Version:    "0.2.0",
			Name:       "Denormalize datasets",
			PythonPath: "d3m.primitives.data_transformation.denormalize.Common",
			Digest:     "6a80776d244347f0d29f4358df1cd0286c25f67e03a7e2ee517c6e853e6a9d1f",
		},
		outputMethods,
		map[string]interface{}{},
		inputs,
	)
}

// NewCSVReaderStep reads data from csv files into a nested dataframe structure.
func NewCSVReaderStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	hyperparams := map[string]interface{}{
		"return_result": "append",
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "989562ac-b50f-4462-99cb-abef80d765b2",
			Version:    "0.1.0",
			Name:       "Columns CSV reader",
			PythonPath: "d3m.primitives.data_preprocessing.csv_reader.DataFrameCommon",
			Digest:     "6f1075710f34deba5954489df0f2a997db5305e805833e0b595b57c5c176c078",
		},
		outputMethods,
		hyperparams,
		inputs,
	)
}

// NewDataFrameFlattenStep searches for nested dataframes and pulls them out.
func NewDataFrameFlattenStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	hyperparams := map[string]interface{}{
		"return_result": "replace",
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "1c4aed23-f3d3-4e6b-9710-009a9bc9b694",
			Version:    "0.1.0",
			Name:       "DataFrame Flatten",
			PythonPath: "d3m.primitives.data_preprocessing.flatten.DataFrameCommon",
			Digest:     "48f7babc531059f17298ea401b9f02e63886179c1190ba15e45310265eac4f60",
		},
		outputMethods,
		hyperparams,
		inputs,
	)
}

// NewColumnParserStep takes obj/string columns in a dataframe and parses them into their
// associated raw python types based on the attached d3m metadata.
func NewColumnParserStep(inputs map[string]DataRef, outputMethods []string, types []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "d510cb7a-1782-4f51-b44c-58f0236e47c7",
			Version:    "0.5.0",
			Name:       "Parses strings into their types",
			PythonPath: "d3m.primitives.data_transformation.column_parser.DataFrameCommon",
			Digest:     "d95eb0ea8a5e6f9abc0965a97e9c4f5d8f74a3df591c11c4145faea3e581cd06",
		},
		outputMethods,
		map[string]interface{}{"parse_semantic_types": types},
		inputs,
	)
}

// NewRemoveColumnsStep removes columns from an input dataframe.  Columns
// are specified by name and the match is case insensitive.
func NewRemoveColumnsStep(inputs map[string]DataRef, outputMethods []string, colIndices []int) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "3b09ba74-cc90-4f22-9e0a-0cf4f29a7e28",
			Version:    "0.1.0",
			Name:       "Removes columns",
			PythonPath: "d3m.primitives.data_transformation.remove_columns.DataFrameCommon",
			Digest:     "d2d01abb8d2183baf0204a9ecb8fefdb43683547a1e26049bf4bf81af1137fa3",
		},
		outputMethods,
		map[string]interface{}{
			"columns": colIndices,
		},
		inputs,
	)
}

// NewRemoveDuplicateColumnsStep removes duplicate columns from a dataframe.
func NewRemoveDuplicateColumnsStep(inputs map[string]DataRef, outputMethods []string) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "130513b9-09ca-4785-b386-37ab31d0cf8b",
			Version:    "0.1.0",
			Name:       "Removes duplicate columns",
			PythonPath: "d3m.primitives.data_transformation.remove_duplicate_columns.DataFrameCommon",
			Digest:     "1504533ed6c09a19c0b3fa3eeae4b8f626e5ceacdcbd247a65996c05f8eb3552",
		},
		outputMethods,
		map[string]interface{}{},
		inputs,
	)
}

// NewTermFilterStep creates a primitive step that filters dataset rows based on a match against a
// term list.  The term match can be partial, or apply to whole terms only.
func NewTermFilterStep(inputs map[string]DataRef, outputMethods []string, colindex int, inclusive bool, terms []string, matchWhole bool) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "a6b27300-4625-41a9-9e91-b4338bfc219b",
			Version:    "0.1.0",
			Name:       "Term list dataset filter",
			PythonPath: "d3m.primitives.data_preprocessing.term_filter.DataFrameCommon",
			Digest:     "48ba9165ceddd92f740bfae8bbcb894986d3dffb430ee3c2269e7952bb2aad0d",
		},
		outputMethods,
		map[string]interface{}{
			"column":      colindex,
			"inclusive":   inclusive,
			"terms":       terms,
			"match_whole": matchWhole,
		},
		inputs,
	)
}

// NewRegexFilterStep creates a primitive step that filter dataset rows based on a regex match.
func NewRegexFilterStep(inputs map[string]DataRef, outputMethods []string, colindex int, inclusive bool, regex string) *StepData {
	hyperparams := map[string]interface{}{
		"column":    colindex,
		"inclusive": inclusive,
		"regex":     regex,
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "cf73bb3d-170b-4ba9-9ead-3dd4b4524b61",
			Version:    "0.1.0",
			Name:       "Regex dataset filter",
			PythonPath: "d3m.primitives.data_preprocessing.regex_filter.DataFrameCommon",
			Digest:     "b6594dce51b2d16d6468cea45619750bc73fcaf9731d52afa1328398b3d54371",
		},
		outputMethods,
		hyperparams,
		inputs,
	)
}

// NewNumericRangeFilterStep creates a primitive step that filters dataset rows based on an
// included/excluded numeric range.  Inclusion of boundaries is controlled by the strict flag.
func NewNumericRangeFilterStep(inputs map[string]DataRef, outputMethods []string, colindex int, inclusive bool, min float64, max float64, strict bool) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "8c246c78-3082-4ec9-844e-5c98fcc76f9d",
			Version:    "0.1.0",
			Name:       "Numeric range filter",
			PythonPath: "d3m.primitives.data_preprocessing.numeric_range_filter.DataFrameCommon",
			Digest:     "031e249edabb35dbd4e6d7453d1e149774678603dfc186d0a1a03c153b132101",
		},
		outputMethods,
		map[string]interface{}{
			"column":    colindex,
			"inclusive": inclusive,
			"min":       min,
			"max":       max,
			"strict":    strict,
		},
		inputs,
	)
}

// NewTimeSeriesLoaderStep creates a primitive step that reads time series values using a dataframe
// containing a file URI column.  The file URIs are expected to point to CSV files, with the
// supplied time and value indices pointing the columns in the CSV that form the series data.
// The result is a new dataframe that stores the timetamps as the column headers,
// and the accompanying values for each file as a row.  Note that the file index column is negative,
// the primitive will use the first CSV file name column if finds.
func NewTimeSeriesLoaderStep(inputs map[string]DataRef, outputMethods []string, fileColIndex int, timeColIndex int, valueColIndex int) *StepData {
	// exclude the file col index val ue in the case of a negative index so that the
	// primitive will infer the colum
	args := map[string]interface{}{
		"time_col_index":  timeColIndex,
		"value_col_index": valueColIndex,
	}
	if fileColIndex >= 0 {
		args["file_col_index"] = fileColIndex
	}

	return NewStepData(
		&pipeline.Primitive{
			Id:         "1689aafa-16dc-4c55-8ad4-76cadcf46086",
			Version:    "0.1.0",
			Name:       "Time series loader",
			PythonPath: "d3m.primitives.distil.TimeSeriesLoader",
			Digest:     "",
		},
		outputMethods,
		args,
		inputs,
	)
}

// NewGoatForwardStep creates a GOAT forward geocoding primitive.  A string column
// containing a place name or address is passed in, and the primitive will
// return a DataFrame containing the lat/lon coords of the place.  If location could
// not be found, the row in the data frame will be empty.
func NewGoatForwardStep(inputs map[string]DataRef, outputMethods []string, placeCol string) *StepData {
	args := map[string]interface{}{
		"target_columns": []string{placeCol},
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "c7c61da3-cf57-354e-8841-664853370106",
			Version:    "1.0.7",
			Name:       "Goat_forward",
			PythonPath: "d3m.primitives.data_cleaning.geocoding.Goat_forward",
			Digest:     "655c3b536ee2b87ec4607ba932650a0655400880de89bba2effee4a7f17df9f8",
		},
		outputMethods,
		args,
		inputs,
	)
}

// NewGoatReverseStep creates a GOAT reverse geocoding primitive.  Columns
// containing lat and lon values are passed in, and the primitive will
// return a DataFrame containing the name of the place, with an
// empty value for coords that no meaningful place could be computed.
func NewGoatReverseStep(inputs map[string]DataRef, outputMethods []string, lonCol string, latCol string) *StepData {
	args := map[string]interface{}{
		"lon_col_index": lonCol,
		"lat_col_index": latCol,
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "f6e4880b-98c7-32f0-b687-a4b1d74c8f99",
			Version:    "1.0.7",
			Name:       "Goat_reverse",
			PythonPath: "d3m.primitives.data_cleaning.geocoding.Goat_reverse",
			Digest:     "2111b6253ac8b3765ccdc1d42b76bf34258b90ef824113d227e1b89a090259b9",
		},
		outputMethods,
		args,
		inputs,
	)
}

// NewJoinStep creates a step that will attempt to join two datasets a key column
// from each.  This is currently a placeholder for testing/debugging only.
func NewJoinStep(inputs map[string]DataRef, outputMethods []string, leftCol string, rightCol string, accuracy float32) *StepData {
	return NewStepData(
		&pipeline.Primitive{
			Id:         "6c3188bf-322d-4f9b-bb91-68151bf1f17f",
			Version:    "0.2.0",
			Name:       "Fuzzy Join Placeholder",
			PythonPath: "d3m.primitives.data_transformation.fuzzy_join.DistilFuzzyJoin",
			Digest:     "",
		},
		outputMethods,
		map[string]interface{}{"left_col": leftCol, "right_col": rightCol, "accuracy": accuracy},
		inputs,
	)
}

// NewDSBoxJoinStep creates a step that will attempt to join two datasets using
// key columns from each dataset.
func NewDSBoxJoinStep(inputs map[string]DataRef, outputMethods []string, leftCols []string, rightCols []string, accuracy float32) *StepData {
	joinType := "exact"
	if accuracy < 0.5 {
		joinType = "approximate"
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "datamart-join",
			Version:    "1.4.4",
			Name:       "Datamart Augmentation",
			PythonPath: "d3m.primitives.data_augmentation.Join.DSBOX",
			Digest:     "",
		},
		outputMethods,
		map[string]interface{}{"left_col": leftCols, "right_col": rightCols, "join_type": joinType},
		inputs,
	)
}

// NewTimeseriesFormatterStep creates a step that will format a time series
// to the long form. The input dataset must be structured using resource
// files for time series data.
func NewTimeseriesFormatterStep(inputs map[string]DataRef, outputMethods []string, mainResID string, fileColIndex int) *StepData {
	args := map[string]interface{}{
		"main_resource_index": mainResID,
	}
	if fileColIndex >= 0 {
		args["file_col_index"] = fileColIndex
	}
	return NewStepData(
		&pipeline.Primitive{
			Id:         "24b09066-836f-4b8f-9773-8c86a5eee26c",
			Version:    "0.2.0",
			Name:       "Time series formatter",
			PythonPath: "d3m.primitives.data_preprocessing.timeseries_formatter.DistilTimeSeriesFormatter",
			Digest:     "",
		},
		outputMethods,
		args,
		inputs,
	)
}
