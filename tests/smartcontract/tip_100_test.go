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

func Test_tip_100_compatible_and_incompatible_contracts(t *testing.T) {
	const notTip100ContractName = "not_" + testconstants.ContractName

	// Name of the SC view to be requested - Defined in smartcontract/iota_sc_utils/src/interfaces.rs > NAME_FUNC_IMPLEMENTS
	const name_function_implements = testconstants.NameFuncImplements
	// Name of the TIP-100 interface - Defined in lib.rs > add_view > implements
	const name_interface_tip100 = testconstants.NameInterfaceTip100
	// Name of the TIP-100 interface - Defined in smartcontract/iota_sc_utils/src/interfaces.rs > HNAME_INTERFACE_TIP_100
	hname_interface_tip100 := testutils.EncodeHName((uint32)(0xeae53bfb))

	notSolo := notsolo.New(t)

	chainName := testconstants.ContractName + "Chain"
	chain := notSolo.Chain.NewChain(nil, chainName)

	// Map contractName to contract file path
	contractNameToContract := make(map[string]string)

	// You can use if file is in SmartContract/rust/tip-100-compatible-sc/pkg
	contractNameToContract[testconstants.ContractName] = testutils.MustGetContractWasmFilePath(t, testconstants.ContractName, true)
	// You can use if file is in SmartContract/rust/tip-100-incompatible-sc/pkg
	contractNameToContract[notTip100ContractName] = testutils.MustGetContractWasmFilePath(t, notTip100ContractName, false)

	// Map contractName to expected value
	contractNameToTest := make(map[string]bool)

	// hexadecimal number and the expected result
	contractNameToTest[testconstants.ContractName] = true
	contractNameToTest[notTip100ContractName] = false

	for contractName, expectedResult := range contractNameToTest {

		t.Run(contractName, func(t *testing.T) {
			// Uploads wasm of the tip100-compatible SC and deploys them into chain
			notSolo.Chain.DeployWasmContract(chain, nil, contractName, contractNameToContract[contractName])

			// Call contract function 'implements'
			response, err := notSolo.Request.View(chain, contractName, name_function_implements, name_interface_tip100, hname_interface_tip100)

			// Expect tip100-compatible SC implements interface
			if expectedResult {
				implements := notSolo.Data.MustGetBool(response[name_interface_tip100])
				require.True(t, expectedResult, implements)
			} else {
				require.Nil(t, response)
				require.Error(t, err)
			}
		})
	}
	// notSolo.Chain.DeployWasmContract(chain, nil, notTip100ContractName, not_tip100_contractWasmFilePath)
}
