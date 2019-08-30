package blockchain

import "encoding/json"

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type TransactionMessage struct {
	Trans     Transaction `json:"trans"`
	Timestamp uint64      `json:"time"`
}

func (tr TransactionMessage) Serialize() string {
	r, _ := json.Marshal(tr)
	return string(r)
}

func (tr *TransactionMessage) Deserialize(str string) {
	if err := json.Unmarshal([]byte(str), tr); err != nil {
		panic(err)
	}
}
