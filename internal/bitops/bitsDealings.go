package bitops

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"

	"github.com/horsedevours/compression-tool/internal/huffman"
)

func BitwiseWrite(file io.Reader, fileOut io.Writer, hTree huffman.HuffmanNode) error {
	treeBuffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&treeBuffer)
	encoder.Encode(hTree)
	treeBytes := treeBuffer.Bytes()
	treeLen := uint64(len(treeBytes))

	// First 8 bytes to store size of encoded Huffman tree
	_, err := fileOut.Write(binary.BigEndian.AppendUint64([]byte{}, treeLen))
	if err != nil {
		return err
	}
	_, err = fileOut.Write(treeBytes)
	if err != nil {
		return err
	}

	codeMap := hTree.BuildCodeMap()

	reader := bufio.NewReader(file)

	// For all bytes, read next byte, encode, write
	cache := byte(0)
	position := 0
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		code := codeMap[b]
		for _, b := range code {
			if position == 8 {
				fileOut.Write([]byte{cache})
				cache &= 0
				position = 0
			}
			if b {
				cache = cache << 1
				cache |= 1
			} else {
				cache = cache << 1
			}
			position++
		}
	}

	if position > 0 {
		fileOut.Write([]byte{cache})
	}
	// Last byte stores offset for next to last byte
	// in case not all bits were used to encode
	fileOut.Write([]byte{byte(position - 1)})

	return nil
}

func BitWiseRead(file io.Reader, inputSize int64, fileOut io.Writer) error {
	contents := make([]byte, inputSize)
	_, err := io.ReadFull(file, contents)
	if err != nil {
		return err
	}

	var treeLen uint64
	// Read length of encoded Huffman tree from first 8 bytes
	_, err = binary.Decode(contents[:8], binary.BigEndian, &treeLen)
	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}
	// Slice out the Huffman tree
	buffer.Write(contents[8 : 8+treeLen])
	hTree := huffman.HuffmanNode{}
	decoder := gob.NewDecoder(&buffer)
	decoder.Decode(&hTree)

	position := byte(7)
	node := hTree
	// Encoded text excludes final byte
	body := contents[8+treeLen : len(contents)-1]
	for i, b := range body {
		// Set position for last byte of body based on offset
		// stored in last byte of file
		if i == len(body)-1 {
			position = contents[len(contents)-1]
		}
		for {
			branch := (b >> position) & 1
			if branch == 0 {
				node = *node.Left
			} else {
				node = *node.Right
			}
			if node.IsLeaf {
				fileOut.Write([]byte{node.Val})
				node = hTree
			}
			if position > 0 {
				position--
			} else {
				position = 7
				break
			}
		}
	}

	return nil
}
