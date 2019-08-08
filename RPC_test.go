package bitcoin

import "testing"

func TestGetBlockChainInfo(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlockchainInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetConnectionCount(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	count, err := b.GetConnectionCount()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%d", count)
}

func TestGetNetworkInfo(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetNetworkInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetNetTotals(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetNetTotals()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}
func TestMiningInfo(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetMiningInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestUptime(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	count, err := b.Uptime()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%d", count)
}

func TestGetPeerInfo(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetPeerInfo()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetRawMempool(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetRawMempool()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetChainTxStats(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	stats, err := b.GetChainTxStats(0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", stats)
}

func TestValidateAddress(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	stats, err := b.ValidateAddress("13q47QSaXBHMZVFHFENtTpfd7rtaWTe3v1")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", stats)
}

func TestHelp(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetHelp()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestGetBestBlockHash(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBestBlockHash()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestGetBlockHash(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlockHash(555432)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestSendRawTransaction(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.SendRawTransaction("0100000002b212995825c72d1abf5505a3017694f564dda29908fe5d45ab45b6a3e1748ca5010000006a473044022030d77fbaf3ec092a5038e7fc0235bb2126ca5099dccac5383d58cea01a8d329302207581678aa91743abe31c4cf54e16b0c58e05631d20ccde33be242d96cb57ff1141210373b675c91c95391b6d7977de12bf081bae193470aa0199efc1f847d799497b67ffffffff428947c3e5fb13faf1e71ea86b209b88c3b2e1cd77b326b696fc6e9fb3d0ced6000000006b483045022100801a316894dc40d1e335a14d25b6a7b5b72691b0842878105942d86303a09e1502204d737aecf5d1972e880e329363681bb7c97334e262cf4a20293fee3d8ca87d7041210373b675c91c95391b6d7977de12bf081bae193470aa0199efc1f847d799497b67ffffffff030a1a0000000000001976a914dc9ad4971a54b52308fa3c958df73eac52fb552f88acc6160000000000001976a914a933500e7326f81a974d4212aa16ae29f92e257188ac00000000000000003d6a3b53656e642076696120612057656368617420626f74206d616465206279206161726f6e363720687474703a2f2f6269742e6c792f333331536d754300000000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%s", res)
}

func TestGetBlock(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	res, err := b.GetBlock("000000000000000005e827eecfc1b8cbb990f4ae458e748480d10b80458faf25")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", res)
}

func TestGetRawTransaction(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.GetRawTransaction("4dd023de1efdd78129780da35ba35411868480fc2240ddd7ba042823ae02cdc9")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", tx)
}

func TestGetRawTransactionHex(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := b.GetRawTransactionHex("4dd023de1efdd78129780da35ba35411868480fc2240ddd7ba042823ae02cdc9")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", tx)
}

func TestGetBlockTemplate(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.GetBlockTemplate()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestGetMiningCandidate(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.GetMiningCandidate()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestSubmitBlock(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	template, err := b.SubmitBlock("00000020c803e566bf5af601f19ecb1878f5a6bc2188841c46530c0200000000000000007493d05e5f3ef63975027947bba8c2e08aae8b74e267d9e11e6a4d10e0c0ca3770223e5dc53c081801991f080101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff5f031c0d09046f223e5d2f706f6f6c696e2e636f6d2ffabe6d6ddd5fe0e090d85e6b39d0a6a19bd9b2ff3e3a05113d8ec3bd4a23e39e9c70aec00100000000000000736899c85c811ff6d3eade8b24b8f71213c3bb6d4d003843004000000000ffffffff02807c814a000000001976a91431f2bece272b5d346aa56d09101dd7306d9a307588ac0000000000000000266a24b9e11b6d4ee6b780e5c5ec8a57bc28030aaf0b43632f319648f8bca6abf1ebe3739511f0502c1759")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}

func TestSubmitMiningSolution(t *testing.T) {
	b, err := New("localhost", 8332, "bitcoin", "Yv5Nua9wLQyhHEUyHtSecMawAEgFlLp4s", false)
	if err != nil {
		t.Fatal(err)
	}

	var (
		miningCandidateID = "ABC"
		nonce             uint32
		coinbase          string
		time              uint32
		version           uint32
	)

	template, err := b.SubmitMiningSolution(miningCandidateID, nonce, coinbase, time, version)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%#v", template)
}