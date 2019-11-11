package nv

var (
	seq byte = 0x00
)

const (
	BUFFER_MAX_LENGTH = 1024
	STX byte = 0x7F
)

const (
	CMD_RESET                             byte = 0x01
	CMD_SET_CHANNEL_INHIBITS              byte = 0x02
	CMD_DISPLAY_ON                        byte = 0x03
	CMD_DISPLAY_OFF                       byte = 0x04
	CMD_SETUP_REQUEST                     byte = 0x05
	CMD_HOST_PROTOCOL_VERSION             byte = 0x06
	CMD_POLL                              byte = 0x07
	CMD_REJECT_BANKNOTE                   byte = 0x08
	CMD_DISABLE                           byte = 0x09
	CMD_ENABLE                            byte = 0x0A
	CMD_GET_SERIAL_NUMBER                 byte = 0x0C
	CMD_UNIT_DATA                         byte = 0x0D
	CMD_CHANNEL_VALUE_REQUEST             byte = 0x0E
	CMD_CHANNEL_SECURITY_DATA             byte = 0x0F
	CMD_CHANNEL_RE_TEACH_DATA             byte = 0x10
	CMD_SYNC                              byte = 0x11
	CMD_LAST_REJECT_CODE                  byte = 0x17
	CMD_HOLD                              byte = 0x18
	CMD_GET_FIRMWARE_VERSION              byte = 0x20
	CMD_GET_DATASET_VERSION               byte = 0x21
	CMD_GET_ALL_LEVELS                    byte = 0x22
	CMD_GET_BAR_CODE_READER_CONFIGURATION byte = 0x23
	CMD_SET_BAR_CODE_CONFIGURATION        byte = 0x24
	CMD_GET_BAR_CODE_INHIBIT_STATUS       byte = 0x25
	CMD_SET_BAR_CODE_INHIBIT_STATUS       byte = 0x26
	CMD_GET_BAR_CODE_DATA                 byte = 0x27
	CMD_SET_REFILL_MODE                   byte = 0x30
	CMD_PAYOUT_AMOUNT                     byte = 0x33
	CMD_SET_DENOMINATION_LEVEL            byte = 0x34
	CMD_GET_DENOMINATION_LEVEL            byte = 0x35
	CMD_COMMUNICATION_PASS_THROUGH        byte = 0x37
	CMD_HALT_PAYOUT                       byte = 0x38
	CMD_SET_DENOMINATION_ROUTE            byte = 0x3B
	CMD_GET_DENOMINATION_ROUTE            byte = 0x3C
	CMD_FLOAT_AMOUNT                      byte = 0x3D
	CMD_GET_MINIMUM_PAYOUT                byte = 0x3E
	CMD_EMPTY_ALL                         byte = 0x3F
	CMD_SET_COIN_MECH_INHIBITS            byte = 0x40
	CMD_GET_NOTE_POSITIONS                byte = 0x41
	CMD_PAYOUT_NOTE                       byte = 0x42
	CMD_STACK_NOTE                        byte = 0x43
	CMD_FLOAT_BY_DENOMINATION             byte = 0x44
	CMD_SET_VALUE_REPORTING_TYPE          byte = 0x45
	CMD_PAYOUT_BY_DENOMINATION            byte = 0x46
	CMD_SET_COIN_MECH_GLOBAL_INHIBIT      byte = 0x49
	CMD_SET_GENERATOR                     byte = 0x4A
	CMD_SET_MODULUS                       byte = 0x4B
	CMD_REQUEST_KEY_EXCHANGE              byte = 0x4C
	CMD_SET_BAUD_RATE                     byte = 0x4D
	CMD_GET_BUILD_REVISION                byte = 0x4F
	CMD_SET_HOPPER_OPTIONS                byte = 0x50
	CMD_GET_HOPPER_OPTIONS                byte = 0x51
	CMD_SMART_EMPTY                       byte = 0x52
	CMD_CASHBOX_PAYOUT_OPERATION_DATA     byte = 0x53
	CMD_CONFIGURE_BEZEL                   byte = 0x54
	CMD_POLL_WITH_ACK                     byte = 0x56
	CMD_EVENT_ACK                         byte = 0x57
	CMD_GET_COUNTERS                      byte = 0x58
	CMD_RESET_COUNTERS                    byte = 0x59
	CMD_COIN_MECH_OPTIONS                 byte = 0x5A
	CMD_DISABLE_PAYOUT_DEVICE             byte = 0x5B
	CMD_ENABLE_PAYOUT_DEVICE              byte = 0x5C
	CMD_SET_FIXED_ENCRYPTION_KEY          byte = 0x60
	CMD_RESET_FIXED_ENCRYPTION_KEY        byte = 0x61
	CMD_REQUEST_TEBS_BARCODE              byte = 0x65
	CMD_REQUEST_TEBS_LOG                  byte = 0x66
	CMD_TEBS_UNLOCK_ENABLE                byte = 0x67
	CMD_TEBS_UNLOCK_DISABLE               byte = 0x68

	POLL_TEBS_CASHBOX_OUT_OF_SERVICE        byte = 0x90
	POLL_TEBS_CASHBOX_TAMPER                byte = 0x91
	POLL_TEBS_CASHBOX_IN_SERVICE            byte = 0x92
	POLL_TEBS_CASHBOX_UNLOCK_ENABLED        byte = 0x93
	POLL_JAM_RECOVERY                       byte = 0xB0
	POLL_ERROR_DURING_PAYOUT                byte = 0xB1
	POLL_SMART_EMPTYING                     byte = 0xB3
	POLL_SMART_EMPTIED                      byte = 0xB4
	POLL_CHANNEL_DISABLE                    byte = 0xB5
	POLL_INITIALISING                       byte = 0xB6
	POLL_COIN_MECH_ERROR                    byte = 0xB7
	POLL_EMPTYING                           byte = 0xC2
	POLL_EMPTIED                            byte = 0xC3
	POLL_COIN_MECH_JAMMED                   byte = 0xC4
	POLL_COIN_MECH_RETURN_PRESSED           byte = 0xC5
	POLL_PAYOUT_OUT_OF_SERVICE              byte = 0xC6
	POLL_NOTE_FLOAT_REMOVED                 byte = 0xC7
	POLL_NOTE_FLOAT_ATTACHED                byte = 0xC8
	POLL_NOTE_TRANSFERED_TO_STACKER         byte = 0xC9
	POLL_NOTE_PAID_INTO_STACKER_AT_POWER_UP byte = 0xCA
	POLL_NOTE_PAID_INTO_STORE_AT_POWER_UP   byte = 0xCB
	POLL_NOTE_STACKING                      byte = 0xCC
	POLL_NOTE_DISPENSED_AT_POWER_UP         byte = 0xCD
	POLL_NOTE_HELD_IN_BEZEL                 byte = 0xCE
	POLL_BAR_CODE_TICKET_ACKNOWLEDGE        byte = 0xD1
	POLL_DISPENSED                          byte = 0xD2
	POLL_JAMMED                             byte = 0xD5
	POLL_HALTED                             byte = 0xD6
	POLL_FLOATING                           byte = 0xD7
	POLL_FLOATED                            byte = 0xD8
	POLL_TIME_OUT                           byte = 0xD9
	POLL_DISPENSING                         byte = 0xDA
	POLL_NOTE_STORED_IN_PAYOUT              byte = 0xDB
	POLL_INCOMPLETE_PAYOUT                  byte = 0xDC
	POLL_INCOMPLETE_FLOAT                   byte = 0xDD
	POLL_CASHBOX_PAID                       byte = 0xDE
	POLL_COIN_CREDIT                        byte = 0xDF
	POLL_NOTE_PATH_OPEN                     byte = 0xE0
	POLL_NOTE_CLEARED_FROM_FRONT            byte = 0xE1
	POLL_NOTE_CLEARED_TO_CASHBOX            byte = 0xE2
	POLL_CASHBOX_REMOVED                    byte = 0xE3
	POLL_CASHBOX_REPLACED                   byte = 0xE4
	POLL_BAR_CODE_TICKET_VALIDATED          byte = 0xE5
	POLL_FRAUD_ATTEMPT                      byte = 0xE6
	POLL_STACKER_FULL                       byte = 0xE7
	POLL_DISABLED                           byte = 0xE8
	POLL_UNSAFE_NOTE_JAM                    byte = 0xE9
	POLL_SAFE_NOTE_JAM                      byte = 0xEA
	POLL_NOTE_STACKED                       byte = 0xEB
	POLL_NOTE_REJECTED                      byte = 0xEC
	POLL_NOTE_REJECTING                     byte = 0xED
	POLL_CREDIT_NOTE                        byte = 0xEE
	POLL_READ_NOTE                          byte = 0xEF
	POLL_SLAVE_RESET                        byte = 0xF1

	RESPONSE_OK                          byte = 0xF0
	RESPONSE_COMMAND_NOT_KNOWN           byte = 0xF2
	RESPONSE_WRONG_NO_PARAMETERS         byte = 0xF3
	RESPONSE_PARAMETER_OUT_OF_RANGE      byte = 0xF4
	RESPONSE_COMMAND_CANNOT_BE_PROCESSED byte = 0xF5
	RESPONSE_SOFTWARE_ERROR              byte = 0xF6
	RESPONSE_FAIL                        byte = 0xF8
	RESPONSE_KEY_NOT_SET                 byte = 0xFA
)

