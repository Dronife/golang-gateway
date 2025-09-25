package serverdto

const RETURN_CODE_OK = 200
const RETURN_BAD_REQUEST = 400
const RETURN_CODE_NOT_FOUND = 404

type Response struct {
	ReturnMessage []string
	ReturnCode    int
}
