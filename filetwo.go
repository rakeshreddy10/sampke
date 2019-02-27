package model

// MOI is the full calculated MOI Object
type MOI struct {
	Sol        *MOISol              `json:"solicitation,omitempty"`
	ProdCells  []MOIProdCell        `json:"productionCells,omitempty"`
	TestCells  []MOITestCell        `json:"testCells,omitempty"`
	Offers     []MOIOffer           `json:"offers,omitempty"`
	Products   []MOIProduct         `json:"products,omitempty"`
	Rates      []MOIRate            `json:"rates,omitempty"`
	Strategies []MOIStrategy        `json:"strategies,omitempty"`
	RespChnls  []MOIResponseChannel `json:"responseChannels,omitempty"`
	Creatives  []MOICreative        `json:"creatives,omitempty"`
	Plastics   []MOIPlastic         `json:"plastics,omitempty"`
	ProdMps    []MOIProdMp          `json:"productMappings,omitempty"`
	ALEs       []MOIALE             `json:"accountLevelExclusions,omitempty"`
}

// MOISol contains the Solicitation Level information of the MOI
type MOISol struct {
	SolID          int      `json:"solicitationId"`
	M1SolLvlStatus *string  `json:"m1UploadReuploadStatus"`
	ApprvlRate     *float32 `json:"anticipatedApprovalRate"`
	GrossRspRt     *float32 `json:"anticipatedGrossResponseRate"`
	CLPID          *string  `json:"multiplePlasticStyleCodes"`
	CntryCd        *string  `json:"countryCode"`
	CrdtBureauCsts *string  `json:"creditBureauCosts"`
	DMSolFlag      *string  `json:"directMailSolicitationFlag"`
	EstTotalPC     *int     `json:"estimatedTotalNumberProductionCells"`
	EstTotalTC     *int     `json:"estimatedTotalNumberTestCells"`
	GenSolNotes    *string  `json:"generalSolicitationNotes"`
	OrgEstMailQty  *int     `json:"originalEstimatedMailQuantity"`
	PrimryCurve    *string  `json:"primaryCurve"`
	ReassgnMSN     *string  `json:"reassignMailingSequenceNumber"`
	RemailCnt      *int     `json:"remailCount"`
	SndSolRemail   *string  `json:"sendSolicitationToRemailMart"`
	TimesMailed    *string  `json:"timesMailed"`
}

// MOIProdCell contains Prod Cell level information for the MOI, for each Prod Cell
type MOIProdCell struct {
	ProdCellID     int     `json:"productionCellId"`
	TestCellID     int     `json:"testCellId"`
	M1PCLvlStatus  *string `json:"marketingOneProductionCellLevelStatus"`
	ProdCellLabel  *string `json:"productionCellLabel"`
	ProdCellTypeID *string `json:"productionCellTypeId"`
	PrntMthdCd     *string `json:"printingMethodCode"`
	COBIndicia     *string `json:"capitalOneBankIndicia"`
}

// MOITestCell contains Test Cell level information for the MOI, for each Test Cell
type MOITestCell struct {
	ProdCellID    int        `json:"productionCellId"`
	TestCellID    int        `json:"testCellId"`
	M1TCLvlStatus *string    `json:"marketingOneTestCellLevelStatus"`
	TestCellLabel *string    `json:"testCellLabel"`
	AcctBkSystm   *string    `json:"accountBookingSystem"`
	CmprsdOf      *string    `json:"comprisedOf"`
	PromoInqry    *string    `json:"promotionInquiry"`
	EstTCQty      *int       `json:"estimatedTestCellQuantity"`
	TCLnchDrpDt   *SpiceDate `json:"testCellLaunchDropDate"`
	AppProcSys    *string    `json:"applicationProcessingSystem"`
	ApprvlCd      *string    `json:"approvalCode"`
	OffrExpDays   *int       `json:"offerExpirationDays"`
	MailClass     *string    `json:"mailClass"`
	DcsngExpDays  *int       `json:"decisioningExpirationDays"`
	PopIntntInd   *string    `json:"populationIntentIndicator"`
	LttrCdRqd     *string    `json:"letterCodeRequired"`
	AddrSlctCrtCd *string    `json:"addressSelectCriteriaCode"`
	SeedSlctrCd   *string    `json:"seedSelectorCode"`
	BckEndCASS    *string    `json:"backendCodingAccuracySupportSystem"`
	BckEndNCOA    *string    `json:"backendNationalChangeOfAddress"`
	ApprvlRt      *float32   `json:"approvalRate"`
	GrssRspRt     *float32   `json:"grossResponseRate"`
	AltAppExpInd  *string    `json:"alternateApplicationExperienceIndicator"`
	CrdtLvl       *string    `json:"creditLevel"`
	LttrVndrCd    *string    `json:"letterVendorCode"`
}