var RejectReasons = map[byte]string{
	0x00: "Note Accepted",
	0x01: "Note length incorrect",
	0x02: "Reject reason 2",
	0x03: "Reject reason 3",
	0x04: "Reject reason 4",
	0x05: "Reject reason 5",
	0x06: "Channel Inhibited",
	0x07: "Second Note Inserted",
	0x08: "Reject reason 8",
	0x09: "Note recognised in more than one channel",
	0x0A: "Reject reason 10",
	0x0B: "Note too long",
	0x0C: "Reject reason 12",
	0x0D: "Mechanism Slow / Stalled",
	0x0E: "Striming Attempt",
	0x0F: "Fraud Channel Reject",
	0x10: "No Notes Inserted",
	0x11: "Peak Detect Fail",
	0x12: "Twisted note detected",
	0x13: "Escrow time-out",
	0x14: "Bar code scan fail",
	0x15: "Rear sensor 2 Fail",
	0x16: "Slot Fail 1",
	0x17: "Slot Fail 2",
	0x18: "Lens Over Sample",
	0x19: "Width Detect Fail",
	0x1A: "Short Note Detected",
}

var Commands = map[byte]string{
	CMD_RESET:                             "",
	CMD_SET_CHANNEL_INHIBITS:              "",
	CMD_DISPLAY_ON:                        "",
	CMD_DISPLAY_OFF:                       "",
	CMD_SETUP_REQUEST:                     "",
	CMD_HOST_PROTOCOL_VERSION:             "",
	CMD_POLL:                              "",
	CMD_REJECT_BANKNOTE:                   "",
	CMD_DISABLE:                           "",
	CMD_ENABLE:                            "",
	CMD_GET_SERIAL_NUMBER:                 "",
	CMD_UNIT_DATA:                         "",
	CMD_CHANNEL_VALUE_REQUEST:             "",
	CMD_CHANNEL_SECURITY_DATA:             "",
	CMD_CHANNEL_RE_TEACH_DATA:             "",
	CMD_SYNC:                              "",
	CMD_LAST_REJECT_CODE:                  "",
	CMD_HOLD:                              "",
	CMD_GET_FIRMWARE_VERSION:              "",
	CMD_GET_DATASET_VERSION:               "",
	CMD_GET_ALL_LEVELS:                    "",
	CMD_GET_BAR_CODE_READER_CONFIGURATION: "",
	CMD_SET_BAR_CODE_CONFIGURATION:        "",
	CMD_GET_BAR_CODE_INHIBIT_STATUS:       "",
	CMD_SET_BAR_CODE_INHIBIT_STATUS:       "",
	CMD_GET_BAR_CODE_DATA:                 "",
	CMD_SET_REFILL_MODE:                   "",
	CMD_PAYOUT_AMOUNT:                     "",
	CMD_SET_DENOMINATION_LEVEL:            "",
	CMD_GET_DENOMINATION_LEVEL:            "",
	CMD_COMMUNICATION_PASS_THROUGH:        "",
	CMD_HALT_PAYOUT:                       "",
	CMD_SET_DENOMINATION_ROUTE:            "",
	CMD_GET_DENOMINATION_ROUTE:            "",
	CMD_FLOAT_AMOUNT:                      "",
	CMD_GET_MINIMUM_PAYOUT:                "",
	CMD_EMPTY_ALL:                         "",
	CMD_SET_COIN_MECH_INHIBITS:            "",
	CMD_GET_NOTE_POSITIONS:                "",
	CMD_PAYOUT_NOTE:                       "",
	CMD_STACK_NOTE:                        "",
	CMD_FLOAT_BY_DENOMINATION:             "",
	CMD_SET_VALUE_REPORTING_TYPE:          "",
	CMD_PAYOUT_BY_DENOMINATION:            "",
	CMD_SET_COIN_MECH_GLOBAL_INHIBIT:      "",
	CMD_SET_GENERATOR:                     "",
	CMD_SET_MODULUS:                       "",
	CMD_REQUEST_KEY_EXCHANGE:              "",
	CMD_SET_BAUD_RATE:                     "",
	CMD_GET_BUILD_REVISION:                "",
	CMD_SET_HOPPER_OPTIONS:                "",
	CMD_GET_HOPPER_OPTIONS:                "",
	CMD_SMART_EMPTY:                       "",
	CMD_CASHBOX_PAYOUT_OPERATION_DATA:     "",
	CMD_CONFIGURE_BEZEL:                   "",
	CMD_POLL_WITH_ACK:                     "",
	CMD_EVENT_ACK:                         "",
	CMD_GET_COUNTERS:                      "",
	CMD_RESET_COUNTERS:                    "",
	CMD_COIN_MECH_OPTIONS:                 "",
	CMD_DISABLE_PAYOUT_DEVICE:             "",
	CMD_ENABLE_PAYOUT_DEVICE:              "",
	CMD_SET_FIXED_ENCRYPTION_KEY:          "",
	CMD_RESET_FIXED_ENCRYPTION_KEY:        "",
	CMD_REQUEST_TEBS_BARCODE:              "",
	CMD_REQUEST_TEBS_LOG:                  "",
	CMD_TEBS_UNLOCK_ENABLE:                "",
	CMD_TEBS_UNLOCK_DISABLE:               "",
}
