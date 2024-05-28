package netservice

import "encoding/json"

func RespExtractor(data map[string]interface{}, inp any) error {
	// map to byte
	m, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Convert to specific data(unmarshall)
	err = json.Unmarshal(m, &inp)
	if err != nil {
		return err
	}
	return nil
}
