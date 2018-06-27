package rules

var PCC_RULES_FILE_PATH = "config/pcc_rules.cfg"
var SDF_RULES_FILE_PATH = "config/sdf_rules.cfg"
var ADC_RULES_FILE_PATH = "config/adc_rules.cfg"

var STATIC_ARP_FILE_PATH = "config/static_arp.cfg"

var SDF_UL_ACL_FILE_PATH = "/tmp/sdf_ul_acl.cfg"
var SDF_DL_ACL_FILE_PATH = "/tmp/sdf_dl_acl.cfg"
var ADC_ACL_FILE_PATH = "/tmp/adc_acl.cfg"

var GLOBAL = "GLOBAL"

// constants for pcc filter
var NUM_PCC_FILTERS = "NUM_PCC_FILTERS"
var STR_PCC_FILTER = "PCC_FILTER"

var UL_AMBR_MTR_PROFILE_IDX = "UL_AMBR_MTR_PROFILE_IDX"
var DL_AMBR_MTR_PROFILE_IDX = "DL_AMBR_MTR_PROFILE_IDX"
var REPORT_LEVEL = "REPORT_LEVEL"
var PRECEDENCE = "PRECEDENCE"
var RULE_NAME = "RULE_NAME"
var RATING_GROUP = "RATING_GROUP"
var SERVICE_ID = "SERVICE_ID"
var RULE_STATUS = "RULE_STATUS"
var CHARGING_MODE = "CHARGING_MODE"
var GATE_STATUS = "GATE_STATUS"
var METERING_METHOD = "METERING_METHOD"
var MUTE_NOTIFY = "MUTE_NOTIFY"
var MONITORING_KEY = "MONITORING_KEY"
var SPONSOR_ID = "SPONSOR_ID"
var REDIRECT_INFO = "REDIRECT_INFO"
var SESSION_CONT = "SESSION_CONT"
var DROP_PKT_COUNT = "DROP_PKT_COUNT"
var SDF_FILTER_IDX = "SDF_FILTER_IDX"
var ADC_FILTER_IDX = "ADC_FILTER_IDX"

//constants for sdf filter
var NUM_SDF_FILTERS = "NUM_SDF_FILTERS"
var STR_SDF_FILTER = "SDF_FILTER"

var DIRECTION = "DIRECTION"
var IPV4_LOCAL = "IPV4_LOCAL"
var IPV4_REMOTE = "IPV4_REMOTE"
var IPV4_LOCAL_MASK = "IPV4_LOCAL_MASK"
var IPV4_REMOTE_MASK = "IPV4_REMOTE_MASK"
var PROTOCOL = "PROTOCOL"
var LOCAL_LOW_LIMIT_PORT = "LOCAL_LOW_LIMIT_PORT"
var LOCAL_HIGH_LIMIT_PORT = "LOCAL_HIGH_LIMIT_PORT"
var REMOTE_LOW_LIMIT_PORT = "REMOTE_LOW_LIMIT_PORT"
var REMOTE_HIGH_LIMIT_PORT = "REMOTE_HIGH_LIMIT_PORT"

//----
var DIRECTION_UL = "uplink_only"
var DIRECTION_DL = "downlink_only"
var DIRECTION_BI = "bidirectional"

//constants for adc filter
var NUM_ADC_RULES = "NUM_ADC_RULES"
var ADC_RULE = "ADC_RULE"

var ADC_TYPE = "ADC_TYPE"
var IP = "IP"
var PREFIX = "PREFIX"
var DOMAIN = "DOMAIN"

//static apr entries
var SECTION_S1U = "s1u"
var SECTION_SGI = "sgi"
