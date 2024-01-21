package tronscan

type GetAccountResponse struct {
	TransactionsOut                    int32  `json:"transactions_out"`
	AcquiredDelegateFrozenForBandWidth int32  `json:"acquiredDelegateFrozenForBandWidth"`
	RewardNum                          int32  `json:"rewardNum"`
	GreyTag                            string `json:"greyTag"`
	OwnerPermission                    struct {
		Keys []struct {
			Address string `json:"address"`
			Weight  int32  `json:"weight"`
		} `json:"keys"`
		Threshold      int32  `json:"threshold"`
		PermissionName string `json:"permission_name"`
	} `json:"ownerPermission"`
	RedTag          string `json:"redTag"`
	PublicTag       string `json:"publicTag"`
	WithPriceTokens []struct {
		TokenPriceInTrx float64 `json:"tokenPriceInTrx"`
		TokenId         string  `json:"tokenId"`
		Balance         string  `json:"balance"`
		TokenName       string  `json:"tokenName"`
		TokenDecimal    int32   `json:"tokenDecimal"`
		TokenAbbr       string  `json:"tokenAbbr"`
		TokenCanShow    int32   `json:"tokenCanShow"`
		TokenType       string  `json:"tokenType"`
		Vip             bool    `json:"vip"`
		TokenLogo       string  `json:"tokenLogo"`
	} `json:"withPriceTokens"`
	DelegateFrozenForEnergy int32 `json:"delegateFrozenForEnergy"`
	Balance                 int64 `json:"balance"`
	FeedbackRisk            bool  `json:"feedbackRisk"`
	VoteTotal               int32 `json:"voteTotal"`
	TotalFrozen             int32 `json:"totalFrozen"`
	Delegated               struct {
	} `json:"delegated"`
	TransactionsIn        int32 `json:"transactions_in"`
	LatestOperationTime   int64 `json:"latest_operation_time"`
	TotalTransactionCount int32 `json:"totalTransactionCount"`
	Representative        struct {
		LastWithDrawTime int32  `json:"lastWithDrawTime"`
		Allowance        int32  `json:"allowance"`
		Enabled          bool   `json:"enabled"`
		Url              string `json:"url"`
	} `json:"representative"`
	FrozenForBandWidth int32         `json:"frozenForBandWidth"`
	Announcement       string        `json:"announcement"`
	Reward             int32         `json:"reward"`
	AddressTagLogo     string        `json:"addressTagLogo"`
	AllowExchange      []interface{} `json:"allowExchange"`
	Address            string        `json:"address"`
	FrozenSupply       []interface{} `json:"frozen_supply"`
	Bandwidth          struct {
		EnergyRemaining   int32   `json:"energyRemaining"`
		TotalEnergyLimit  int64   `json:"totalEnergyLimit"`
		TotalEnergyWeight int64   `json:"totalEnergyWeight"`
		NetUsed           int32   `json:"netUsed"`
		StorageLimit      int32   `json:"storageLimit"`
		StoragePercentage float64 `json:"storagePercentage"`
		NetPercentage     float64 `json:"netPercentage"`
		StorageUsed       int32   `json:"storageUsed"`
		StorageRemaining  int32   `json:"storageRemaining"`
		FreeNetLimit      int32   `json:"freeNetLimit"`
		EnergyUsed        int32   `json:"energyUsed"`
		FreeNetRemaining  int32   `json:"freeNetRemaining"`
		NetLimit          int32   `json:"netLimit"`
		NetRemaining      int32   `json:"netRemaining"`
		EnergyLimit       int32   `json:"energyLimit"`
		FreeNetUsed       int32   `json:"freeNetUsed"`
		TotalNetWeight    int64   `json:"totalNetWeight"`
		FreeNetPercentage float64 `json:"freeNetPercentage"`
		EnergyPercentage  float64 `json:"energyPercentage"`
		TotalNetLimit     int64   `json:"totalNetLimit"`
	} `json:"bandwidth"`
	DateCreated int64         `json:"date_created"`
	AccountType int32         `json:"accountType"`
	Exchanges   []interface{} `json:"exchanges"`
	Frozen      struct {
		Total    int32         `json:"total"`
		Balances []interface{} `json:"balances"`
	} `json:"frozen"`
	AccountResource struct {
		FrozenBalanceForEnergy struct {
		} `json:"frozen_balance_for_energy"`
	} `json:"accountResource"`
	Transactions                    int32  `json:"transactions"`
	BlueTag                         string `json:"blueTag"`
	Witness                         int32  `json:"witness"`
	DelegateFrozenForBandWidth      int32  `json:"delegateFrozenForBandWidth"`
	Name                            string `json:"name"`
	FrozenForEnergy                 int32  `json:"frozenForEnergy"`
	Activated                       bool   `json:"activated"`
	AcquiredDelegateFrozenForEnergy int32  `json:"acquiredDelegateFrozenForEnergy"`
	ActivePermissions               []struct {
		Operations string `json:"operations"`
		Keys       []struct {
			Address string `json:"address"`
			Weight  int32  `json:"weight"`
		} `json:"keys"`
		Threshold      int32  `json:"threshold"`
		Id             int32  `json:"id"`
		Type           string `json:"type"`
		PermissionName string `json:"permission_name"`
	} `json:"activePermissions"`
}

