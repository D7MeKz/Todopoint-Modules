package codes

type WebCode uint

const (
	ZeroCode WebCode = 0

	BadAuthorizationHeader WebCode = 40001
	BadImageHeader         WebCode = 40002

	// Common
	GlobalInternalServerError WebCode = 50000
	FileUploadError           WebCode = 50001

	// User
	// NOTE : 200(OK)
	UserSuccess           WebCode = 120000
	UserLoginSuccess      WebCode = 120001
	UserTokenSetupSuccess WebCode = 120002
	ProfileSuccess        WebCode = 120003

	// 201 : Created
	UserCreationSuccess WebCode = 120101

	// 400 : Bad Request
	UserInvalidJson  WebCode = 140001
	UserInvalidUri   WebCode = 140002
	UserAlreadyExist WebCode = 140003

	// 401
	TokenInvalid     WebCode = 140100
	TokenExpired     WebCode = 140101
	TokenCreationErr WebCode = 140102
	UserNotFound     WebCode = 140110

	UserInternalServerError WebCode = 150000
	UserCreationError       WebCode = 150001
	TokenCreationError      WebCode = 150002
	TokenExpiredErr         WebCode = 150003
	ProfileUpdateError      WebCode = 150004

	UserRedisErr    WebCode = 150010
	UserRedisSetErr WebCode = 150011
	// ---------- Task ---------------
	//NOTE : 200
	TaskOneSuccess    WebCode = 220000
	TaskListSuccess   WebCode = 220001
	TaskUpdateSuccess WebCode = 2200002

	SubtaskOneSuccess  WebCode = 220050
	SubtaskListSuccess WebCode = 220051

	// NOTE : 201
	TaskCreationSuccess    WebCode = 220101
	SubtaskCreationSuccess WebCode = 220151

	// NOTE : 400
	TaskInvalidJson  WebCode = 240001
	TaskInvalidUri   WebCode = 240002
	TaskInvalidQuery WebCode = 240003
	TaskDoesNotFound WebCode = 240400

	SubtaskInvalidJson  WebCode = 240051
	SubtaskDoesNotFound WebCode = 240050
	// NOTE : 500
	TaskCreationError WebCode = 250001
	TaskListError     WebCode = 250002
	TaskDecodingErr   WebCode = 250003

	SubtaskCreationErr WebCode = 250051
	SubtaskAdditionErr WebCode = 250054
	SubtaskUpdateErr   WebCode = 250055

	// NOTE : 503
	TaskUserUnavailable WebCode = 250301
)

/*
GetStatus Function
NOTE : 140400 (1 : label, 404: status codes, 00 : Meta)
*/
func GetStatus(c WebCode) int {

	deletedMeta := c / 100
	code := deletedMeta % 1000
	return int(code)
}

// ConvertFrom
func ConvertFrom(a any) WebCode {
	switch v := a.(type) {
	case float64:
		code := WebCode(v)
		return code
	default:
		panic("Unexpected error to convert WebCode")
		return ZeroCode
	}
}
