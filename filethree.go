package model

// MasterPlan contains information about the Master Planning document
type MasterPlan struct {
	Offers []MasterPlanOffer `json:"masterPlanningOffers"`
}

// MasterPlanOffer contains all the information about each Master Planning offer
type MasterPlanOffer struct {
	SolID                   int        `json:"solicitationId" db:"solicitation_id"`
	ProdCellID              int        `json:"productionCellId" db:"prod_cell_id"`
	TestCellID              int        `json:"testCellId" db:"test_cell_id"`
	OffrNum                 int        `json:"offerNumber" db:"ofr_num"`
	M1UpldStat              *string    `json:"m1UploadReuploadStatus" db:"m1_upld_reupld_stat"`
	TCLnchDt                *SpiceDate `json:"testCellLaunchDate" db:"tc_lnch_dt"`
	OffrExpDays             *int       `json:"offerExpirationDays" db:"ofr_expirn_days"`
	DcsngExpDays            *int       `json:"decisioningExpirationDays" db:"dcsng_exprin_days"`
	ProdCellTypeID          *string    `json:"productionCellTypeId" db:"prod_cell_type_id"`
	TestCellLabel           *string    `json:"testCellLabel" db:"tc_lbl"`
	AspenEffDt              *SpiceDate `json:"aspenEffectiveDate" db:"aspen_eff_dt"`
	RollingTsrInd           *string    `json:"rollingTeaserIndicator" db:"rollg_teasr_ind"`
	ICatProdID              *int       `json:"iCatalystProductId" db:"ictlst_prod_id"`
	BalXferStrctID          *string    `json:"balanceTransferStructureId" db:"bt_strc_id"`
	RDMProdLineID           *int       `json:"referenceDataManagementProductLineId" db:"rdm_prod_ln_id"`
	RDMProdID               *int       `json:"referenceDataManagementProductId" db:"rdm_prod_id"`
	FraudSegID              *string    `json:"fraudSegmentId" db:"frd_seg_ind"`
	UnqSegLibKey            *string    `json:"uniqueSegmentLibraryKey" db:"unq_seg_lib_key"`
	CustMgmtStrgyStrctID    *string    `json:"customerManagementStrategyStructureId" db:"cust_mgmt_strgy_strc_id"`
	SCCallInQNum            *string    `json:"securedCardCallInQuestionNumber" db:"sc_call_in_ques_num"`
	UnqChnlLibKey           *string    `json:"uniqueChannelsLibraryKey" db:"unq_chnls_lib_key"`
	PrntAddlDsclsrLtrRqdInd *string    `json:"printAdditionalDisclosureLetterRequiredIndicator" db:"prnt_addl_dsclsr_ltr_reqd_ind"`
	OnlineDiscReqInd        *string    `json:"onlineDiscloserRequiredIndicator" db:"onln_dsclsr_reqd_ind"`
	AcctLvlExclStrctID      *string    `json:"accountLevelExclusionStructureId" db:"acct_lvl_excln_strc_id"`
	ExtProdKey              *string    `json:"externalProductKey" db:"extrnl_prod_key"`
	InitAppRespChnl         *string    `json:"initiatingApplicationResponseChannel" db:"inittg_appn_respns_cnl"`
	AltAppExpInd            *string    `json:"alternateApplicationExperienceIndicator" db:"alt_appln_expr_ind"`
	MktingHypeInd           *string    `json:"marketingHypeId" db:"mktg_hype_ind"`
	ApprvlRtPct             *float32   `json:"approvalRatePercentage" db:"apprl_rt_pct"`
	WelcPkgNum              *int       `json:"welcomePackageNumber" db:"wlcm_pkg_num"`
	BankIDNum               *int       `json:"bankIdNumber" db:"bin"`
	ChipSelCd               *string    `json:"chipSelectionCode" db:"chip_seln_cd"`
	PlasStyleStrctCd        *string    `json:"plasticStyleStructureCode" db:"plstc_style_strc_cd"`
	BenNotes                *string    `json:"benefitNotes" db:"ben_notes"`
	PrevSolAndTCID          *string    `json:"previousSolicitationAndTestCellId" db:"prev_sol_and_tc_id"`
	LangCd                  *string    `json:"languageCode" db:"lang_cd"`
	MoblChnlInd             *string    `json:"mobileChannelIndicator" db:"mbl_chnl_ind"`
	AspenPrequalChnlInd     *string    `json:"aspenPrequalificationChannelIndicator" db:"aspen_prqualn_chnl_ind"`
	AspenNonprequalChnlInd  *string    `json:"aspenNonprequalificationChannelIndicator" db:"aspen_nonprqualn_chnl_ind"`
	EmailChnlInd            *string    `json:"emailChannelIndicator" db:"email_chnl_ind"`
	BannerChnlInd           *string    `json:"bannersChannelIndicator" db:"bnnrs_chnl_ind"`
	SrchEngMktChnlInd       *string    `json:"searchEngineMarketingChannelIndicator" db:"srch_engn_mktg_chnl_ind"`
	UnsolChnlInd            *string    `json:"unsolicitedChannelIndicator" db:"unsolctd_chnl_ind"`
	AffltChnlInd            *string    `json:"affiliatesChannelIndicator" db:"afils_chnl_ind"`
	BankChnlInd             *string    `json:"bankChannelIndicator" db:"bank_chnl_ind"`
	TeleSalesChnlInd        *string    `json:"telesalesChannelIndicator" db:"telesales_chnl_ind"`
	SmlBsBnkChnlInd         *string    `json:"smallBusinessBankChannelIndicator" db:"sb_bank_chnl_ind"`
	SmlBsUnsolChnlInd       *string    `json:"smallBusinessUnsolicitationChannelIndicator" db:"sb_unsol_chnl_ind"`
	AffinChnlInd            *string    `json:"affinitiesChannelIndicator" db:"affinities_chnl_ind"`
	AppPhoneChnlInd         *string    `json:"applicationPhoneChannelIndicator" db:"appn_phn_chnl_ind"`
	GenProdNotesCd          *string    `json:"generalProductNotesCode" db:"genl_prod_notes_cd"`
	CardNm                  *string    `json:"cardName" db:"card_nm"`
	AppBannerCd             *string    `json:"applicationBannerCode" db:"appn_bnnr_cd"`
	SimpSchumerInd          *string    `json:"simp_schumer_ind" db:"simp_schumer_ind"`
}

