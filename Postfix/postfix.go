package Postfix

type Resp struct{
	PostfixID        string `json:"postfix_id"`
	ServerID         string `json:"server_id"`
	PostfixSignature string `json:"postfix_signature"`
	PostfixUpdated   string `json:"postfix_updated"`
	PostfixPosition  string `json:"postfix_position"`
	PostfixSize      string `json:"postfix_size"`
	PostfixStatus    string `json:"postfix_status"`
	PostfixPath      string `json:"postfix_path"`
}
