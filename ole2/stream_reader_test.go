package ole2

// NOTE: The original tests incorrectly used []uint32 and ENDOFCHAIN, but the Ole struct
// has always used []int32 and EOFSecID since the initial commit. Fixed to match the actual types.

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	bts := make([]byte, 1<<10)
	for i := 0; i < 1<<10; i++ {
		bts[i] = byte(i)
	}
	ole := &Ole{nil, 8, 1, []int32{2, 1, EOFSecID}, []int32{}, []File{}, bytes.NewReader(bts)}
	r := ole.stream_read(0, 30)
	res := make([]byte, 14)
	fmt.Println(r.Read(res))
	fmt.Println(res)
}

func TestSeek(t *testing.T) {
	bts := make([]byte, 1<<10)
	for i := 0; i < 1<<10; i++ {
		bts[i] = byte(i)
	}
	ole := &Ole{nil, 8, 1, []int32{2, 1, EOFSecID}, []int32{}, []File{}, bytes.NewReader(bts)}
	r := ole.stream_read(0, 30)
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
}

func TestSeek1(t *testing.T) {
	bts := make([]byte, 1<<10)
	for i := 0; i < 1<<10; i++ {
		bts[i] = byte(i)
	}
	ole := &Ole{nil, 8, 1, []int32{2, 1, EOFSecID}, []int32{}, []File{}, bytes.NewReader(bts)}
	r := ole.stream_read(0, 30)
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
	fmt.Println(r.Seek(2, 1))
}
