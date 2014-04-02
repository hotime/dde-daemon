// This file is automatically generated, please don't edit manully.
package main

// Get key type
func getSettingWirelessSecurityKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		t = ktypeString
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		t = ktypeUint32
	}
	return
}

// Get key's default value
func getSettingWirelessSecurityKeyDefaultValueJSON(key string) (valueJSON string) {
	value := getSettingWirelessSecurityKeyDefaultValue(key)
	t := getSettingWirelessSecurityKeyType(key)
	valueJSON, err := keyValueToJSON(value, t)
	if err != nil {
		LOGGER.Error("getSettingWirelessSecurityKeyDefaultValueJSON:", err)
	}
	return
}
func getSettingWirelessSecurityKeyDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		LOGGER.Error("invalid key:", key)
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		value = 0
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		value = make([]string, 0)
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		value = make([]string, 0)
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		value = make([]string, 0)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		value = 0
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		value = 0
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		value = 0
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		value = nil
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		value = 0
	}
	return
}

// Get JSON value generally
func generalGetSettingWirelessSecurityKeyJSON(data _ConnectionData, key string) (value string) {
	if !isConnectionDataKeyExists(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, key) {
		return getSettingWirelessSecurityKeyDefaultValueJSON(key)
	}
	switch key {
	default:
		LOGGER.Error("generalGetSettingWirelessSecurityKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SECURITY_KEY_MGMT:
		value = getSettingWirelessSecurityKeyMgmtJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX:
		value = getSettingWirelessSecurityWepTxKeyidxJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_AUTH_ALG:
		value = getSettingWirelessSecurityAuthAlgJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PROTO:
		value = getSettingWirelessSecurityProtoJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PAIRWISE:
		value = getSettingWirelessSecurityPairwiseJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_GROUP:
		value = getSettingWirelessSecurityGroupJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME:
		value = getSettingWirelessSecurityLeapUsernameJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY0:
		value = getSettingWirelessSecurityWepKey0JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY1:
		value = getSettingWirelessSecurityWepKey1JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY2:
		value = getSettingWirelessSecurityWepKey2JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY3:
		value = getSettingWirelessSecurityWepKey3JSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS:
		value = getSettingWirelessSecurityWepKeyFlagsJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE:
		value = getSettingWirelessSecurityWepKeyTypeJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PSK:
		value = getSettingWirelessSecurityPskJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS:
		value = getSettingWirelessSecurityPskFlagsJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD:
		value = getSettingWirelessSecurityLeapPasswordJSON(data)
	case NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS:
		value = getSettingWirelessSecurityLeapPasswordFlagsJSON(data)
	}
	return
}

// Getter
func getSettingWirelessSecurityKeyMgmt(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT).(string)
	return
}
func getSettingWirelessSecurityWepTxKeyidx(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX).(uint32)
	return
}
func getSettingWirelessSecurityAuthAlg(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG).(string)
	return
}
func getSettingWirelessSecurityProto(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO).([]string)
	return
}
func getSettingWirelessSecurityPairwise(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE).([]string)
	return
}
func getSettingWirelessSecurityGroup(data _ConnectionData) (value []string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP).([]string)
	return
}
func getSettingWirelessSecurityLeapUsername(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME).(string)
	return
}
func getSettingWirelessSecurityWepKey0(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0).(string)
	return
}
func getSettingWirelessSecurityWepKey1(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1).(string)
	return
}
func getSettingWirelessSecurityWepKey2(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2).(string)
	return
}
func getSettingWirelessSecurityWepKey3(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3).(string)
	return
}
func getSettingWirelessSecurityWepKeyFlags(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS).(uint32)
	return
}
func getSettingWirelessSecurityWepKeyType(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE).(uint32)
	return
}
func getSettingWirelessSecurityPsk(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK).(string)
	return
}
func getSettingWirelessSecurityPskFlags(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS).(uint32)
	return
}
func getSettingWirelessSecurityLeapPassword(data _ConnectionData) (value string) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD).(string)
	return
}
func getSettingWirelessSecurityLeapPasswordFlags(data _ConnectionData) (value uint32) {
	value, _ = getConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS).(uint32)
	return
}

// Setter
func setSettingWirelessSecurityKeyMgmt(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, value)
}
func setSettingWirelessSecurityWepTxKeyidx(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, value)
}
func setSettingWirelessSecurityAuthAlg(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, value)
}
func setSettingWirelessSecurityProto(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, value)
}
func setSettingWirelessSecurityPairwise(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, value)
}
func setSettingWirelessSecurityGroup(data _ConnectionData, value []string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, value)
}
func setSettingWirelessSecurityLeapUsername(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, value)
}
func setSettingWirelessSecurityWepKey0(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, value)
}
func setSettingWirelessSecurityWepKey1(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, value)
}
func setSettingWirelessSecurityWepKey2(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, value)
}
func setSettingWirelessSecurityWepKey3(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, value)
}
func setSettingWirelessSecurityWepKeyFlags(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, value)
}
func setSettingWirelessSecurityWepKeyType(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, value)
}
func setSettingWirelessSecurityPsk(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, value)
}
func setSettingWirelessSecurityPskFlags(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, value)
}
func setSettingWirelessSecurityLeapPassword(data _ConnectionData, value string) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, value)
}
func setSettingWirelessSecurityLeapPasswordFlags(data _ConnectionData, value uint32) {
	setConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, value)
}