// MOIOffer contains Offer level information for the MOI, for each Offer
type MOIOffer struct {
	OfferNum      int     `json:"offerNumber"`
	M1OCLvlStatus *string `json:"marketingOneOfferCellLevelStatus"`
	OfferLabel    *string `json:"offerLabel"`
	OfferType     *string `json:"offerType"`
	OfferDesc     *string `json:"offerDescription"`
	SCCIQNum      *string `json:"securedCardCallInQuestionNumber"`
	OfferTypeID   *string `json:"offerTypeId"`
	MktgCh        *string `json:"marketingChannel"`
	PymtAlloc     *string `json:"paymentAllocation"`
	ApvlRtDistPct *string `json:"approvalRateDistributionPercentage"`
	FraudSegInd   *string `json:"fraudSegmentIndicator"`
	PreActBunID   *string `json:"preActivationBundleId"`
	RewEarnID     *string `json:"rewardsEarnId"`
	RewRedID      *string `json:"rewardsRedemptionId"`
	VisaProdType  *string `json:"visaProductType"`
	VisaProgNum   *string `json:"visaProgramNumber"`
	MrchNtwkID    *string `json:"merchantNetworkId"`
	ChipSelen     *string `json:"chipSelection"`
	FrnLngInd     *string `json:"foreignLanguageIndicator"`
	BankIDNum     *int    `json:"bankIdNumber"`
	RDMProdID     *int    `json:"referenceDataManagementProductId"`
	RDMProdLnID   *int    `json:"referenceDataManagementProductLineId"`
	ICatProdID    *string `json:"iCatalystProductId"`
	Hype          *string `json:"hype"`
}

