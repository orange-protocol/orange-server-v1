package did

type DataField struct {
	Provider_did string `json:"provider_did"`
	Method       string `json:"method"`
	Name         string `json:"name"`
	UserDid      string `json:"user_did"`
	BindData     string `json:"bind_data"`
	Data         string `json:"data"`
}

type AlgorithmField struct {
	Provider_did string `json:"provider_did"`
	Method       string `json:"method"`
	Name         string `json:"name"`
	Result       string `json:"result"`
}

type OrangeCredential struct {
	Data      *DataField      `json:"data"`
	Algorithm *AlgorithmField `json:"algorithm"`
}

type Resolver interface {
	ValidateSig(did string, idx int, msg string, sig string) (bool, error)
	CreateCredential(did string, idx int, content interface{}, commit bool) (string, string, error)
	RevokeCredential(did string, idx int, cred string) (string, error)
	SignData(data []byte) ([]byte, error)
	EncrypteDataWithDID(data []byte, did string) ([]byte, error)
}
