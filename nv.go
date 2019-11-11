//Encryption is mandatory for all payout devices and optional for
//pay in devices. Encrypted data and commands are transported
//between the host and the slave(s) using the transport mechanism
//described above, the encrypted information is stored in the data
//field in the format shown below

//+------------------------------------------------------------+
//|                     Encryption Layer                       |
//+------------------------------------------------------------+
//
//+------+----------------+----------+--------+--------+-------+
//| STX  |  SEQ/SLAVE ID  |  LENGTH  |  DATA  |  CRCL  |  CRCH |
//+------+----------------+----------+--------+--------+-------+
//
//
//+------------------------------------------------------------+
//|DATA                                                        |
//+---------+--------------------------------------------------+
//|STEX     |             Encrypted Data                       |
//+---------+--------------------------------------------------+
//
//
//+------------------------------------------------------------+
//|Encrypted Data                                              |
//+---------+----------+---------+------------+---------+------+
//|eLENGTH  |  eCOUNT  |  eDATA  |  ePACKING  |  eCRCL  | eCRCH|
//+---------+----------+---------+------------+---------+------+

package nv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/tarm/serial"
	"io"
	"log"
	"sync"
	"time"
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

	UnitData    *UnitData
	ChannelData *[]ChannelData
}

type ChannelData struct {
	Value     uint16
	Channel   byte
	Currency  []byte
	Level     uint16
	Recycling bool
}

type UnitData struct {
	UnitType        string
	FirmwareVersion string
	CountryCode     string
	ValueMultiplier uint16
	ProtocolVersion uint16
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
		return err
	}

	log.Printf("[INFO] Connect:")

	s.port = op
	return err

}

func (s *Service) Disconnect() (err error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.port != nil {
		err := s.port.Close()
		if err != nil {
			log.Printf("[ERROR]")
			return err
		}
	}

	return
}

func (s *Service) ResetFixedEncryptionKey() (*Response, error) {

	//Encryption Required:
	//No

	//Supported on devices:
	//SMART Hopper, SMART Payout, NV11

	//Description:
	//Resets the fixed encryption key to the device default.
	//The device may have extra security requirements before
	//it will accept this command (e.g. The Hopper must be empty)
	//if these requirements are not met, the device will reply
	//with Command Cannot be Processed. If successful, the device
	//will reply OK, then reset. When it starts up the fixed
	//key will be the default.

	//Example
	//7F 80 01 61 46 03
	//7F 80 01 F0 23 80

	return
}

func (s *Service) SetFixedEncryptionKey() (*Response, error) {

	//Encryption Required:
	//No

	//Supported on devices:
	//SMART Hopper, SMART Payout, NV11

	//Description:
	//A command to allow the host to change the fixed part of
	//the eSSP key. The eight data bytes are a 64 bit number
	//representing the fixed part of the key. This command must
	//be encrypted.

	return
}

func (s *Service) EnablePayoutDevice() (*Response, error) {

	//Encryption Required:
	//No

	//Supported on devices:
	//SMART Payout, NV11

	//Description:
	//A command to enable the attached payout device for
	//storing/paying out notes. A successful enable will return
	//OK, If there is a problem the reply will be generic
	//response COMMAND_CANNOT_BE_PROCESSED, followed by
	//an error code.

	//Example Fail Response:
	//7F 80 02 F5 02 30 3E
	//The device responds with COMMAND CANNOT BE PROCESSED
	//and an error byte for failure to enable. In this
	//example an invalid currency miss-match was detected
	//etween the validator and connected payout.

	//Payout Enable Error codes
	//+-----------------------------------+---------------+
	//|             Error reason          |  Error code   |
	//+---------------------------------------------------+
	//|  No device connected              | 1             |
	//+---------------------------------------------------+
	//|  Invalid currency detected        | 2             |
	//+---------------------------------------------------+
	//|  Device busy                      | 3             |
	//+---------------------------------------------------+
	//|  Empty only (Note float only)     | 4             |
	//+---------------------------------------------------+
	//|  Device error                     | 5             |
	//+-----------------------------------+---------------+

	return
}

func (s *Service) DisablePayoutDevice() (*Response, error) {

	//Encryption Required:
	//No

	//Supported on devices:
	//NV11

	//Description:
	//All accepted notes will be routed to the stacker
	//and payout commands will not be accepted.

	return
}

//Coin Mech Option

func (s *Service) ResetCounters() (*Response, error) {

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB SMART Payout NV11

	//Description:
	//Resets the note activity counters described in Get Counters
	//command to all zero values

	return
}