// MOIProduct contains Product level information for the MOI, for each Product
type MOIProduct struct {
	ProdLabel     *string `json:"productLabel"`
	PullFrom      *string `json:"pullFrom"`
	ProdIDDesc    *string `json:"productIdDescription"`
	ElecSvcReq    *string `json:"electronicServiceRequested"`
	CreaProdDesc  *string `json:"creativeProductDescription"`
	SCCIQNum      *string `json:"mainstreetCoreMarketingIndicator"`
	MsNTCMktgInd  *string `json:"mainstreetNewToCreditMarketingIndicator"`
	UpMktMktgInd  *string `json:"upmarketMarketingIndicator"`
	BusSeg        *string `json:"businessSegment"`
	CrdName       *string `json:"cardName"`
	ICREAppBnnrs  *string `json:"integratedChannelAndResponseEngineApplicationBanners"`
	SmplSchumer   *string `json:"simpleSchumer"`
	MstGenProdNts *string `json:"masterGeneralProductNotes"`
	SecCrdDepAmt  *string `json:"securedCardDepositAmount"`
	MaxCredLim    *string `json:"maximumCreditLimit"`
	MinCredLim    *string `json:"minimumCreditLimit"`
	UpToCredLim   *string `json:"upToCreditLimit"`
	NoMntnCredLim *string `json:"noMentionCreditLimit"`
	NoMntnAPR     *string `json:"noMentionAnnualPercentageRate"`
	AppStgExp     *string `json:"applicationStageExpiration"`
	BTOnApp       *string `json:"balanceTransferOnApplication"`
	MaxBTAmt      *string `json:"maximumBalanceTransferAmount"`
	MinBTAmt      *string `json:"minimumBalanceTransferAmount"`
	BTFee         *string `json:"balanceTransferFee"`
	BTFeePct      *string `json:"balanceTransferFeePercentage"`
	PtlBTFlag     *string `json:"partialBalanceTransferFlag"`
	BTOffrID      *string `json:"balanceTransferOfferId"`
	BTSrcID       *string `json:"balanceTransferSourceId"`
	BTOffrStartDt *string `json:"balanceTransferOfferStartDate"`
	BTOffrEndDt   *string `json:"balanceTransferOfferEndDate"`
	AllowFastCash *string `json:"allowFastCash"`
	BenefitNotes  *string `json:"benefitNotes"`
	CrdCarCd      *string `json:"cardCarrierCode"`
	CashAdvPct    *string `json:"cashAdvancePercentage"`
	PaidInFullPct *string `json:"paidInFullPercentage"`
	MaxBalWrOfDbt *string `json:"maximumBalanceToWriteOffDebt"`
	CrdProdTypeCd *string `json:"cardProductTypeCode"`
	MaxNumPlstc   *string `json:"maximumNumberPlastics"`
	PctACLAppChg  *string `json:"percentageAboveCreditLimitApprovedForCharge"`
	PctACLToColl  *string `json:"percentageAboveCreditLimitToCollections"`
	PctACLOOColl  *string `json:"percentageAboveCreditLimitOutOfCollections"`
	AvgDlyBalMeth *string `json:"averageDailyBalanceMethod"`
	IncIADBCalc   *string `json:"includeInterestInAverageDailyBalanceCalculation"`
	SvcOwner      *string `json:"serviceOwner"`
	LOBLeaf       *string `json:"lineOfBusinessLeaf"`
	ProfitCtr     *string `json:"profitCenter"`
	PrivChkPt     *string `json:"privacyCheckpoint"`
	CashAdvFeeAmt *string `json:"cashAdvanceFeeAmount"`
	CashAdvFeePct *string `json:"cashAdvanceFeePercentage"`
	FeeAmtMembPd  *string `json:"feeAmountPerMembershipPeriod"`
	MembPd        *string `json:"membershipPeriod"`
	DfrPd         *string `json:"deferralPeriod"`
	MinFinChg     *string `json:"minimumFinanceCharge"`
	MaxResIntToWv *string `json:"maximumResidualInterestToWaive"`
	MinPayFeeStct *string `json:"minimumPayFeeStructure"`
	OvrLimFeeStct *string `json:"overLimitFeeStructure"`
	PstDueFeeStct *string `json:"pastDueFeeStructure"`
	RetChkFeeAmt  *string `json:"returnCheckFeeAmount"`
	AutoClrHsFlg  *string `json:"autoClearingHouseFlag"`
	ReloadFlg     *string `json:"reloadFlag"`
	WlcmKitNum    *string `json:"welcomeKitNumber"`
	ChkType       *string `json:"checkType"`
	MultCrdFamCrd *string `json:"multiCardFamilyCard"`
	SrvExcl       *string `json:"serviceExclusion"`
	PartnerID     *string `json:"partnerId"`
	BrandID       *string `json:"brandId"`
}

