package common

// common
const (
	// Version openedge version
	Version = "0.9.8.13"
	// log entry fields
	LogID                   = "id"
	LogHost                 = "host"
	LogClient               = "client"
	LogUsername             = "username"
	LogMQTTVersion          = "mqttVersion"
	LogCleanSession         = "cleanSession"
	LogApp                  = "app"
	LogComponent            = "component"
	LogMessage              = "message"
	LogPacketID             = "packetID"
	LogMessageTopic         = "messageTopic"
	LogMessagePayloadLength = "messagePayloadLength"
	LogMessageQOS           = "messageQOS"
	LogWillTopic            = "willTopic"
	// LogPubTopic             = "pubTopic"
	// LogSubTopic             = "subTopic"
	LogUnSubTopic = "unsubTopic"
	LogSinkTopic  = "sinkTopic"
	LogQueueName  = "queueName"
	LogFunction   = "function"
	LogContext    = "context"
	LogName       = "name"
	LogInvoke     = "invoke"
	LogInstance   = "instance"
	LogProcess    = "process"
	LogSequence   = "sequence"
	LogOffset     = "offset"
	LogSink       = "sink"
	LogTimestamp  = "timestamp"
	LogVersion    = "version"
	LogEvent      = "event"
	LogSource     = "source"
	LogTarget     = "target"
	LogTopic      = "topic"
	LogQOS        = "qos"

	// topic validate fields
	MaxSlashCount          = 8
	MaxTopicNameLen        = 255
	MaxPayloadLengthString = "Byte"

	// wildcard topic fields
	TopicSeparator    = "/"
	SingleWildCard    = "+"
	MultipleWildCard  = "#"
	SysCmdPrefix      = "$"
	FuncTopicPrefix   = "$function/"
	RemoteTopicPrefix = "$remote/"

	PrefixSub  = "sub."
	PrefixPub  = "pub."
	PrefixSess = "$session/"
	PrefixTmp  = "$session/tmp/"
	RuleMsgQ0  = "$rule/msgqos0"
	RuleTopic  = "$rule/topic"
)

// Queue的Bucket名字不能包含'.'，非Queue（特别是metadata）的Bucket命名推荐'.'起头
var (
	// BucketNameDotSubscription stores session's subscription
	BucketNameDotSubscription = []byte(".subscription")
	// BucketNameDotRetained stores session's retained message
	BucketNameDotRetained = []byte(".retained")
	// BucketNameDotWill stores session's will message ?
	BucketNameDotWill = []byte(".will")
	// BucketNameDotAuth stores auth data
	BucketNameDotAuth = []byte(".auth")
	// BucketNameDotMapping stores topic mappings
	BucketNameDotMapping = []byte(".mapping")
)

// Subscription
var (
	SubTopicSysRemoteFormat = "$SYS/remote/%s/%s"
	SubTopicSysFuncPrefix   = "$SYS/function/"
	SubTypeFunc             = "function"
	SubTypeRemote           = "remote"
	// SubVersionV0          = "v0"
)
