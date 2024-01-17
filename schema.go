package tronscan

type GetAccountResponse struct {
	TransactionsOut                    int    `json:"transactions_out"`
	AcquiredDelegateFrozenForBandWidth int    `json:"acquiredDelegateFrozenForBandWidth"`
	RewardNum                          int    `json:"rewardNum"`
	GreyTag                            string `json:"greyTag"`
	OwnerPermission                    struct {
		Keys []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		PermissionName string `json:"permission_name"`
	} `json:"ownerPermission"`
	RedTag          string `json:"redTag"`
	PublicTag       string `json:"publicTag"`
	WithPriceTokens []struct {
		Amount          string `json:"amount"`
		TokenPriceInTrx int    `json:"tokenPriceInTrx"`
		TokenId         string `json:"tokenId"`
		Balance         string `json:"balance"`
		TokenName       string `json:"tokenName"`
		TokenDecimal    int    `json:"tokenDecimal"`
		TokenAbbr       string `json:"tokenAbbr"`
		TokenCanShow    int    `json:"tokenCanShow"`
		TokenType       string `json:"tokenType"`
		Vip             bool   `json:"vip"`
		TokenLogo       string `json:"tokenLogo"`
	} `json:"withPriceTokens"`
	DelegateFrozenForEnergy int   `json:"delegateFrozenForEnergy"`
	Balance                 int64 `json:"balance"`
	FeedbackRisk            bool  `json:"feedbackRisk"`
	VoteTotal               int   `json:"voteTotal"`
	TotalFrozen             int   `json:"totalFrozen"`
	Delegated               struct {
	} `json:"delegated"`
	TransactionsIn        int   `json:"transactions_in"`
	LatestOperationTime   int64 `json:"latest_operation_time"`
	TotalTransactionCount int   `json:"totalTransactionCount"`
	Representative        struct {
		LastWithDrawTime int    `json:"lastWithDrawTime"`
		Allowance        int    `json:"allowance"`
		Enabled          bool   `json:"enabled"`
		Url              string `json:"url"`
	} `json:"representative"`
	FrozenForBandWidth int           `json:"frozenForBandWidth"`
	Announcement       string        `json:"announcement"`
	Reward             int           `json:"reward"`
	AddressTagLogo     string        `json:"addressTagLogo"`
	AllowExchange      []interface{} `json:"allowExchange"`
	Address            string        `json:"address"`
	FrozenSupply       []interface{} `json:"frozen_supply"`
	Bandwidth          struct {
		EnergyRemaining   int   `json:"energyRemaining"`
		TotalEnergyLimit  int64 `json:"totalEnergyLimit"`
		TotalEnergyWeight int64 `json:"totalEnergyWeight"`
		NetUsed           int   `json:"netUsed"`
		StorageLimit      int   `json:"storageLimit"`
		StoragePercentage int   `json:"storagePercentage"`
		NetPercentage     int   `json:"netPercentage"`
		StorageUsed       int   `json:"storageUsed"`
		StorageRemaining  int   `json:"storageRemaining"`
		FreeNetLimit      int   `json:"freeNetLimit"`
		EnergyUsed        int   `json:"energyUsed"`
		FreeNetRemaining  int   `json:"freeNetRemaining"`
		NetLimit          int   `json:"netLimit"`
		NetRemaining      int   `json:"netRemaining"`
		EnergyLimit       int   `json:"energyLimit"`
		FreeNetUsed       int   `json:"freeNetUsed"`
		TotalNetWeight    int64 `json:"totalNetWeight"`
		FreeNetPercentage int   `json:"freeNetPercentage"`
		EnergyPercentage  int   `json:"energyPercentage"`
		TotalNetLimit     int64 `json:"totalNetLimit"`
	} `json:"bandwidth"`
	DateCreated int64         `json:"date_created"`
	AccountType int           `json:"accountType"`
	Exchanges   []interface{} `json:"exchanges"`
	Frozen      struct {
		Total    int           `json:"total"`
		Balances []interface{} `json:"balances"`
	} `json:"frozen"`
	AccountResource struct {
		FrozenBalanceForEnergy struct {
		} `json:"frozen_balance_for_energy"`
	} `json:"accountResource"`
	Transactions                    int    `json:"transactions"`
	BlueTag                         string `json:"blueTag"`
	Witness                         int    `json:"witness"`
	DelegateFrozenForBandWidth      int    `json:"delegateFrozenForBandWidth"`
	Name                            string `json:"name"`
	FrozenForEnergy                 int    `json:"frozenForEnergy"`
	Activated                       bool   `json:"activated"`
	AcquiredDelegateFrozenForEnergy int    `json:"acquiredDelegateFrozenForEnergy"`
	ActivePermissions               []struct {
		Operations string `json:"operations"`
		Keys       []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		Id             int    `json:"id"`
		Type           string `json:"type"`
		PermissionName string `json:"permission_name"`
	} `json:"activePermissions"`
}

