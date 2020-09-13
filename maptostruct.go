package maptostruct

import jsoniter "github.com/json-iterator/go"

func MapToStruct(mapInput interface{}, structOutput interface{}) error {

	data, err := jsoniter.Marshal(mapInput)
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(data, &structOutput)
}