func (s *Service) GetCounters() (*Response, error) {

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB SMART Payout NV11

	//Description:
	//A command to return a global note activity counter set for
	// the slave device. The response is formatted as in the table
	// below and the counter values are persistent in memory after
	// a power down- power up cycle. These counters are note set
	// independent and will wrap to zero and begin again if their
	// maximum value is reached. Each counter is made up of
	// 4 bytes of data giving a max value of 4294967295.

	//Response
	//The device responds generic OK if supported and then with data
	//representing the cumulative note count for the device. In this
	//example, we have a device with 300 notes stacked, 210 notes stored,
	//180 notes dispensed and 25 notes rejected.

	//7F 80 16 F0 05 2C 01 00 00 D2 00 00 00 B4 00 00 00 68 01 00 00 19
	//00 00 00 F1 82

	//Response Schema
	//+-------------------+---------------+---------------------------------------+
	//|  Data byte offset |   Size bytes  |              Description              |
	//+---------------------------------------------------------------------------+
	//|      0            | 1             |Number of counters in set              |
	//+---------------------------------------------------------------------------+
	//|      1|4          | 4             |Notes stacked                          |
	//+---------------------------------------------------------------------------+
	//|      5|8          | 4             |Notes stored                           |
	//+---------------------------------------------------------------------------+
	//|      9|12         | 4             |Notes dispensed                        |
	//+---------------------------------------------------------------------------+
	//|      14|16        | 4             |Notes transferred from store to stacker|
	//+---------------------------------------------------------------------------+
	//|      17|20        | 4             |Notes rejected                         |
	//+-------------------+-------------------------------------------------------+

	return
}

func (s *Service) EventACK() (*Response, error) {

	//Encryption Required:
	//Yes

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 SMART Hopper NV11

	//Description:
	//This command will clear a repeating Poll ACK response
	//and allow further note operations

	return
}

func (s *Service) Sync() (*Response, error) {

	//Description:
	//A command to establish communications with a slave device.
	//A Sync command resets the seq bit of the packet so that the
	//slave device expects the next seq bit to be 0. The host then
	//sets its next seq bit to 0 and the seq sequence is synchronised

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB, NV10USB, BV20, BV50, BV100, NV200,
	//SMART Hopper, SMART Payout, NV11

	log.Printf("[INFO] Sync:")

	cmd, err := s.command(CMD_SYNC, []byte{})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) SetGenerator() (*Response, error) {

	//Description:
	//Eight data bytes are a 64 bit number representing
	//the Generator this must be a 64bit prime number.
	//The slave will reply with OK or PARAMETER_OUT_OF_RANGE
	//if the number is not prime.

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB, NV10USB, BV20, BV50, BV100,
	//NV200, SMART Hopper, SMART Payout, NV11

	log.Printf("[INFO] SetGenerator:")

	data := make([]byte, 8)
	//prime number: 982451653
	binary.LittleEndian.PutUint64(data, uint64(982451653))

	cmd, err := s.command(CMD_SET_GENERATOR, data)
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) SetModulus() (*Response, error) {

	//Description:
	//Eight data bytes are a 64 bit number representing the
	//modulus this must be a 64 bit prime number. The slave
	//will reply with OK or PARAMETER_OUT_OF_RANGE if the number
	//is not prime

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 SMART Hopper SMART Payout NV11

	log.Printf("[INFO] SetModulus:")

	data := make([]byte, 8)
	//prime number: 982451653
	binary.LittleEndian.PutUint64(data, uint64(1287821))

	cmd, err := s.command(CMD_SET_MODULUS, []byte{})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) HostProtocolVersion() (*Response, error) {

	//Description:
	//Dual byte command, the first byte is the command; the second
	//byte is the version of the protocol that is implemented on
	//the host. So for example, to enable events on BNV to protocol
	//version 6, send 06, 06. The device will respond with OK if the
	//device supports version 6, or FAIL (0xF8) if it does not.

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 SMART Hopper SMART Payout NV11

	log.Printf("[INFO] HostProtocolVersion:")

	data := make([]byte, 1)
	data[0] = 0x06

	cmd, err := s.command(CMD_HOST_PROTOCOL_VERSION, data)
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) RequestKeyExchange() (*Response, error) {

	//Description:
	//The eight data bytes are a 64 bit number representing
	//the Host intermediate key. If the Generator and Modulus
	//have been set the slave will calculate the reply with
	//the generic response and eight data bytes representing
	//the slave intermediate key. The host and slave will
	//then calculate the key. If Generator and Modulus are
	//not set then the slave will reply FAIL

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 SMART Hopper SMART Payout NV11

	log.Printf("[INFO] RequestKeyExchange:")

	data := make([]byte, 8)
	//Host intermediate: 982451653
	binary.LittleEndian.PutUint64(data, uint64(982451653))

	cmd, err := s.command(CMD_REQUEST_KEY_EXCHANGE, []byte{})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) ChannelValueRequest() (*Response, error) {

	//Description:
	//Returns channel value data for a banknote validator. Formatted
	//as: byte 0 - the highest channel used the a value byte representing
	//each of the denomination values. The real value is obtained
	//by multiplying by the value multiplier. If the validator is greater
	//than or equal to protocol version 6 then the channel values response
	//will be given as: Highest Channel, Value Per Channel (0 for expanded values),
	//3 Byte ASCI country code for each channel, 4- byte Full channel Value
	//for each channel.

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 NV11

	log.Printf("[INFO] ChannelValueRequest:")

	r, err := s.command(CMD_CHANNEL_VALUE_REQUEST, []byte{})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	//Highest Channel
	_ = r.Data[4]

	return r, nil
}

