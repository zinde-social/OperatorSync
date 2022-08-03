package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/server/global"
	"github.com/Crossbell-Box/OperatorSync/app/server/models"
	"github.com/Crossbell-Box/OperatorSync/app/server/types"
	"github.com/Crossbell-Box/OperatorSync/app/server/utils"
	commonConsts "github.com/Crossbell-Box/OperatorSync/common/consts"
	commonTypes "github.com/Crossbell-Box/OperatorSync/common/types"
)

func feedOnChainDispatchWork(account *models.Account, feeds []models.Feed) {
	if utils.IsAccountOnChainPaused(account) {
		global.Logger.Errorf("OnChain process has been paused for account %s#%d with reason: %s", account.Platform, account.ID, account.OnChainPauseMessage)
		return
	}

	for _, feed := range feeds {

		var err error
		var onChainRequestBytes []byte

		if onChainRequestBytes, err = json.Marshal(&commonTypes.OnChainRequest{
			FeedID:               feed.ID,
			CrossbellCharacterID: account.CrossbellCharacterID,
			Platform:             account.Platform,
			Username:             account.Username,
			RawFeed:              feed.RawFeed,
		}); err != nil {
			global.Logger.Errorf("Failed to parse OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, err.Error())
			utils.AccountOnChainPause(account, fmt.Sprintf("Failed to parse OnChain work for feed %s#%d", account.Platform, feed.ID))
			break
		}

		onchainResponseMsg, err := global.MQ.Request(commonConsts.MQSETTINGS_OnChainChannelName, onChainRequestBytes, commonConsts.MQSETTINGS_OnChainRequestTimeOut)
		if err != nil {
			global.Logger.Errorf("Failed to dispatch OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, err.Error())
			utils.AccountOnChainPause(account, fmt.Sprintf("Failed to dispatch OnChain work for feed %s#%d", account.Platform, feed.ID))
			break
		}

		// Parse response
		var workRespond commonTypes.OnChainRespond
		if err := json.Unmarshal(onchainResponseMsg.Data, &workRespond); err != nil {
			global.Logger.Errorf("Failed to parse respond: %s", string(onchainResponseMsg.Data))
			utils.AccountOnChainPause(account, fmt.Sprintf("Failed to parse respond: %s", string(onchainResponseMsg.Data)))
			break
		}

		// Validate response
		if !workRespond.IsSucceeded {
			global.Logger.Errorf("Failed to finish OnChain work for feed %s#%d with error: %s", account.Platform, feed.ID, workRespond.Message)
			utils.AccountOnChainPause(account, fmt.Sprintf("Failed to finish OnChain work for feed %s#%d", account.Platform, feed.ID))
			break
		}

		feed.IPFSUri = workRespond.IPFSUri
		feed.Transaction = workRespond.Transaction
	}

	// Update feed from database
	if err := global.DB.Scopes(models.FeedTable(models.Feed{
		Feed: types.Feed{
			Platform: account.Platform,
		},
	})).Save(&feeds).Error; err != nil {
		global.Logger.Error("Failed to save updated feeds", feeds)
	}
}