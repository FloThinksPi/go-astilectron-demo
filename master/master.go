package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	kcp "github.com/xtaci/kcp-go"
)

func main() {
	limit := 10

	crypt, err := kcp.NewAESBlockCrypt([]byte("1234567890123456"))
	if err != nil {
		panic(err)
	}
	for i := 0; i < limit; i++ {

		conn, err := kcp.DialWithOptions("127.0.0.1:8100", crypt, 10, 3)
		if err != nil {
			panic(err)
		}

		timeout := time.Second * 30

		configureKCPConnection(conn, timeout)
		mausposx, mausposy := robotgo.GetMousePos()
		sendData(conn, []byte(strconv.Itoa(mausposx)+" "+strconv.Itoa(mausposy)))
		//msg, err := recvData(conn)

		//if err != nil {
		//	fmt.Println(err)
		//	} else {
		//		fmt.Println("message size recevied is", len(msg))
		//	}

		conn.Close()

		time.Sleep(time.Millisecond * 100)
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
