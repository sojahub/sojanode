package sojagen

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/MakeNowJust/heredoc"

	"github.com/Sojahub/sojanode/tools/sojagen/key"
	"github.com/Sojahub/sojanode/tools/sojagen/network"
	"github.com/Sojahub/sojanode/tools/sojagen/node"
	"github.com/Sojahub/sojanode/tools/sojagen/utils"
)

type Sojagen struct {
	chainID *string
}

func NewSojagen(chainID *string) Sojagen {
	return Sojagen{
		chainID: chainID,
	}
}

func (s Sojagen) NewNetwork(keyringBackend string) *network.Network {
	return &network.Network{
		ChainID: *s.chainID,
		CLI:     utils.NewCLI(*s.chainID, keyringBackend),
	}
}

func (s Sojagen) NetworkCreate(count int, outputDir, startingIPAddress string, outputFile string) {
	net := network.NewNetwork(*s.chainID)
	summary, err := net.Build(count, outputDir, startingIPAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = ioutil.WriteFile(outputFile, []byte(*summary), 0600); err != nil {
		log.Fatal(err)
		return
	}
}

func (s Sojagen) NetworkReset(networkDir string) {
	if err := network.Reset(*s.chainID, networkDir); err != nil {
		log.Fatal(err)
	}
}

func (s Sojagen) NewNode(keyringBackend string) *node.Node {
	return &node.Node{
		ChainID: *s.chainID,
		CLI:     utils.NewCLI(*s.chainID, keyringBackend),
	}
}

func (s Sojagen) NodeReset(nodeHomeDir *string) {
	if err := node.Reset(*s.chainID, nodeHomeDir); err != nil {
		log.Fatal(err)
	}
}

func (s Sojagen) KeyGenerateMnemonic(name, password string) {
	newKey := key.NewKey(name, password)
	newKey.GenerateMnemonic()
	fmt.Println(newKey.Mnemonic)
}

func (s Sojagen) KeyRecoverFromMnemonic(mnemonic string) {
	newKey := key.NewKey("", "")
	if err := newKey.RecoverFromMnemonic(mnemonic); err != nil {
		log.Fatal(err)
	}

	fmt.Println(heredoc.Doc(`
		Address: ` + newKey.Address + `
		Validator Address: ` + newKey.ValidatorAddress + `
		Consensus Address: ` + newKey.ConsensusAddress + `
	`))
}
