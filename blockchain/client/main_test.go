package main

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func TestParseMessage(t *testing.T) {
	exampleXrpTxResult := `{"engine_result":"tesSUCCESS","engine_result_code":0,"engine_result_message":"The transaction was applied. Only final in a validated ledger.","ledger_hash":"0DB3566DC7E57228DB6E2D8A0A697C1B17A8F7251559B22966FD723760075F18","ledger_index":18139963,"meta":{"AffectedNodes":[{"ModifiedNode":{"FinalFields":{"Account":"rsGPNkSLt36BDLMgPAYKifFvCphQJZ2qJw","Balance":"1999999892","Flags":0,"OwnerCount":0,"Sequence":10},"LedgerEntryType":"AccountRoot","LedgerIndex":"131B6989F00BE80F0C5ACA4C642CBD534BB18E2F04E59488F9BAF9ED7017E140","PreviousFields":{"Balance":"2999999904","Sequence":9},"PreviousTxnID":"509FC6682AAE4AFFFB0123EE812AD5A7CEC2F68188B045786841F0DF662750F1","PreviousTxnLgrSeq":18139935}},{"ModifiedNode":{"FinalFields":{"Account":"rs16hESfGChwAnK97oSdRJq4A18gcJbE7j","Balance":"17999996952","Flags":0,"OwnerCount":1,"Sequence":4},"LedgerEntryType":"AccountRoot","LedgerIndex":"FEE0D27B64E9EC97D880E34C172EBEC49147D2D156169BC7A63430EAFD214790","PreviousFields":{"Balance":"16999996952"},"PreviousTxnID":"509FC6682AAE4AFFFB0123EE812AD5A7CEC2F68188B045786841F0DF662750F1","PreviousTxnLgrSeq":18139935}}],"TransactionIndex":0,"TransactionResult":"tesSUCCESS","delivered_amount":"1000000000"},"status":"closed","transaction":{"Account":"rsGPNkSLt36BDLMgPAYKifFvCphQJZ2qJw","Amount":"1000000000","Destination":"rs16hESfGChwAnK97oSdRJq4A18gcJbE7j","Fee":"12","Flags":2147483648,"LastLedgerSequence":18139964,"Memos":[{"Memo":{"MemoData":"7573647861646472727347504E6B534C74333642444C4D675041594B69664676437068514A5A32714A77"}}],"Sequence":9,"SigningPubKey":"02CA3BD0848A81769A35BA0A663098A75A9D7182E8A40418788FDCDFABDA67C660","TransactionType":"Payment","TxnSignature":"3045022100D8525293800A845B561280B8B90AAA08D44CB15485F09DEC6C94EC4E05B5E3F002201641A5C0270C207FDFA89E1E194FE72974A6BF42EB86FBA931C3FC92CCD3AC6D","date":607062701,"hash":"B254CCF198156440230DC9E2F89FB6A3FA87075DE9009BB6EF29E80376E33DD0"},"type":"transaction","validated":true}`
	res := parseMessage(exampleXrpTxResult)
	t.Log(res)

	if res.Transaction.Hash != "B254CCF198156440230DC9E2F89FB6A3FA87075DE9009BB6EF29E80376E33DD0" {
		t.Error("did not parse hash correctly")
	}
}

func TestHandleNewXrpTx(t *testing.T) {
	exampleXrpTxResult := `{"engine_result":"tesSUCCESS","engine_result_code":0,"engine_result_message":"The transaction was applied. Only final in a validated ledger.","ledger_hash":"0DB3566DC7E57228DB6E2D8A0A697C1B17A8F7251559B22966FD723760075F18","ledger_index":18139963,"meta":{"AffectedNodes":[{"ModifiedNode":{"FinalFields":{"Account":"rsGPNkSLt36BDLMgPAYKifFvCphQJZ2qJw","Balance":"1999999892","Flags":0,"OwnerCount":0,"Sequence":10},"LedgerEntryType":"AccountRoot","LedgerIndex":"131B6989F00BE80F0C5ACA4C642CBD534BB18E2F04E59488F9BAF9ED7017E140","PreviousFields":{"Balance":"2999999904","Sequence":9},"PreviousTxnID":"509FC6682AAE4AFFFB0123EE812AD5A7CEC2F68188B045786841F0DF662750F1","PreviousTxnLgrSeq":18139935}},{"ModifiedNode":{"FinalFields":{"Account":"rs16hESfGChwAnK97oSdRJq4A18gcJbE7j","Balance":"17999996952","Flags":0,"OwnerCount":1,"Sequence":4},"LedgerEntryType":"AccountRoot","LedgerIndex":"FEE0D27B64E9EC97D880E34C172EBEC49147D2D156169BC7A63430EAFD214790","PreviousFields":{"Balance":"16999996952"},"PreviousTxnID":"509FC6682AAE4AFFFB0123EE812AD5A7CEC2F68188B045786841F0DF662750F1","PreviousTxnLgrSeq":18139935}}],"TransactionIndex":0,"TransactionResult":"tesSUCCESS","delivered_amount":"1000000000"},"status":"closed","transaction":{"Account":"rsGPNkSLt36BDLMgPAYKifFvCphQJZ2qJw","Amount":"1000000000","Destination":"rs16hESfGChwAnK97oSdRJq4A18gcJbE7j","Fee":"12","Flags":2147483648,"LastLedgerSequence":18139964,"Memos":[{"Memo":{"MemoData":"7573647861646472727347504E6B534C74333642444C4D675041594B69664676437068514A5A32714A77"}}],"Sequence":9,"SigningPubKey":"02CA3BD0848A81769A35BA0A663098A75A9D7182E8A40418788FDCDFABDA67C660","TransactionType":"Payment","TxnSignature":"3045022100D8525293800A845B561280B8B90AAA08D44CB15485F09DEC6C94EC4E05B5E3F002201641A5C0270C207FDFA89E1E194FE72974A6BF42EB86FBA931C3FC92CCD3AC6D","date":607062701,"hash":"B254CCF198156440230DC9E2F89FB6A3FA87075DE9009BB6EF29E80376E33DD0"},"type":"transaction","validated":true}`
	validatorAddress := []byte{}
	jobQueue := make(chan Job, 100)

	err := handleNewXrpTx(jobQueue, validatorAddress, exampleXrpTxResult)

	if err != nil {
		t.Error(err)
	}
	if len(jobQueue) != 1 {
		t.Error("didn't receive the expected number of jobs")
	}
}

func TestWorker(t *testing.T) {
	jobQueue := make(chan Job, 10)
	done := make(chan bool)
	runner := func(j Job) {
		t.Logf("Processed job %s", j)
		done <- true
	}
	go worker(jobQueue, runner)

	jobQueue <- Job{}
	<-done // wait until job has been processed
}

func TestCompleteAndBroadcastTxCLI(t *testing.T) {

	err := CompleteAndBroadcastTxCLI(authtxb.TxBuilder{}, context.CLIContext{}, "password", Job{})

	if err == nil {
		t.Error("should have errored")
	}
	t.Log(err)
}