type GetAccountAuthSecurityResponse struct {
	ApproveProjectCount      int32          `json:"approveProjectCount"`
	ApproveTokenCount        int32          `json:"approveTokenCount"`
	ApproveAddressCount      int32          `json:"approveAddressCount"`
	ApproveRiskContractCount int32          `json:"approveRiskContractCount"`
	ApproveRiskAccountCount  int32          `json:"approveRiskAccountCount"`
	ApproveRiskAddressCount  int32          `json:"approveRiskAddressCount"`
	RiskApprove              []*RiskApprove `json:"riskApprove"`
}

type RiskApprove struct {
	Amount          string `json:"amount"`
	Unlimited       bool   `json:"unlimited"`
	ToAddress       string `json:"to_address"`
	ContractAddress string `json:"contract_address"`
	FromAddress     string `json:"from_address"`
	TokenInfo       struct {
		TokenId      string `json:"tokenId"`
		TokenAbbr    string `json:"tokenAbbr"`
		TokenName    string `json:"tokenName"`
		TokenDecimal int32  `json:"tokenDecimal"`
		TokenCanShow int32  `json:"tokenCanShow"`
		TokenType    string `json:"tokenType"`
		TokenLogo    string `json:"tokenLogo"`
		TokenLevel   string `json:"tokenLevel"`
		IssuerAddr   string `json:"issuerAddr"`
		Vip          bool   `json:"vip"`
	} `json:"tokenInfo"`
	Project struct {
		Id string `json:"id"`
	} `json:"project"`
	ProjectId   string `json:"project_id"`
	ProjectSort int32  `json:"project_sort"`
	OperateTime int64  `json:"operate_time"`
}

type ListTrc20TransfersWithStatusRequest struct {
	Start     int32  `json:"start"`                // start index，default is 0
	Limit     int32  `json:"limit"`                // number of transfers per page
	Trc20Id   string `json:"trc20Id"`              // trc20 token address
	Address   string `json:"address"`              // account address
	Direction int32  `json:"direction,omitempty"`  // transfer in or transfer out 0:all 1:transfer out 2:transfer in
	DbVersion int32  `json:"db_version,omitempty"` // return data is contains approval transfer or not. 1: contains, 0: do not contains
	Reverse   bool   `json:"reverse,omitempty"`    // sort by create time, value is true or false
}

