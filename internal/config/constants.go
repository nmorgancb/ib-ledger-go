package config

const (
	TransactionStatusPending  = "PENDING"
	TransactionStatusFilled   = "FILLED"
	TransactionStatusCanceled = "CANCELED"
	TransactionStatusFailed   = "FAILED"

	// Used to map Entry Postgres UUIDs
	VenueFeeSenderEntryKey    = "venueFeeSenderEntry"
	VenueFeeReceiverEntryKey  = "venueFeeReceiverEntry"
	RetailFeeSenderEntryKey   = "retailFeeSenderEntry"
	RetailFeeReceiverEntryKey = "retailFeeReceiverEntry"
	SenderEntryKey            = "senderEntry"
	ReceiverEntryKey          = "receiverEntry"

	Debit  = "DEBIT"
	Credit = "CREDIT"
	Buy    = "BUY"
	Sell   = "SELL"

	// QLDB Table Names
	AccountsTable = "Accounts"
	LedgerTable   = "Ledger"
)
