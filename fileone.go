//strategy.go
package calcmoi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.kdc.capitalone.com/SPICE/M1-Cloud/pkg/config"
	"github.kdc.capitalone.com/SPICE/M1-Cloud/pkg/dao"
	"github.kdc.capitalone.com/SPICE/M1-Cloud/pkg/model"
	"github.kdc.capitalone.com/SPICE/M1-Cloud/pkg/util/genutil"
)

// Calculate strategies for the MOI.
func calcCIMStrgys(ctx *config.SpiceContext, dbx *sqlx.DB, mastPlan model.MasterPlan) ([]model.MOIStrategy, error) {
	cache := map[string]model.CustMgmtStrgy{}
	strgys := make([]model.MOIStrategy, len(mastPlan.Offers))
	for i, offr := range mastPlan.Offers {
		if offr.CustMgmtStrgyStrctID == nil {
			strgy := getStrgy(offr, model.CustMgmtStrgy{})
			strgys[i] = strgy
			continue
		}
		cached, exists := cache[*offr.CustMgmtStrgyStrctID]
		if exists {
			strgy := getStrgy(offr, cached)
			strgys[i] = strgy
			continue
		}
		strgyInfo, err := dao.GetCustMgmtStrgy(ctx, dbx, *offr.CustMgmtStrgyStrctID)
		if err != nil {
			return nil, fmt.Errorf("calcCIMStrgys: error getting strategy for strategy id (%s), %+v",
				*offr.CustMgmtStrgyStrctID, err)
		}
		cache[*offr.CustMgmtStrgyStrctID] = strgyInfo
		strgy := getStrgy(offr, strgyInfo)
		strgys[i] = strgy
	}
	return strgys, nil
}

func getStrgy(offr model.MasterPlanOffer, strgy model.CustMgmtStrgy) model.MOIStrategy {
	strategy := model.MOIStrategy{}
	strategy.CxMgtStratStctID = genutil.StrPtrCpy(offr.CustMgmtStrgyStrctID)
	strategy.HxPrRprStrtID = genutil.IntPtrCpy(strgy.HlxPrcRprcStrgyID)
	strategy.HxPrRprStrtIDPrmNum = genutil.IntPtrCpy(strgy.HlxPrcRprcStrgyParmNum)
	strategy.HxOverLimStrtID = genutil.IntPtrCpy(strgy.HlxOvrLmtStrgyID)
	strategy.HxDelStrtID = genutil.IntPtrCpy(strgy.HlxDlqyStrgyID)
	strategy.BrndProdID = genutil.IntPtrCpy(strgy.BrnddProdID)
	strategy.GradStrtID = genutil.IntPtrCpy(strgy.GrdunStrgyID)
	strategy.StmtCrStrtID = genutil.IntPtrCpy(strgy.StmtCrStrgyID)
	strategy.HxCrossSellStrtID = genutil.IntPtrCpy(strgy.HlxCrsSellStrgyID)
	strategy.HxRetStrtIDOffr = genutil.IntPtrCpy(strgy.HlxRetntnStrgyOffrID)
	strategy.HxRetStrtIDElig = genutil.IntPtrCpy(strgy.HlxRetntnStrgyEligyID)
	strategy.HxRctvCrLnIncProgStrtIDElig = genutil.IntPtrCpy(strgy.HlxRactvCLIPStrgyEligyID)
	strategy.HxRctvCrLnIncProgStrtIDOffr = genutil.IntPtrCpy(strgy.HlxRactvCLIPStrgyOffrID)
	strategy.HxLnMgmtStrtID = genutil.IntPtrCpy(strgy.HlxLnmgmtStrgyID)
	strategy.HxLnMgmtStrtIDPrmNum = genutil.IntPtrCpy(strgy.HlxLnmgmtStrgyParmNum)
	strategy.MaxCxLvlPlAmtHxLnMgmtStrtIDPrmVal = genutil.StrPtrCpy(strgy.HlxLnmgmtStrgyParmValNum)
	strategy.ProdCellID = offr.ProdCellID
	strategy.TestCellID = offr.TestCellID
	strategy.OffrNum = offr.OffrNum
	return strategy
}