// MOIRate contains Rate level information for the MOI, for each Rate
type MOIRate struct {
	ISSRtType       *string `json:"introductorySpecialSegmentRateType"`
	ISSPrdRtType    *string `json:"introductorySpecialSegmentPeriodicRateType"`
	ISSGrcPer       *string `json:"introductorySpecialSegmentGracePeriod"`
	ISSInd          *string `json:"introductorySpecialSegmentIndicator"`
	ISSExpDt        *string `json:"introductorySpecialSegmentExpirationDate"`
	ISSLen          *string `json:"introductorySpecialSegmentLength"`
	ISSEffRt        *string `json:"introductorySpecialSegmentEffectiveRate"`
	ISSVRtIndType   *string `json:"introductorySpecialSegmentVariableRateIndexType"`
	ISSVRtIndVal    *string `json:"introductorySpecialSegmentVariableRateIndexValue"`
	ISSVRtIndEffDt  *string `json:"introductorySpecialSegmentVariableRateIndexEffectiveDate"`
	ISSVRtOffset    *string `json:"introductorySpecialSegmentVariableRateOffset"`
	ISSVRtOffsetSn  *string `json:"introductorySpecialSegmentVariableRateOffsetSign"`
	ISSVCap         *string `json:"introductorySpecialSegmentVariableCap"`
	ISSVFloor       *string `json:"introductorySpecialSegmentVariableFloor"`
	ISSPRBRpr       *string `json:"introductorySpecialSegmentPreventRiskBasedReprice"`
	IPRtType        *string `json:"introductoryPurchaseRateType"`
	IPPerRtType     *string `json:"introductoryPurchasePeriodicRateType"`
	IPGrcPer        *string `json:"introductoryPurchaseGracePeriod"`
	IPExpDt         *string `json:"introductoryPurchaseExpirationDate"`
	IPLen           *string `json:"introductoryPurchaseLength"`
	IPEffRt         *string `json:"introductoryPurchaseEffectiveRate"`
	IPVRtIndType    *string `json:"introductoryPurchaseVariableRateIndexType"`
	IPVRtIndVal     *string `json:"introductoryPurchaseVariableRateIndexValue"`
	IPVRtIndEffDt   *string `json:"introductoryPurchaseVariableRateIndexEffectiveDate"`
	IPVRtOfst       *string `json:"introductoryPurchaseVariableRateOffset"`
	IPVRtOfstSn     *string `json:"introductoryPurchaseVariableRateOffsetSign"`
	IPVCap          *string `json:"introductoryPurchaseVariableCap"`
	IPVFloor        *string `json:"introductoryPurchaseVariableFloor"`
	IPPRBRpr        *string `json:"introductoryPurchasePreventRiskBasedReprice"`
	ICRtType        *string `json:"introductoryCashRateType"`
	ICPerRtType     *string `json:"introductoryCashPeriodicRateType"`
	ICExpDt         *string `json:"introductoryCashExpirationDate"`
	ICLen           *string `json:"introductoryCashLength"`
	ICEffRt         *string `json:"introductoryCashEffectiveRate"`
	ICVRtIndType    *string `json:"introductoryCashVariableRateIndexType"`
	ICVRIndVal      *string `json:"introductoryCashVariableRateIndexValue"`
	ICVRtIndEffDt   *string `json:"introductoryCashVariableRateIndexEffectiveDate"`
	ICVRtOfst       *string `json:"introductoryCashVariableRateOffset"`
	ICVRtOfstSn     *string `json:"introductoryCashVariableRateOffsetSign"`
	ICVCap          *string `json:"introductoryCashVariableCap"`
	ICVFloor        *string `json:"introductoryCashVariableFloor"`
	ICPRBRpr        *string `json:"introductoryCashPreventRiskBasedReprice"`
	NISSRtType      *string `json:"nonintroductorySpecialSegmentRateType"`
	NISSPrdRtType   *string `json:"nonintroductorySpecialSegmentPeriodicRateType"`
	NISSGrcPer      *string `json:"nonintroductorySpecialSegmentGracePeriod"`
	NISSInd         *string `json:"nonintroductorySpecialSegmentIndicator"`
	NISSEffRt       *string `json:"nonintroductorySpecialSegmentEffectiveRate"`
	NISSVRtIndType  *string `json:"nonintroductorySpecialSegmentVariableRateIndexType"`
	NISSVRtIndVal   *string `json:"nonintroductorySpecialSegmentVariableRateIndexValue"`
	NISSVRtIndEffDt *string `json:"nonintroductorySpecialSegmentVariableRateIndexEffectiveDate"`
	NISSVRtOfst     *string `json:"nonintroductorySpecialSegmentVariableRateOffset"`
	NISSVRtOfstSn   *string `json:"nonintroductorySpecialSegmentVariableRateOffsetSign"`
	NISSVCap        *string `json:"nonintroductorySpecialSegmentVariableCap"`
	NISSVFloor      *string `json:"nonintroductorySpecialSegmentVariableFloor"`
	NISSPRBRpr      *string `json:"nonintroductorySpecialSegmentPreventRiskBasedReprice"`
	NIPRtType       *string `json:"nonintroductoryPurchaseRateType"`
	NIPPerRtType    *string `json:"nonintroductoryPurchasePeriodicRateType"`
	NIPGrcPer       *string `json:"nonintroductoryPurchaseGracePeriod"`
	NIPEffRt        *string `json:"nonintroductoryPurchaseEffectiveRate"`
	NIPVRtIndType   *string `json:"nonintroductoryPurchaseVariableRateIndexType"`
	NIPVRtIndVal    *string `json:"nonintroductoryPurchaseVariableRateIndexValue"`
	NIPVRtIndEffDt  *string `json:"nonintroductoryPurchaseVariableRateIndexEffectiveDate"`
	NIPVRtOfst      *string `json:"nonintroductoryPurchaseVariableRateOffset"`
	NIPVRtOfstSn    *string `json:"nonintroductoryPurchaseVariableRateOffsetSign"`
	NIPVCap         *string `json:"nonintroductoryPurchaseVariableCap"`
	NIPVFloor       *string `json:"nonintroductoryPurchaseVariableFloor"`
	NIPPRBRpr       *string `json:"nonintroductoryPurchasePreventRiskBasedReprice"`
	NICRtType       *string `json:"nonintroductoryCashRateType"`
	NICPerRtType    *string `json:"nonintroductoryCashPeriodicRateType"`
	NICEffRt        *string `json:"nonintroductoryCashEffectiveRate"`
	NICVRtIndType   *string `json:"nonintroductoryCashVariableRateIndexType"`
	NICVRtIndVal    *string `json:"nonintroductoryCashVariableRateIndexValue"`
	NICVRtIndEffDt  *string `json:"nonintroductoryCashVariableRateIndexEffectiveDate"`
	NICVRtOfst      *string `json:"nonintroductoryCashVariableRateOffset"`
	NICVRtOfstSn    *string `json:"nonintroductoryCashVariableRateOffsetSign"`
	NICVCap         *string `json:"nonintroductoryCashVariableCap"`
	NICVFloor       *string `json:"nonintroductoryCashVariableFloor"`
	NICPRBRpr       *string `json:"nonintroductoryCashPreventRiskBasedReprice"`
}

