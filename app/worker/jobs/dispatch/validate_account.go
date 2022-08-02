package dispatch

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/Crossbell-Box/OperatorSync/app/worker/jobs/callback"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/medium"
	"github.com/Crossbell-Box/OperatorSync/app/worker/platforms/tiktok"
	"github.com/Crossbell-Box/OperatorSync/app/worker/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
	"github.com/nats-io/nats.go"
	"strings"
)

func ValidateAccounts(m *nats.Msg) {
	global.Logger.Debug("New validate request received: ", string(m.Data))

	var validateReq commonTypes.ValidateRequest

	if err := json.Unmarshal(m.Data, &validateReq); err != nil {
		global.Logger.Error("Failed to parse validate request.", err.Error())
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_FAILED_TO_PARSE_JSON, "Failed to parse validate request")
		return
	}

	handle, err := utils.GetCrossbellHandleFromID(validateReq.CrossbellCharacterID)
	if err != nil {
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_HTTP_REQUEST_FAILED, err.Error())
	}

	validateString := strings.ToLower(fmt.Sprintf("Crossbell@%s#%s", handle, validateReq.CrossbellCharacterID))

	switch validateReq.Platform {
	case "medium":
		medium.Account(m.Reply, validateReq.Username, validateString)
	case "tiktok":
		tiktok.Account(m.Reply, validateReq.Username, validateString)
	default:
		callback.ValidateHandleFailed(m.Reply, commonConsts.ERROR_CODE_UNSUPPORTED_PLATFORM, "Unsupported platform")
	}

}
