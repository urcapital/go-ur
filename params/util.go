// Copyright 2015 The go-ur Authors
// This file is part of the go-ur library.
//
// The go-ur library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ur library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ur library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/ur-technology/go-ur/common"
)

var (
	TestNetGenesisHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d") // Testnet genesis hash to enforce below configs on
	MainNetGenesisHash = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3") // Mainnet genesis hash to enforce below configs on

	TestNetHomesteadBlock = big.NewInt(1) // Testnet homestead block
	MainNetHomesteadBlock = big.NewInt(1) // Mainnet homestead block

	TestNetHomesteadGasRepriceBlock = big.NewInt(1)      // Testnet gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(444000) // Mainnet gas reprice block

	// TestNetHomesteadGasRepriceHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d") // Testnet gas reprice block hash (used by fast sync)
	TestNetHomesteadGasRepriceHash = common.Hash{} // Testnet gas reprice block hash (used by fast sync)
	// MainNetHomesteadGasRepriceHash = common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0") // Mainnet gas reprice block hash (used by fast sync)
	MainNetHomesteadGasRepriceHash = common.HexToHash("5c45acfc654722b17e701fe566a16a80f99e932f34266041e3de28e21239e057") // Mainnet gas reprice block hash (used by fast sync)

	// GBU
	MainNetGBUBlock     = big.NewInt(1)
	MainNetGBUBlockHash = common.Hash{}

	TestNetSpuriousDragon = big.NewInt(1)
	MainNetSpuriousDragon = big.NewInt(444500)

	TestNetChainID = big.NewInt(3) // Test net default chain ID
	MainNetChainID = big.NewInt(1) // main net default chain ID
)