// MOIStrategy contains Strategy level information for the MOI, for each Strategy
type MOIStrategy struct {
	CxMgtStratStctID                  *string `json:"customerManagementStrategyStructureIdentifier"`
	HxPrRprStrtID                     *int    `json:"helixPriceRepriceStrategyId"`
	HxPrRprStrtIDPrmNum               *int    `json:"helixPriceRepriceStrategyIdParameterNumber"`
	RBPrDefRtHxPrStrtIDPrmVal         *string `json:"riskBasedPricingDefaultRateHelixPriceStrategyIdParameterValue"`
	HxOverLimStrtID                   *int    `json:"helixOverLimitStrategyId"`
	HxDelStrtID                       *int    `json:"helixDeliquencyStrategyId"`
	BrndProdID                        *int    `json:"brandedProductIdentifier"`
	GradStrtID                        *int    `json:"graduationStrategyId"`
	StmtCrStrtID                      *int    `json:"statementCreditStrategyId"`
	HxCrossSellStrtID                 *int    `json:"helixCrossSellStrategyId"`
	HxRetStrtIDElig                   *int    `json:"helixRetentionStrategyIdEligibility"`
	HxRetStrtIDOffr                   *int    `json:"helixRetentionStrategyIdOffer"`
	HxRctvCrLnIncProgStrtIDElig       *int    `json:"helixReactiveCreditLineIncreaseProgramStrategyIdEligibility"`
	HxRctvCrLnIncProgStrtIDOffr       *int    `json:"helixReactiveCreditLineIncreaseProgramStrategyIdOffer"`
	HxLnMgmtStrtID                    *int    `json:"helixLineManagementStrategyId"`
	HxLnMgmtStrtIDPrmNum              *int    `json:"helixLineManagementStrategyIdParameterNumber"`
	MaxCxLvlPlAmtHxLnMgmtStrtIDPrmVal *string `json:"maximumCustomerLevelPlasticsAmountHelixLineManagementStrategyIdParameterValue"`
	ProdCellID                        int     `json:"productionCellId"`
	TestCellID                        int     `json:"testCellId"`
	OffrNum                           int     `json:"offerNumber"`
}