// JSON Getter
func getSettingWirelessSecurityKeyMgmtJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_KEY_MGMT))
	return
}
func getSettingWirelessSecurityWepTxKeyidxJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX))
	return
}
func getSettingWirelessSecurityAuthAlgJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_AUTH_ALG))
	return
}
func getSettingWirelessSecurityProtoJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PROTO))
	return
}
func getSettingWirelessSecurityPairwiseJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PAIRWISE))
	return
}
func getSettingWirelessSecurityGroupJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_GROUP))
	return
}
func getSettingWirelessSecurityLeapUsernameJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME))
	return
}
func getSettingWirelessSecurityWepKey0JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY0))
	return
}
func getSettingWirelessSecurityWepKey1JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY1))
	return
}
func getSettingWirelessSecurityWepKey2JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY2))
	return
}
func getSettingWirelessSecurityWepKey3JSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY3))
	return
}
func getSettingWirelessSecurityWepKeyFlagsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS))
	return
}
func getSettingWirelessSecurityWepKeyTypeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE))
	return
}
func getSettingWirelessSecurityPskJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK))
	return
}
func getSettingWirelessSecurityPskFlagsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS))
	return
}
func getSettingWirelessSecurityLeapPasswordJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD))
	return
}
func getSettingWirelessSecurityLeapPasswordFlagsJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS))
	return
}

// JSON Setter
func setSettingWirelessSecurityKeyMgmtJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_KEY_MGMT))
}
func setSettingWirelessSecurityWepTxKeyidxJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX))
}
func setSettingWirelessSecurityAuthAlgJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_AUTH_ALG))
}
func setSettingWirelessSecurityProtoJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PROTO))
}
func setSettingWirelessSecurityPairwiseJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PAIRWISE))
}
func setSettingWirelessSecurityGroupJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_GROUP))
}
func setSettingWirelessSecurityLeapUsernameJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME))
}
func setSettingWirelessSecurityWepKey0JSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY0))
}
func setSettingWirelessSecurityWepKey1JSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY1))
}
func setSettingWirelessSecurityWepKey2JSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY2))
}
func setSettingWirelessSecurityWepKey3JSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY3))
}
func setSettingWirelessSecurityWepKeyFlagsJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS))
}
func setSettingWirelessSecurityWepKeyTypeJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE))
}
func setSettingWirelessSecurityPskJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK))
}
func setSettingWirelessSecurityPskFlagsJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS))
}
func setSettingWirelessSecurityLeapPasswordJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD))
}
func setSettingWirelessSecurityLeapPasswordFlagsJSON(data _ConnectionData, valueJSON string) {
	setConnectionDataKeyJSON(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS, valueJSON, getSettingWirelessSecurityKeyType(NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS))
}

// Remover
func removeSettingWirelessSecurityKeyMgmt(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT)
}
func removeSettingWirelessSecurityWepTxKeyidx(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_TX_KEYIDX)
}
func removeSettingWirelessSecurityAuthAlg(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_AUTH_ALG)
}
func removeSettingWirelessSecurityProto(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PROTO)
}
func removeSettingWirelessSecurityPairwise(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PAIRWISE)
}
func removeSettingWirelessSecurityGroup(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_GROUP)
}
func removeSettingWirelessSecurityLeapUsername(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_USERNAME)
}
func removeSettingWirelessSecurityWepKey0(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY0)
}
func removeSettingWirelessSecurityWepKey1(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY1)
}
func removeSettingWirelessSecurityWepKey2(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY2)
}
func removeSettingWirelessSecurityWepKey3(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY3)
}
func removeSettingWirelessSecurityWepKeyFlags(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_FLAGS)
}
func removeSettingWirelessSecurityWepKeyType(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_WEP_KEY_TYPE)
}
func removeSettingWirelessSecurityPsk(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK)
}
func removeSettingWirelessSecurityPskFlags(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_PSK_FLAGS)
}
func removeSettingWirelessSecurityLeapPassword(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD)
}
func removeSettingWirelessSecurityLeapPasswordFlags(data _ConnectionData) {
	removeConnectionDataKey(data, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_LEAP_PASSWORD_FLAGS)
}
