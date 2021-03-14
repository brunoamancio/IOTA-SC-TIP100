package libtest

import (
	"testing"

	"github.com/brunoamancio/IOTA-SC-ERC721/tests/testutils"
	"github.com/brunoamancio/IOTA-SC-ERC721/tests/testutils/testconstants"
	notsolo "github.com/brunoamancio/NotSolo"
	"github.com/stretchr/testify/require"
)

//  -----------------------------------------------  //
//  See code samples in Tests/testutils/codesamples  //
//  -----------------------------------------------  //

func TestLib(t *testing.T) {
	contractWasmFilePath := testutils.MustGetContractWasmFilePath(t, testconstants.ContractName) // You can use if file is in SmartContract/pkg

	// Name of the SC view to be requested - Defined in smartcontract/iota_sc_utils/src/interfaces.rs > NAME_FUNC_IMPLEMENTS
	const name_function_implements = testconstants.NameFuncImplements
	// Name of the TIP-100 interface - Defined in lib.rs > add_view > implements
	const name_interface_tip100 = testconstants.NameInterfaceTip100
	// Name of the TIP-100 interface - Defined in smartcontract/iota_sc_utils/src/interfaces.rs > HNAME_INTERFACE_TIP_100
	hname_interface_tip100 := testutils.EncodeHName((uint32)(0xeae53bfb))

	notSolo := notsolo.New(t)

	chainName := testconstants.ContractName + "Chain"
	chain := notSolo.Chain.NewChain(nil, chainName)

	// Uploads wasm of SC and deploys it into chain
	notSolo.Chain.DeployWasmContract(chain, nil, testconstants.ContractName, contractWasmFilePath)

	// Call contract 'tip_100', function 'implements'
	response := notSolo.Request.MustView(chain, testconstants.ContractName, name_function_implements, name_interface_tip100, hname_interface_tip100)
	implements := notSolo.Data.MustGetBool(response[name_interface_tip100])

	require.True(t, implements)
}