// MOIResponseChannel contains Response Channel level information for the MOI, for each Response Channel
type MOIResponseChannel struct {
	DMRespChName               *string  `json:"directMailResponseChannelName"`
	DMRespChCurve              *string  `json:"directMailResponseChannelCurve"`
	DMRespChPct                *float32 `json:"directMailResponseChannelPercentage"`
	DMOLDsclsrReq              *string  `json:"directMailOnlineDisclosureRequired"`
	GMCRespChName              *string  `json:"getMyCardResponseChannelName"`
	GMCRespChCurve             *string  `json:"getMyCardResponseChannelCurve"`
	GMCRespChPct               *float32 `json:"getMyCardResponseChannelPercentage"`
	GMCDefIntChRespEngAppName  *string  `json:"getMyCardDefaultIntegratedChannelAndResponseEngineApplicationName"`
	GMCIntChRespEngQSetID      *string  `json:"getMyCardIntegratedChannelAndResponseEngineQuestionSetId"`
	GMCURL                     *string  `json:"getMyCardUniformResourceLocator"`
	GMCOLDsclsrReq             *string  `json:"getMyCardOnlineDisclosureRequired"`
	IBRespChName               *string  `json:"inboundResponseChannelName"`
	IBRespChCurve              *string  `json:"inboundResponseChannelCurve"`
	IBRespChPct                *float32 `json:"inboundResponseChannelPercentage"`
	IBScriptAppName            *string  `json:"inboundScriptApplicationName"`
	IBPhNum                    *string  `json:"inboundPhoneNumber"`
	IBOLDsclsrReq              *string  `json:"inboundOnlineDisclosureRequired"`
	ITMRTAELVRespChName        *string  `json:"inboundTelemarketingRealTimeAtExternalVendorResponseChannelName"`
	ITMRTAELVRespChCurve       *string  `json:"inboundTelemarketingRealTimeAtExternalVendorResponseChannelCurve"`
	ITMRTAELVRespChPct         *float32 `json:"inboundTelemarketingRealTimeAtExternalVendorResponseChannelPercentage"`
	ITMRTAELVScriptAppName     *string  `json:"inboundTelemarketingRealTimeAtExternalVendorScriptApplicationName"`
	ITMRTAELVPhNum             *string  `json:"inboundTelemarketingRealTimeAtExternalVendorPhoneNumber"`
	ITMRTAELVOLDsclsrReq       *string  `json:"inboundTelemarketingRealTimeAtExternalVendorOnlineDisclosureRequired"`
	IRTRespChName              *string  `json:"internetRealTimeResponseChannelName"`
	IRTRespChCurve             *string  `json:"internetRealTimeResponseChannelCurve"`
	IRTRespChPct               *float32 `json:"internetRealTimeResponseChannelPercentage"`
	IRTDefIntChRespEngAppName  *string  `json:"internetRealTimeDefaultIntegratedChannelAndResponseEngineApplicationName"`
	IRTIntChRespEngQSetID      *string  `json:"internetRealTimeIntegratedChannelAndResponseEngineQuestionSetId"`
	IRTOLDsclsrReq             *string  `json:"internetRealTimeOnlineDisclosureRequired"`
	ISRTRespChName             *string  `json:"inStoreRealTimeResponseChannelName"`
	ISRTRespChCurve            *string  `json:"inStoreRealTimeResponseChannelCurve"`
	ISRTRespChPct              *float32 `json:"inStoreRealTimeResponseChannelPercentage"`
	ISRTDefIntChRespEngAppName *string  `json:"inStoreRealTimeDefaultIntegratedChannelAndResponseEngineApplicationName"`
	ISRTOLDsclsrReq            *string  `json:"inStoreRealTimeOnlineDisclosureRequired"`
	RTBRespChName              *string  `json:"realTimeBookingResponseChannelName"`
	RTBRespChCurve             *string  `json:"realTimeBookingResponseChannelCurve"`
	RTBRespChPct               *float32 `json:"realTimeBookingResponseChannelPercentage"`
	RTBDefIntChRespEngAppName  *string  `json:"realTimeBookingDefaultIntegratedChannelAndResponseEngineApplicationName"`
	RTBIntChRespEngQSetID      *string  `json:"realTimeBookingIntegratedChannelAndResponseEngineQuestionSetId"`
	RTBOLDsclsrReq             *string  `json:"realTimeBookingOnlineDisclosureRequired"`
}

