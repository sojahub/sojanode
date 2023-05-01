package symbol_translator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	sojahubDenomFeedface = "ibc/FEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACE"
	ethereumSymbolFeeface = "Face"
)

func TestNewSymbolTranslatorFromJsonBytes(t *testing.T) {
	_, err := NewSymbolTranslatorFromJSONBytes([]byte("foo"))
	assert.Error(t, err)

	q := ` {"ibc/FEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACE": "Face"} `
	x, err := NewSymbolTranslatorFromJSONBytes([]byte(q))
	assert.NoError(t, err)
	assert.NotNil(t, x)
	assert.Equal(t, x.SojahubToEthereum(sojahubDenomFeedface), ethereumSymbolFeeface)
	assert.Equal(t, x.EthereumToSojahub(ethereumSymbolFeeface), sojahubDenomFeedface)
	assert.Equal(t, x.SojahubToEthereum("verbatim"), "verbatim")
	assert.Equal(t, x.EthereumToSojahub("verbatim"), "verbatim")
}

func TestNewSymbolTranslator(t *testing.T) {
	s := NewSymbolTranslator()
	assert.Equal(t, s.SojahubToEthereum("something"), "something")
}
