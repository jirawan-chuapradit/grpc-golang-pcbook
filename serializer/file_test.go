package serializer_test

import (
	"testing"

	"github.com/jirawan-chuapradit/grpc-golang-pcbook/sample"
	"github.com/stretchr/testify/require"
	"github.com/jirawan-chuapradit/grpc-golang-pcbook/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"


	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)


} 