type GetAccountAuthSecurityResponse struct {
	ApproveProjectCount      int            `json:"approveProjectCount"`
	ApproveTokenCount        int            `json:"approveTokenCount"`
	ApproveAddressCount      int            `json:"approveAddressCount"`
	ApproveRiskContractCount int            `json:"approveRiskContractCount"`
	ApproveRiskAccountCount  int            `json:"approveRiskAccountCount"`
	ApproveRiskAddressCount  int            `json:"approveRiskAddressCount"`
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
		TokenDecimal int    `json:"tokenDecimal"`
		TokenCanShow int    `json:"tokenCanShow"`
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
	ProjectSort int    `json:"project_sort"`
	OperateTime int64  `json:"operate_time"`
}

type ListTrc20TransfersWithStatusRequest struct {
	Start     int    `json:"start"`                // start index，default is 0
	Limit     int    `json:"limit"`                // number of transfers per page
	Trc20Id   string `json:"trc20Id"`              // trc20 token address
	Address   string `json:"address"`              // account address
	Direction int    `json:"direction,omitempty"`  // transfer in or transfer out 0:all 1:transfer out 2:transfer in
	DbVersion int    `json:"db_version,omitempty"` // return data is contains approval transfer or not. 1: contains, 0: do not contains
	Reverse   bool   `json:"reverse,omitempty"`    // sort by create time, value is true or false
}

type ListTrc20TransfersWithStatusResponse struct {
	TokenInfo struct {
		TokenId      string `json:"tokenId"`
		TokenAbbr    string `json:"tokenAbbr"`
		TokenName    string `json:"tokenName"`
		TokenDecimal int    `json:"tokenDecimal"`
		TokenCanShow int    `json:"tokenCanShow"`
		TokenType    string `json:"tokenType"`
		TokenLogo    string `json:"tokenLogo"`
		TokenLevel   string `json:"tokenLevel"`
		IssuerAddr   string `json:"issuerAddr"`
		Vip          bool   `json:"vip"`
	} `json:"tokenInfo"`
	PageSize int                         `json:"page_size"`
	Code     int                         `json:"code"`
	Data     []*Trc20TransfersWithStatus `json:"data"`
}

type Trc20TransfersWithStatus struct {
	Amount          string `json:"amount"`
	Status          int    `json:"status"`
	ApprovalAmount  string `json:"approval_amount"`
	BlockTimestamp  int64  `json:"block_timestamp"`
	Block           int    `json:"block"`
	From            string `json:"from"`
	To              string `json:"to"`
	Hash            string `json:"hash"`
	ContractAddress string `json:"contract_address"`
	Confirmed       int    `json:"confirmed"`
	ContractType    string `json:"contract_type"`
	ContractType1   int    `json:"contractType"`
	Revert          int    `json:"revert"`
	ContractRet     string `json:"contract_ret"`
	FinalResult     string `json:"final_result"`
	EventType       string `json:"event_type"`
	IssueAddress    string `json:"issue_address"`
	Decimals        int    `json:"decimals"`
	TokenName       string `json:"token_name"`
	Id              string `json:"id"`
	Direction       int    `json:"direction"`
}

type ListAccountTokensRequest struct {
	Address   string `json:"address"`
	Start     int    `json:"start"`               // start index，default is 0
	Limit     int    `json:"limit"`               // number of transfers per page
	Token     string `json:"token"`               // Specify token ID or token address
	Hidden    int    `json:"hidden,omitempty"`    // Whether to hide tokens with small balance. 0: hide (default) 1: show
	Show      int    `json:"show,omitempty"`      // token type. 1: TRC20 2: TRC721 3: ALL (default) 4: TRC1155
	SortType  int    `json:"sortType,omitempty"`  // Sort field. 1: price 2: amount (default) 3: quantity
	SortBy    int    `json:"sortBy,omitempty"`    // Sort order. 0: descending order (default) 1: ascending order
	AssetType int    `json:"assetType,omitempty"` // Asset type. 0: all 1: pass-through only (default) 2: credential only
}

type ListAccountTokensResponse struct {
	Total int             `json:"total"`
	Data  []*AccountToken `json:"data"`
}

type AccountToken struct {
	Amount           float64 `json:"amount,omitempty"`
	Quantity         float64 `json:"quantity"`
	TokenId          string  `json:"tokenId"`
	TokenPriceInUsd  float64 `json:"tokenPriceInUsd,omitempty"`
	TokenName        string  `json:"tokenName"`
	TokenAbbr        string  `json:"tokenAbbr"`
	TokenCanShow     int     `json:"tokenCanShow"`
	TokenLogo        string  `json:"tokenLogo"`
	TokenPriceInTrx  float64 `json:"tokenPriceInTrx,omitempty"`
	AmountInUsd      float64 `json:"amountInUsd,omitempty"`
	Balance          string  `json:"balance"`
	TokenDecimal     int     `json:"tokenDecimal"`
	TokenType        string  `json:"tokenType"`
	Vip              bool    `json:"vip"`
	OwnerAddress     string  `json:"owner_address,omitempty"`
	TransferCount    int     `json:"transferCount,omitempty"`
	NrOfTokenHolders int     `json:"nrOfTokenHolders,omitempty"`
	Project          string  `json:"project,omitempty"`
}
