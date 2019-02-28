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

func TestCalcCIM_MultPlstc(t *testing.T) {
	ctx := new(config.SpiceContext)
	mp := GetTestMP(t)
	mp.Offers[0].PlasStyleStrctCd = genutil.StrPtrAssgn("12345")
	solID := 1234
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	dbx := sqlx.NewDb(mockDB, "sqlmock")
	key := `akey`
	rows1 := sqlmock.NewRows([]string{"seg_lib_key", "no_ment_cl", "lob_leaf_cd", "mult_card_fmly_card", "svc_ownr_cd", "cr_lvl_cd"}).
		AddRow(key, "vala", "valb", "valc", "vald", "vale")
	intVal := 0
	strVal := "string value"
	rows2 := sqlmock.NewRows([]string{"cust_mgmt_strgy_strc_cd", "helix_price_rprc_strgy_id", "helix_price_rprc_strgy_parm_num",
		"helix_over_lmt_strgy_id", "helix_delqy_strgy_id", "brndd_prod_id", "grdun_strgy_id", "stmt_cr_strgy_id",
		"helix_crs_sell_strgy_id", "helix_retntn_strgy_id", "helix_retntn_strgy_eligy_id", "helix_ractv_clip_strgy_eligy_id",
		"helix_ractv_clip_strgy_id", "helix_lnmgmt_strgy_id", "helix_lnmgmt_strgy_parm_num", "helix_lnmgmt_strgy_parm_val_num"}).
		AddRow(key, intVal, intVal, intVal, intVal, intVal, intVal, intVal,
			intVal, intVal, intVal, intVal, intVal, intVal, intVal, strVal)
	mock.ExpectQuery("SELECT").WillReturnRows(rows1)
	mock.ExpectQuery("SELECT").WillReturnRows(rows2)
	// mock.ExpectQuery("SELECT").WillReturnRows(rows2)
	result, err := CalcCIM(ctx, dbx, solID, mp)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	if result.Sol == nil {
		t.Fatalf("unexpected nil solicitation")
	}
	if result.Sol.SolID != solID {
		t.Errorf("expected Sol id %d, got %d", solID, result.Sol.SolID)
	}
	if len(result.TestCells) != 2 {
		t.Errorf("unexpected test cell size: %d", len(result.TestCells))
	}
	if len(result.ProdCells) != 2 {
		t.Errorf("unexpected prod cell size: %d", len(result.ProdCells))
	}
	if result.Sol.CLPID == nil {
		t.Fatalf("unexpected nil value for CLPID")
	}
	if *result.Sol.CLPID != "Y" {
		t.Errorf("unexpected value for CLPID: %s", *result.Sol.CLPID)
	}
	if len(result.Strategies) != 8 {
		t.Errorf("unexpected Strategies size: %d", len(result.Strategies))
	}
	if result.Strategies[0].CxMgtStratStctID == nil {
		t.Fatalf("unexpected nil value for CxMgtStratStctID")
	}
	strgyID := *result.Strategies[0].CxMgtStratStctID
	if strgyID != "cust-mgmt-strgy-strct-ID" {
		t.Errorf("Strategy customer management structure ID:\ngot:  %s\nwant: %s", *result.Strategies[0].CxMgtStratStctID, "cust-mgmt-strgy-strct-ID")
	}
}
