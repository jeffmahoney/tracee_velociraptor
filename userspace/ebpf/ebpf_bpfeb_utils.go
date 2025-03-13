//go:build armbe || arm64be || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64

package ebpf

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func getEbpfBytes() []byte {
	rb := bytes.NewReader(FileUserspaceEbpfEbpfBpfebO)
	r, err := gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return data
}
