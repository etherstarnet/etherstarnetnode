// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

var MainnetBootnodes = []string{
	// Eth Star Network Boot Nodes Go Bootnodes
	"enode://a6ee2eaca69b93e630e9067b142fc7a803cf6361fd7de03b5071243a69b71b6878a9786b873658eb13345a2e92c62e0ce62e90f8e68bbccd84814fe0b3e90274@165.227.42.184:30303",
	"enode://34ec67434a183c1e8ed698d6fab70453b7731a77045608c0f3eab7f54b9d39c8dec8859a29608b462e7d65f0802120bf50354fda69770fb929a6a1ca3537a7be@68.183.24.195:30303",
	"enode://0c233552b6379f5e074386959d5f49ce61cda2dc2b2fc81c67fbc06c7c1836221a5b48056aa79bf0d99ddb2538d2d8145db50df71d457f732ebeb1b7819bca75@165.227.210.22:30303",
	"enode://93c306167c0c61dbe19210aed9f5d34753c177923b836c77f0bb895b677fe732ab6274927d2b32384539ed2868abc7d0639a576feb208a5204044ae00540ace8@94.156.174.83:30303",
	"enode://7f06a3803daab395ebc30031b23b15ca2399089c107b546f2b27570aada092dd68d5f884fab280764c90b6eb7a89361b7583b427e4275847ad44887fee8fb85b@94.156.174.81:30303",
	"enode://013ccebadb6dad06e1bf0bfff3133c4fd5e8286929de488bce080289702b937b4b5ebb3f24d8f17b697bb4259f9847dc4658dafc7964555ca9589443f4f70e1c@20.195.207.140:30303",
}
var RopstenBootnodes = []string{
}
var SepoliaBootnodes = []string{
}
var RinkebyBootnodes = []string{
}
var GoerliBootnodes = []string{
}
var KilnBootnodes = []string{
}
var V5Bootnodes = []string{
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