type ListTrc20TransfersWithStatusResponse struct {
	TokenInfo struct {
		TokenId      string `json:"tokenId"`
		TokenAbbr    string `json:"tokenAbbr"`
		TokenName    string `json:"tokenName"`
		TokenDecimal int32  `json:"tokenDecimal"`
		TokenCanShow int32  `json:"tokenCanShow"`
		TokenType    string `json:"tokenType"`
		TokenLogo    string `json:"tokenLogo"`
		TokenLevel   string `json:"tokenLevel"`
		IssuerAddr   string `json:"issuerAddr"`
		Vip          bool   `json:"vip"`
	} `json:"tokenInfo"`
	PageSize int32                       `json:"page_size"`
	Code     int32                       `json:"code"`
	Data     []*Trc20TransfersWithStatus `json:"data"`
}

type Trc20TransfersWithStatus struct {
	Amount          string `json:"amount"`
	Status          int32  `json:"status"`
	ApprovalAmount  string `json:"approval_amount"`
	BlockTimestamp  int64  `json:"block_timestamp"`
	Block           int64  `json:"block"`
	From            string `json:"from"`
	To              string `json:"to"`
	Hash            string `json:"hash"`
	ContractAddress string `json:"contract_address"`
	Confirmed       int32  `json:"confirmed"`
	ContractType    string `json:"contract_type"`
	ContractType1   int32  `json:"contractType"`
	Revert          int32  `json:"revert"`
	ContractRet     string `json:"contract_ret"`
	FinalResult     string `json:"final_result"`
	EventType       string `json:"event_type"`
	IssueAddress    string `json:"issue_address"`
	Decimals        int32  `json:"decimals"`
	TokenName       string `json:"token_name"`
	Id              string `json:"id"`
	Direction       int32  `json:"direction"`
}

type ListAccountTokensRequest struct {
	Address   string `json:"address"`
	Start     int32  `json:"start"`               // start index，default is 0
	Limit     int32  `json:"limit"`               // number of transfers per page
	Token     string `json:"token"`               // Specify token ID or token address
	Hidden    int32  `json:"hidden,omitempty"`    // Whether to hide tokens with small balance. 0: hide (default) 1: show
	Show      int32  `json:"show,omitempty"`      // token type. 1: TRC20 2: TRC721 3: ALL (default) 4: TRC1155
	SortType  int32  `json:"sortType,omitempty"`  // Sort field. 1: price 2: amount (default) 3: quantity
	SortBy    int32  `json:"sortBy,omitempty"`    // Sort order. 0: descending order (default) 1: ascending order
	AssetType int32  `json:"assetType,omitempty"` // Asset type. 0: all 1: pass-through only (default) 2: credential only
}

type ListAccountTokensResponse struct {
	Total int32           `json:"total"`
	Data  []*AccountToken `json:"data"`
}

type AccountToken struct {
	Amount           float64 `json:"amount,omitempty"`
	Quantity         float64 `json:"quantity"`
	TokenId          string  `json:"tokenId"`
	TokenPriceInUsd  float64 `json:"tokenPriceInUsd,omitempty"`
	TokenName        string  `json:"tokenName"`
	TokenAbbr        string  `json:"tokenAbbr"`
	TokenCanShow     int32   `json:"tokenCanShow"`
	TokenLogo        string  `json:"tokenLogo"`
	TokenPriceInTrx  float64 `json:"tokenPriceInTrx,omitempty"`
	AmountInUsd      float64 `json:"amountInUsd,omitempty"`
	Balance          string  `json:"balance"`
	TokenDecimal     int32   `json:"tokenDecimal"`
	TokenType        string  `json:"tokenType"`
	Vip              bool    `json:"vip"`
	OwnerAddress     string  `json:"owner_address,omitempty"`
	TransferCount    int32   `json:"transferCount,omitempty"`
	NrOfTokenHolders int32   `json:"nrOfTokenHolders,omitempty"`
	Project          string  `json:"project,omitempty"`
}