func (s *Service) UnitData() (*Response, error) {

	//Description:
	//Returns, Unit type (1 Byte integer), Firmware Version (4 bytes ASCII string),
	//Country Code (3 Bytes ASCII string), Value Multiplier (3 bytes integer),
	// Protocol Version (1 Byte, integer)

	//Encryption Required:
	//No

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 NV11

	log.Printf("[INFO] UnitData:")

	r, err := s.command(CMD_UNIT_DATA, []byte{})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	var unitType string
	switch ut := r.Data[4]; ut {
	case 0x00:
		unitType = "Validator"
	case 0x03:
		unitType = "SMART Hopper"
	case 0x06:
		unitType = "SMART Payout"
	case 0x07:
		unitType = "NV11"
	default:
		unitType = "Unknown Type"
	}

	var firmwareVersion string
	for _, item := range r.Data[5:9] {
		firmwareVersion += string(item)
	}

	var country string
	for _, item := range r.Data[9:12] {
		country += string(item)
	}

	valueMultiplier := uint16(r.Data[14]*100 + r.Data[13]*10 + r.Data[12]*1)
	protocolVersion := uint16(r.Data[15])

	r.UnitData = &UnitData{
		UnitType:        unitType,
		FirmwareVersion: firmwareVersion,
		CountryCode:     country,
		ValueMultiplier: valueMultiplier,
		ProtocolVersion: protocolVersion,
	}

	return r, nil
}

func (s *Service) Enable() (*Response, error) {

	//Supported on devices:
	//NV9USB NV10USB BV20 BV50 BV100 NV200 SMART Hopper NV11

	//Encryption Required:
	//No

	//Send this command to enable a disabled device.

	log.Printf("[INFO] Enable:")

	cmd, err := s.command(CMD_ENABLE, []byte{})
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) ConfigureBezel() (*Response, error) {

	//Description:
	//This command allows the host to configure a supported BNV bezel.
	//If the bezel is not supported the command will return generic
	//response COMMAND NOT KNOWN 0xF2.

	//Encryption Required:
	//No

	//Supported on devices:
	//NV200

	log.Printf("[INFO] ConfigureBezel:")

	data := make([]byte, 4)
	data[0] = 0x00 //Red intensity (0-255)
	data[1] = 0x00 //Green intensity (0-255)
	data[2] = 0x00 //Blue intensity (0-255)
	data[3] = 0x00 //Config 0 for volatile,1 - for non-volatile.

	cmd, err := s.command(CMD_CONFIGURE_BEZEL, data)
	if err != nil {
		log.Printf("[ERROR]")
		return nil, err
	}

	return cmd, nil
}

func (s *Service) command(cmd byte, data []byte) (*Response, error) {

	log.Printf("[INFO] Command:")

	response, err := s.request(cmd, data)
	if err != nil {
		log.Printf("[ERROR]")
		return response, err
	}

	return response, nil
}

func (s *Service) request(cmd byte, data []byte) (*Response, error) {

	log.Printf("[INFO] Request:")

	if seq == 0x80 {
		seq = 0x00
	} else {
		seq = 0x80
	}

	len := byte(len(data) + 1)

	var b bytes.Buffer
	//b := new(bytes.Buffer)
	b.WriteByte(STX)
	b.WriteByte(seq)
	b.WriteByte(len)
	b.WriteByte(cmd)
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