// MOICreative contains Creative level information for the MOI, for each Creative
type MOICreative struct {
	CrtvLbl *string `db:"creativeLabel"`
}

// MOIALE contains Account Level Exclusion level information for the MOI, for each Account Level Exclusion
type MOIALE struct {
	AcctLvlExc *string `db:"accountLevelExclusions"`
}

// MOIPlastic contains Plastic level information for the MOI, for each Plastic
type MOIPlastic struct {
	PlStyleNum  *string  `db:"plasticStyleNumber"`
	PlBrand     *string  `db:"plasticBrand"`
	PlCrdType   *string  `db:"plasticCardType"`
	PlStyleDesc *string  `db:"plasticStyleDescription"`
	AliasUsed   *string  `db:"aliasUsed"`
	PlStyleLd   *string  `db:"plasticStyleLead"`
	DefPlLd     *string  `db:"defaultPlasticLead"`
	PlUsagePct  *float32 `db:"plasticUsagePercentage"`
	FrontUGCode *string  `db:"frontUltragraphicCode"`
	RearUGCode  *string  `db:"rearUltragraphicCode"`
}

// MOIProdMp contains Product Mapping level information for the MOI, for each Product Mapping
type MOIProdMp struct {
	ExtProdKey    *string    `db:"externalProductKey"`
	InitAppRespCh *string    `db:"initiatingApplicationResponseChannel"`
	AspnEffDt     *SpiceDate `db:"aspenEffectiveDate"`
}

// SegLibRcrd contains the information in a Segment Library Record
type SegLibRcrd struct {
	Key            string  `db:"seg_lib_key"`
	NoMentCl       *string `db:"no_ment_cl"`
	LOBLeafCd      *string `db:"lob_leaf_cd"`
	MultCrdFmlyCrd *string `db:"mult_card_fmly_card"`
	SvcOwnrCd      *string `db:"svc_ownr_cd"`
	CrLvlCd        *string `db:"cr_lvl_cd"`
}

