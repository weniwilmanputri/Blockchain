  // Copyright 2017 The go-ethereum Authors
  // This file is part of the go-ethereum library.

  // The go-ethereum library is free software: you can redistribute it and/or modify
  // it under the terms of the GNU Lesser General Public License as published by
  // the Free Software Foundation, either version 3 of the License, or
  // (at your option) any later version.

  // The go-ethereum library is distributed in the hope that it will be useful,
  // but WITHOUT ANY WARRANTY; without even the implied warranty of
  // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
  // GNU Lesser General Public License for more details.
 
  // You should have received a copy of the GNU Lesser General Public License
  // along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
  // paket consesus menggunakan implementasi ethereum consesnsus engine yang berbeda
package consensus

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)
  // ChainHeader Reader mendefeinisikan kumpulan method yang membutuhkan akses ke lokal
type ChainHeaderReader interface {
  // configurasi untuk mengambil chain configurasi
	Config() *params.ChainConfig

  // currentHeader mengambil data header sebelumnya di lokal chain
	CurrentHeader() *types.Header

	// GetHeader mengambil data block header dari database dengan hash dan nomor
	GetHeader(hash common.Hash, number uint64) *types.Header

	// GetHeaderByNumber mengambil data block header dari database number
	GetHeaderByNumber(number uint64) *types.Header

	// GetHeaderByHash mengambil data block header dari database its hash
	GetHeaderByHash(hash common.Hash) *types.Header

	//GetTd mengambil data total kesulitan dari database hash dan number 
	GetTd(hash common.Hash, number uint64) *big.Int
}

  // ChainReader mendefinisi kumpulan method yang dibutuhkan akses ke lokal
type ChainReader interface {
	ChainHeaderReader

	//GetBlock mengambil data total kesulitan dari database hash dan number 
	GetBlock(hash common.Hash, number uint64) *types.Block
}

  // Engine merupakan algoritma agnistik konsesus engine
type Engine interface {
  // Author  mengambil alamat ethereum dari akun
	Author(header *types.Header) (common.Address, error)

  // VerifyHeader mencek apakah header sesuai dengan aturan
	VerifyHeader(chain ChainHeaderReader, header *types.Header, seal bool) error

	// VerifyHeaders mirip dengan method VerifyHeader, tetapi VerifyHeaders memverifikasi sekumpulan header secara bersamaan 
	VerifyHeaders(chain ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error)

  // Prepare menginisialisasi bidang konsensus dari blok header sesuai dengan aturan dari engine tertentu
	Prepare(chain ChainHeaderReader, header *types.Header) error

  // VerifyUncles memverifikasi block yang diberikan sudah sesuai dengan aturan dari yang diberikan engine
	VerifyUncles(chain ChainReader, block *types.Block) error

	// Finalize menjalankan berbagai post-transaction state modifikasi teatpi tidak menyatukan blok
	Finalize(chain ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
		uncles []*types.Header)

	// FinalizeAndAssemble menjalankan berbagai post-transaction state modifikasi dan menggabungkan blok terakhir
	FinalizeAndAssemble(chain ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
		uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error)

	// Seal menghasilkan permintaan penyegelan baru untuk blok input yang diberikan dan mendorong hasilnya ke saluran yang diberikan.
	Seal(chain ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error

	// SealHash kembali ke hash blok  sebelumnya untuk disegel
	SealHash(header *types.Header) common.Hash

	// CalcDifficulty merupakan algoritma pengaturan kesulitan
	CalcDifficulty(chain ChainHeaderReader, time uint64, parent *types.Header) *big.Int

	// APIs kembali RPC APIs pada konsensus engine  
	APIs(chain ChainHeaderReader) []rpc.API

	// Close merupakan perintah mengakhiri serbagai background tread dari consesus engine
	Close() error
}

  // PoW merupakan consesus engine dengan base proof-of-work
type PoW interface {
	Engine

	// Hashrate kembali ke mining hastrate sebelumnya dari PoW consensus engine
	Hashrate() float64
}
