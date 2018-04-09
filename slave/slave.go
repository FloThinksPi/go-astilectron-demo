package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	token := make([]byte, 32)
	rand.Read(token)

	crypt, err := kcp.NewAESBlockCrypt([]byte("1234567890123456"))
	if err != nil {
		panic(err)
	}

	l, err := kcp.ListenWithOptions("0.0.0.0:8100", crypt, 10, 3)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	timeout := time.Second * 30

	fmt.Println("running...")
	for {
		conn, err := l.AcceptKCP()
		if err != nil {
			fmt.Println(err)
			continue
		}

		configureKCPConnection(conn, timeout)
		msg, err := recvData(conn)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("received:", string(msg))
		}

		n, e := sendData(conn, []byte("Response"))

		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println("wrote", n, "bytes")
		}

		conn.Close()
	}
}

func configureKCPConnection(conn *kcp.UDPSession, timeout time.Duration) {
	conn.SetStreamMode(true)
	conn.SetWindowSize(512, 512)
	conn.SetNoDelay(1, 40, 2, 1)
	conn.SetACKNoDelay(false)

	conn.SetReadDeadline(time.Now().Add(timeout))
	conn.SetWriteDeadline(time.Now().Add(timeout))
}

func recvData(r io.Reader) ([]byte, error) {
	buf := make([]byte, 4)

	_, err := r.Read(buf)
	if err != nil {
		return nil, err
	}

	if len(buf) == 0 {
		return buf, nil
	}

	// read message size as 4 bytes from the beginning of the message
	size := int(binary.LittleEndian.Uint32(buf))

	mbuf := bytes.NewBuffer([]byte{})

	buf = make([]byte, 4096)

	total := 0
	for total < size {
		n, err := r.Read(buf)
		if err != nil {
			return nil, err
		}
		mbuf.Write(buf[:n])
		total += n
		if err != nil {
			break
		}
	}

	return mbuf.Bytes(), nil
}

func sendData(w io.Writer, msg []byte) (int, error) {
	b := []byte{0, 0, 0, 0}
	binary.LittleEndian.PutUint32(b, uint32(len(msg)))

	mbuf := bytes.NewBuffer(b)
	mbuf.Write(msg)

	return w.Write(mbuf.Bytes())
}