// ChnlLibRcrd contains the information in a Channel Library Record
type ChnlLibRcrd struct {
	Key               string   `db:"unq_chnls_lib_key"`
	MktgChnlCd        *string  `db:"mktg_chnl_cd"`
	IBRspnsChnlCd     *string  `db:"ib_respns_chnl_cd"`
	IBRspnsChnlCrv    *string  `db:"ib_respns_chnl_curve"`
	IBRspnseChnlPct   *float32 `db:"ib_respns_chnl_pct"`
	IBScptAppNm       *string  `db:"ib_scpt_app_nm"`
	IBPhnNum          *string  `db:"ib_phn_num"`
	IERespnsChnlCd    *string  `db:"ie_respns_chnl_cd"`
	IERespnsChnlCrv   *string  `db:"ie_respns_chnl_curve"`
	IERespnsChnlPct   *float32 `db:"ie_respons_chnl_pct"`
	IEScptAppNm       *string  `db:"ie_scpt_app_nm"`
	IEPhnNum          *string  `db:"ie_phn_num"`
	IRRespnsChnlCd    *string  `db:"ir_respns_chnl_cd"`
	IRRespnsChnlCdCrv *string  `db:"ir_respns_chnl_curve"`
	IRRespnsChnlCdPct *float32 `db:"ir_respns_chnl_pct"`
	IRDfltIcoreAppNm  *string  `db:"ir_dflt_icore_app_nm"`
	ISRspnsChnlCd     *string  `db:"is_respns_chnl_cd"`
	ISRspnsChnlCrv    *string  `db:"is_respns_chnl_curve"`
	ISRspnsChnlPct    *float32 `db:"is_respns_chnl_pct"`
	ISDfltIcoreAppNm  *string  `db:"is_dflt_icore_app_nm"`
	RIRspnsChnlCd     *string  `db:"ri_respns_chnl_cd"`
	RIRspnsChnlCrv    *string  `db:"ri_respns_chnl_curve"`
	RIRspnsChnlPct    *float32 `db:"ri_respns_chnl_pct"`
	RIDfltIcoreAppNm  *string  `db:"ri_dflt_icore_app_nm"`
}

// BTStrcRcrd contains the information in a Balance Transfer Structure library record
type BTStrcRcrd struct {
	Key         string  `db:"bt_strc_id"`
	BTOnAppInd  *string `db:"bt_on_app_ind"`
	BTOfrID     *int    `db:"bt_ofr_id"`
	BTSrcID     *int    `db:"bt_src_id"`
	AppStageInd *string `db:"app_stage_ind"`
	PartBTInd   *string `db:"part_bt_ind"`
}

// CardModule contains the information in a Card Module library record.
type CardModule struct {
	Key    int  `db:"card_module_id"`
	BrndCd *int `db:"card_module_brand_code"`
}

// CustMgmtStrgy contains the information in a Customer Management Strategy library lookup.
type CustMgmtStrgy struct {
	Key                      string  `db:"cust_mgmt_strgy_strc_cd"`
	HlxPrcRprcStrgyID        *int    `db:"helix_price_rprc_strgy_id"`
	HlxPrcRprcStrgyParmNum   *int    `db:"helix_price_rprc_strgy_parm_num"`
	HlxOvrLmtStrgyID         *int    `db:"helix_over_lmt_strgy_id"`
	HlxDlqyStrgyID           *int    `db:"helix_delqy_strgy_id"`
	BrnddProdID              *int    `db:"brndd_prod_id"`
	GrdunStrgyID             *int    `db:"grdun_strgy_id"`
	StmtCrStrgyID            *int    `db:"stmt_cr_strgy_id"`
	HlxCrsSellStrgyID        *int    `db:"helix_crs_sell_strgy_id"`
	HlxRetntnStrgyOffrID     *int    `db:"helix_retntn_strgy_id"`
	HlxRetntnStrgyEligyID    *int    `db:"helix_retntn_strgy_eligy_id"`
	HlxRactvCLIPStrgyEligyID *int    `db:"helix_ractv_clip_strgy_eligy_id"`
	HlxRactvCLIPStrgyOffrID  *int    `db:"helix_ractv_clip_strgy_id"`
	HlxLnmgmtStrgyID         *int    `db:"helix_lnmgmt_strgy_id"`
	HlxLnmgmtStrgyParmNum    *int    `db:"helix_lnmgmt_strgy_parm_num"`
	HlxLnmgmtStrgyParmValNum *string `db:"helix_lnmgmt_strgy_parm_val_num"`
}

// PlsticRec contains the information in a Plastic library record
type PlsticRec struct {
	Key		    int	    `db:"card_module_id"`
	PlStyleDes  *string `db:"plastic_style_code_desc"`
	PlCardtype  *int    `db:"plastic_card_type"`
}
