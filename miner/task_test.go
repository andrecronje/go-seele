/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package miner

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/core/types"
	"github.com/seeleteam/go-seele/crypto"
)

func Test_PopBestTx(t *testing.T) {
	txs := make(map[common.Address][]*types.Transaction)

	addr1 := *crypto.MustGenerateRandomAddress()
	t1, _ := types.NewTransaction(addr1, addr1, big.NewInt(0), big.NewInt(11), 10)
	t2, _ := types.NewTransaction(addr1, addr1, big.NewInt(0), big.NewInt(5), 11)
	txs[addr1] = []*types.Transaction{t1, t2}

	addr2 := *crypto.MustGenerateRandomAddress()
	t3, _ := types.NewTransaction(addr2, addr2, big.NewInt(0), big.NewInt(10), 10)
	t4, _ := types.NewTransaction(addr2, addr2, big.NewInt(0), big.NewInt(12), 11)
	txs[addr2] = []*types.Transaction{t3, t4}

	tt1 := popBestFeeTx(txs)
	assert.Equal(t, tt1, t1, "1")
	fmt.Println(tt1.Data.Fee)

	tt2 := popBestFeeTx(txs)
	assert.Equal(t, tt2, t3, "2")
	fmt.Println(tt2.Data.Fee)

	tt3 := popBestFeeTx(txs)
	assert.Equal(t, tt3, t4, "3")
	fmt.Println(tt3.Data.Fee)

	tt4 := popBestFeeTx(txs)
	assert.Equal(t, tt4, t2, "4")
	fmt.Println(tt4.Data.Fee)

	tt5 := popBestFeeTx(txs)
	assert.Equal(t, tt5, (*types.Transaction)(nil), "5")
}
