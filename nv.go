package nv

import (
	"bytes"
	"errors"
	"github.com/tarm/serial"
	"io"
	"log"
	"sync"
	"time"
)

const (
	BUFFER_MAX_LENGTH = 1024
)

var (
	seq byte = 0x80
)

type Config struct {
	BaudRate    int
	PortName    string
	Address     byte
	ReadTimeout time.Duration
}

type Service struct {
	mu         sync.Mutex
	config     *Config
	port       *serial.Port
	portIsOpen bool
}

type Response struct {
	ErrorCode    []byte
	ErrorMessage string
	Data         []byte
	DataLen      uint16
}

func NewService(config *Config) *Service {
	return &Service{
		config:     config,
		portIsOpen: false,
	}
}

func (s *Service) Connect() (err error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.portIsOpen {
		err := s.Disconnect()
		if err != nil {
			return err
		}
	}

	c := &serial.Config{
		Name:        s.config.PortName,
		Baud:        s.config.BaudRate,
		ReadTimeout: s.config.ReadTimeout,
	}

	op, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("[ERROR]")
		s.portIsOpen = false
		return err
	}

	log.Printf("[INFO] Connect:")

	s.port = op
	s.portIsOpen = true

	return err

}

func (s *Service) Disconnect() (err error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.portIsOpen {
		err := s.port.Close()
		if err != nil {
			log.Printf("[ERROR]")
			s.portIsOpen = false
			return err
		}
	}

	log.Printf("[INFO] Disconnect:")
	s.portIsOpen = false
	return
}

func (s *Service) Sync() (*Response, error) {

	//Supported on devices:
	//NV9USB, NV10USB, BV20, BV50, BV100, NV200,
	//SMART Hopper, SMART Payout, NV11

	//Encryption Required:
	//No

	//A command to establish communications with a slave device.
	//A Sync command resets the seq bit of the packet so that the
	//slave device expects the next seq bit to be 0. The host then
	//sets its next seq bit to 0 and the seq sequence is synchronised

	log.Printf("[INFO] Sync:")

	cmd, err := s.command([]byte{CMD_SYNC})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) Enable() (*Response, error) {

	log.Printf("[INFO] Enable:")

	cmd, err := s.command([]byte{0X11})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) command(data []byte) (*Response, error) {

	log.Printf("[INFO] Command:")

	response, err := s.request(data)
	if err != nil {
		log.Printf("[ERROR]")
		return response, err
	}

	return response, nil
}

func (s *Service) request(data []byte) (*Response, error) {

	log.Printf("[INFO] Request:")

	if seq == 0x80 {
		seq = 0x00
	} else {
		seq = 0x80
	}

	dataLen := byte(len(data))

	var b bytes.Buffer
	b.WriteByte(STX)
	b.WriteByte(seq)
	b.WriteByte(dataLen)
	b.Write(data)
	b.Write(crc16(b.Bytes()[1:]))

	if b.Len() > BUFFER_MAX_LENGTH {
		return nil, errors.New("")
	}

	_, err := s.pWrite(b.Bytes())
	buf, _, err := s.pRead()

	if err != nil {
		return nil, err
	}

	var response Response
	response.DataLen = uint16(buf[2])
	response.Data = buf

	return &response, nil

}

func (s *Service) pRead() ([]byte, int, error) {

	if !s.portIsOpen {
		return nil, 0, errors.New("[ERROR]")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	buf := make([]byte, BUFFER_MAX_LENGTH)

	time.Sleep(50 * time.Millisecond)
	i := 0
	for {
		readLen, err := s.port.Read(buf[i:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}

		if readLen == 0 {
			break
		}

		i += readLen
	}

	log.Printf("[INFO] Read: Read buffer:[% X] len:%v", buf[:i], i)

	return buf, i, nil

}

func (s *Service) pWrite(data []byte) (int, error) {

	if !s.portIsOpen {
		return 0, errors.New("[ERROR]")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	writeLine, err := s.port.Write(data)
	if err != nil {
		log.Printf("[ERROR] Write: Write error:%s", err)
		return 0, err
	}

	log.Printf("[INFO] Write: Write data:[% X] len:%v", data, len(data))

	return writeLine, nil
}

func crc16(data []byte) []byte {
	seed := uint16(0xFFFF)
	poly := uint16(0x8005)
	crc := seed

	for _, d := range data {
		crc ^= uint16(d) << 8
		for i := 0; i < 8; i++ {
			bit := (crc & 0x8000) != 0
			crc <<= 1
			if bit {
				crc ^= poly
			}
		}
	}

	b := [2]byte{
		byte(crc & 0xFF),
		byte((crc >> 8) & 0xFF),
	}

	return b[:]
}
