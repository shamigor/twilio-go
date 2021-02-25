/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UsageRecordTodayCategory the model 'UsageRecordTodayCategory'
type UsageRecordTodayCategory string

// List of usage_record_today_category
const (
	USAGERECORDTODAYCATEGORY_AGENT_CONFERENCE                                       UsageRecordTodayCategory = "agent-conference"
	USAGERECORDTODAYCATEGORY_ANSWERING_MACHINE_DETECTION                            UsageRecordTodayCategory = "answering-machine-detection"
	USAGERECORDTODAYCATEGORY_AUTHY_AUTHENTICATIONS                                  UsageRecordTodayCategory = "authy-authentications"
	USAGERECORDTODAYCATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageRecordTodayCategory = "authy-calls-outbound"
	USAGERECORDTODAYCATEGORY_AUTHY_MONTHLY_FEES                                     UsageRecordTodayCategory = "authy-monthly-fees"
	USAGERECORDTODAYCATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageRecordTodayCategory = "authy-phone-intelligence"
	USAGERECORDTODAYCATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageRecordTodayCategory = "authy-phone-verifications"
	USAGERECORDTODAYCATEGORY_AUTHY_SMS_OUTBOUND                                     UsageRecordTodayCategory = "authy-sms-outbound"
	USAGERECORDTODAYCATEGORY_CALL_PROGESS_EVENTS                                    UsageRecordTodayCategory = "call-progess-events"
	USAGERECORDTODAYCATEGORY_CALLERIDLOOKUPS                                        UsageRecordTodayCategory = "calleridlookups"
	USAGERECORDTODAYCATEGORY_CALLS                                                  UsageRecordTodayCategory = "calls"
	USAGERECORDTODAYCATEGORY_CALLS_CLIENT                                           UsageRecordTodayCategory = "calls-client"
	USAGERECORDTODAYCATEGORY_CALLS_GLOBALCONFERENCE                                 UsageRecordTodayCategory = "calls-globalconference"
	USAGERECORDTODAYCATEGORY_CALLS_INBOUND                                          UsageRecordTodayCategory = "calls-inbound"
	USAGERECORDTODAYCATEGORY_CALLS_INBOUND_LOCAL                                    UsageRecordTodayCategory = "calls-inbound-local"
	USAGERECORDTODAYCATEGORY_CALLS_INBOUND_MOBILE                                   UsageRecordTodayCategory = "calls-inbound-mobile"
	USAGERECORDTODAYCATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageRecordTodayCategory = "calls-inbound-tollfree"
	USAGERECORDTODAYCATEGORY_CALLS_OUTBOUND                                         UsageRecordTodayCategory = "calls-outbound"
	USAGERECORDTODAYCATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageRecordTodayCategory = "calls-pay-verb-transactions"
	USAGERECORDTODAYCATEGORY_CALLS_RECORDINGS                                       UsageRecordTodayCategory = "calls-recordings"
	USAGERECORDTODAYCATEGORY_CALLS_SIP                                              UsageRecordTodayCategory = "calls-sip"
	USAGERECORDTODAYCATEGORY_CALLS_SIP_INBOUND                                      UsageRecordTodayCategory = "calls-sip-inbound"
	USAGERECORDTODAYCATEGORY_CALLS_SIP_OUTBOUND                                     UsageRecordTodayCategory = "calls-sip-outbound"
	USAGERECORDTODAYCATEGORY_CARRIER_LOOKUPS                                        UsageRecordTodayCategory = "carrier-lookups"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS                                          UsageRecordTodayCategory = "conversations"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS_API_REQUESTS                             UsageRecordTodayCategory = "conversations-api-requests"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageRecordTodayCategory = "conversations-conversation-events"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageRecordTodayCategory = "conversations-endpoint-connectivity"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS_EVENTS                                   UsageRecordTodayCategory = "conversations-events"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageRecordTodayCategory = "conversations-participant-events"
	USAGERECORDTODAYCATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageRecordTodayCategory = "conversations-participants"
	USAGERECORDTODAYCATEGORY_CPS                                                    UsageRecordTodayCategory = "cps"
	USAGERECORDTODAYCATEGORY_FRAUD_LOOKUPS                                          UsageRecordTodayCategory = "fraud-lookups"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS                                            UsageRecordTodayCategory = "group-rooms"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageRecordTodayCategory = "group-rooms-data-track"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageRecordTodayCategory = "group-rooms-encrypted-media-recorded"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageRecordTodayCategory = "group-rooms-media-downloaded"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageRecordTodayCategory = "group-rooms-media-recorded"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageRecordTodayCategory = "group-rooms-media-routed"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageRecordTodayCategory = "group-rooms-media-stored"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageRecordTodayCategory = "group-rooms-participant-minutes"
	USAGERECORDTODAYCATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageRecordTodayCategory = "group-rooms-recorded-minutes"
	USAGERECORDTODAYCATEGORY_IMP_V1_USAGE                                           UsageRecordTodayCategory = "imp-v1-usage"
	USAGERECORDTODAYCATEGORY_LOOKUPS                                                UsageRecordTodayCategory = "lookups"
	USAGERECORDTODAYCATEGORY_MARKETPLACE                                            UsageRecordTodayCategory = "marketplace"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageRecordTodayCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageRecordTodayCategory = "marketplace-cadence-transcription"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageRecordTodayCategory = "marketplace-cadence-translation"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageRecordTodayCategory = "marketplace-capio-speech-to-text"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageRecordTodayCategory = "marketplace-convriza-ababa"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageRecordTodayCategory = "marketplace-deepgram-phrase-detector"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageRecordTodayCategory = "marketplace-digital-segment-business-info"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageRecordTodayCategory = "marketplace-facebook-offline-conversions"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageRecordTodayCategory = "marketplace-google-speech-to-text"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageRecordTodayCategory = "marketplace-ibm-watson-message-insights"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageRecordTodayCategory = "marketplace-ibm-watson-message-sentiment"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageRecordTodayCategory = "marketplace-ibm-watson-recording-analysis"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageRecordTodayCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageRecordTodayCategory = "marketplace-icehook-systems-scout"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageRecordTodayCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageRecordTodayCategory = "marketplace-keen-io-contact-center-analytics"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageRecordTodayCategory = "marketplace-marchex-cleancall"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageRecordTodayCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageRecordTodayCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageRecordTodayCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageRecordTodayCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageRecordTodayCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageRecordTodayCategory = "marketplace-nomorobo-spam-score"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageRecordTodayCategory = "marketplace-payfone-tcpa-compliance"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageRecordTodayCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageRecordTodayCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageRecordTodayCategory = "marketplace-telo-opencnam"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageRecordTodayCategory = "marketplace-truecnam-true-spam"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageRecordTodayCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageRecordTodayCategory = "marketplace-twilio-carrier-information-lookup"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageRecordTodayCategory = "marketplace-voicebase-pci"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageRecordTodayCategory = "marketplace-voicebase-transcription"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageRecordTodayCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageRecordTodayCategory = "marketplace-whitepages-pro-caller-identification"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageRecordTodayCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageRecordTodayCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageRecordTodayCategory = "marketplace-wolfarm-spoken-results"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageRecordTodayCategory = "marketplace-wolfram-short-answer"
	USAGERECORDTODAYCATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageRecordTodayCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGERECORDTODAYCATEGORY_MEDIASTORAGE                                           UsageRecordTodayCategory = "mediastorage"
	USAGERECORDTODAYCATEGORY_MMS                                                    UsageRecordTodayCategory = "mms"
	USAGERECORDTODAYCATEGORY_MMS_INBOUND                                            UsageRecordTodayCategory = "mms-inbound"
	USAGERECORDTODAYCATEGORY_MMS_INBOUND_LONGCODE                                   UsageRecordTodayCategory = "mms-inbound-longcode"
	USAGERECORDTODAYCATEGORY_MMS_INBOUND_SHORTCODE                                  UsageRecordTodayCategory = "mms-inbound-shortcode"
	USAGERECORDTODAYCATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageRecordTodayCategory = "mms-messages-carrierfees"
	USAGERECORDTODAYCATEGORY_MMS_OUTBOUND                                           UsageRecordTodayCategory = "mms-outbound"
	USAGERECORDTODAYCATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageRecordTodayCategory = "mms-outbound-longcode"
	USAGERECORDTODAYCATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageRecordTodayCategory = "mms-outbound-shortcode"
	USAGERECORDTODAYCATEGORY_MONITOR_READS                                          UsageRecordTodayCategory = "monitor-reads"
	USAGERECORDTODAYCATEGORY_MONITOR_STORAGE                                        UsageRecordTodayCategory = "monitor-storage"
	USAGERECORDTODAYCATEGORY_MONITOR_WRITES                                         UsageRecordTodayCategory = "monitor-writes"
	USAGERECORDTODAYCATEGORY_NOTIFY                                                 UsageRecordTodayCategory = "notify"
	USAGERECORDTODAYCATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageRecordTodayCategory = "notify-actions-attempts"
	USAGERECORDTODAYCATEGORY_NOTIFY_CHANNELS                                        UsageRecordTodayCategory = "notify-channels"
	USAGERECORDTODAYCATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageRecordTodayCategory = "number-format-lookups"
	USAGERECORDTODAYCATEGORY_PCHAT                                                  UsageRecordTodayCategory = "pchat"
	USAGERECORDTODAYCATEGORY_PCHAT_USERS                                            UsageRecordTodayCategory = "pchat-users"
	USAGERECORDTODAYCATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageRecordTodayCategory = "peer-to-peer-rooms-participant-minutes"
	USAGERECORDTODAYCATEGORY_PFAX                                                   UsageRecordTodayCategory = "pfax"
	USAGERECORDTODAYCATEGORY_PFAX_MINUTES                                           UsageRecordTodayCategory = "pfax-minutes"
	USAGERECORDTODAYCATEGORY_PFAX_MINUTES_INBOUND                                   UsageRecordTodayCategory = "pfax-minutes-inbound"
	USAGERECORDTODAYCATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageRecordTodayCategory = "pfax-minutes-outbound"
	USAGERECORDTODAYCATEGORY_PFAX_PAGES                                             UsageRecordTodayCategory = "pfax-pages"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS                                           UsageRecordTodayCategory = "phonenumbers"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS_CPS                                       UsageRecordTodayCategory = "phonenumbers-cps"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS_EMERGENCY                                 UsageRecordTodayCategory = "phonenumbers-emergency"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS_LOCAL                                     UsageRecordTodayCategory = "phonenumbers-local"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS_MOBILE                                    UsageRecordTodayCategory = "phonenumbers-mobile"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS_SETUPS                                    UsageRecordTodayCategory = "phonenumbers-setups"
	USAGERECORDTODAYCATEGORY_PHONENUMBERS_TOLLFREE                                  UsageRecordTodayCategory = "phonenumbers-tollfree"
	USAGERECORDTODAYCATEGORY_PREMIUMSUPPORT                                         UsageRecordTodayCategory = "premiumsupport"
	USAGERECORDTODAYCATEGORY_PROXY                                                  UsageRecordTodayCategory = "proxy"
	USAGERECORDTODAYCATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageRecordTodayCategory = "proxy-active-sessions"
	USAGERECORDTODAYCATEGORY_PSTNCONNECTIVITY                                       UsageRecordTodayCategory = "pstnconnectivity"
	USAGERECORDTODAYCATEGORY_PV                                                     UsageRecordTodayCategory = "pv"
	USAGERECORDTODAYCATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageRecordTodayCategory = "pv-composition-media-downloaded"
	USAGERECORDTODAYCATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageRecordTodayCategory = "pv-composition-media-encrypted"
	USAGERECORDTODAYCATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageRecordTodayCategory = "pv-composition-media-stored"
	USAGERECORDTODAYCATEGORY_PV_COMPOSITION_MINUTES                                 UsageRecordTodayCategory = "pv-composition-minutes"
	USAGERECORDTODAYCATEGORY_PV_RECORDING_COMPOSITIONS                              UsageRecordTodayCategory = "pv-recording-compositions"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS                                   UsageRecordTodayCategory = "pv-room-participants"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageRecordTodayCategory = "pv-room-participants-au1"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageRecordTodayCategory = "pv-room-participants-br1"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageRecordTodayCategory = "pv-room-participants-ie1"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageRecordTodayCategory = "pv-room-participants-jp1"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageRecordTodayCategory = "pv-room-participants-sg1"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageRecordTodayCategory = "pv-room-participants-us1"
	USAGERECORDTODAYCATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageRecordTodayCategory = "pv-room-participants-us2"
	USAGERECORDTODAYCATEGORY_PV_ROOMS                                               UsageRecordTodayCategory = "pv-rooms"
	USAGERECORDTODAYCATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageRecordTodayCategory = "pv-sip-endpoint-registrations"
	USAGERECORDTODAYCATEGORY_RECORDINGS                                             UsageRecordTodayCategory = "recordings"
	USAGERECORDTODAYCATEGORY_RECORDINGSTORAGE                                       UsageRecordTodayCategory = "recordingstorage"
	USAGERECORDTODAYCATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageRecordTodayCategory = "rooms-group-bandwidth"
	USAGERECORDTODAYCATEGORY_ROOMS_GROUP_MINUTES                                    UsageRecordTodayCategory = "rooms-group-minutes"
	USAGERECORDTODAYCATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageRecordTodayCategory = "rooms-peer-to-peer-minutes"
	USAGERECORDTODAYCATEGORY_SHORTCODES                                             UsageRecordTodayCategory = "shortcodes"
	USAGERECORDTODAYCATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageRecordTodayCategory = "shortcodes-customerowned"
	USAGERECORDTODAYCATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageRecordTodayCategory = "shortcodes-mms-enablement"
	USAGERECORDTODAYCATEGORY_SHORTCODES_MPS                                         UsageRecordTodayCategory = "shortcodes-mps"
	USAGERECORDTODAYCATEGORY_SHORTCODES_RANDOM                                      UsageRecordTodayCategory = "shortcodes-random"
	USAGERECORDTODAYCATEGORY_SHORTCODES_UK                                          UsageRecordTodayCategory = "shortcodes-uk"
	USAGERECORDTODAYCATEGORY_SHORTCODES_VANITY                                      UsageRecordTodayCategory = "shortcodes-vanity"
	USAGERECORDTODAYCATEGORY_SMALL_GROUP_ROOMS                                      UsageRecordTodayCategory = "small-group-rooms"
	USAGERECORDTODAYCATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageRecordTodayCategory = "small-group-rooms-data-track"
	USAGERECORDTODAYCATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageRecordTodayCategory = "small-group-rooms-participant-minutes"
	USAGERECORDTODAYCATEGORY_SMS                                                    UsageRecordTodayCategory = "sms"
	USAGERECORDTODAYCATEGORY_SMS_INBOUND                                            UsageRecordTodayCategory = "sms-inbound"
	USAGERECORDTODAYCATEGORY_SMS_INBOUND_LONGCODE                                   UsageRecordTodayCategory = "sms-inbound-longcode"
	USAGERECORDTODAYCATEGORY_SMS_INBOUND_SHORTCODE                                  UsageRecordTodayCategory = "sms-inbound-shortcode"
	USAGERECORDTODAYCATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageRecordTodayCategory = "sms-messages-carrierfees"
	USAGERECORDTODAYCATEGORY_SMS_MESSAGES_FEATURES                                  UsageRecordTodayCategory = "sms-messages-features"
	USAGERECORDTODAYCATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageRecordTodayCategory = "sms-messages-features-senderid"
	USAGERECORDTODAYCATEGORY_SMS_OUTBOUND                                           UsageRecordTodayCategory = "sms-outbound"
	USAGERECORDTODAYCATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageRecordTodayCategory = "sms-outbound-content-inspection"
	USAGERECORDTODAYCATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageRecordTodayCategory = "sms-outbound-longcode"
	USAGERECORDTODAYCATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageRecordTodayCategory = "sms-outbound-shortcode"
	USAGERECORDTODAYCATEGORY_SPEECH_RECOGNITION                                     UsageRecordTodayCategory = "speech-recognition"
	USAGERECORDTODAYCATEGORY_STUDIO_ENGAGEMENTS                                     UsageRecordTodayCategory = "studio-engagements"
	USAGERECORDTODAYCATEGORY_SYNC                                                   UsageRecordTodayCategory = "sync"
	USAGERECORDTODAYCATEGORY_SYNC_ACTIONS                                           UsageRecordTodayCategory = "sync-actions"
	USAGERECORDTODAYCATEGORY_SYNC_ENDPOINT_HOURS                                    UsageRecordTodayCategory = "sync-endpoint-hours"
	USAGERECORDTODAYCATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageRecordTodayCategory = "sync-endpoint-hours-above-daily-cap"
	USAGERECORDTODAYCATEGORY_TASKROUTER_TASKS                                       UsageRecordTodayCategory = "taskrouter-tasks"
	USAGERECORDTODAYCATEGORY_TOTALPRICE                                             UsageRecordTodayCategory = "totalprice"
	USAGERECORDTODAYCATEGORY_TRANSCRIPTIONS                                         UsageRecordTodayCategory = "transcriptions"
	USAGERECORDTODAYCATEGORY_TRUNKING_CPS                                           UsageRecordTodayCategory = "trunking-cps"
	USAGERECORDTODAYCATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageRecordTodayCategory = "trunking-emergency-calls"
	USAGERECORDTODAYCATEGORY_TRUNKING_ORIGINATION                                   UsageRecordTodayCategory = "trunking-origination"
	USAGERECORDTODAYCATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageRecordTodayCategory = "trunking-origination-local"
	USAGERECORDTODAYCATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageRecordTodayCategory = "trunking-origination-mobile"
	USAGERECORDTODAYCATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageRecordTodayCategory = "trunking-origination-tollfree"
	USAGERECORDTODAYCATEGORY_TRUNKING_RECORDINGS                                    UsageRecordTodayCategory = "trunking-recordings"
	USAGERECORDTODAYCATEGORY_TRUNKING_SECURE                                        UsageRecordTodayCategory = "trunking-secure"
	USAGERECORDTODAYCATEGORY_TRUNKING_TERMINATION                                   UsageRecordTodayCategory = "trunking-termination"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES                                          UsageRecordTodayCategory = "turnmegabytes"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageRecordTodayCategory = "turnmegabytes-australia"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_BRASIL                                   UsageRecordTodayCategory = "turnmegabytes-brasil"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_GERMANY                                  UsageRecordTodayCategory = "turnmegabytes-germany"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_INDIA                                    UsageRecordTodayCategory = "turnmegabytes-india"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_IRELAND                                  UsageRecordTodayCategory = "turnmegabytes-ireland"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_JAPAN                                    UsageRecordTodayCategory = "turnmegabytes-japan"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_SINGAPORE                                UsageRecordTodayCategory = "turnmegabytes-singapore"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_USEAST                                   UsageRecordTodayCategory = "turnmegabytes-useast"
	USAGERECORDTODAYCATEGORY_TURNMEGABYTES_USWEST                                   UsageRecordTodayCategory = "turnmegabytes-uswest"
	USAGERECORDTODAYCATEGORY_TWILIO_INTERCONNECT                                    UsageRecordTodayCategory = "twilio-interconnect"
	USAGERECORDTODAYCATEGORY_VERIFY_PUSH                                            UsageRecordTodayCategory = "verify-push"
	USAGERECORDTODAYCATEGORY_VIDEO_RECORDINGS                                       UsageRecordTodayCategory = "video-recordings"
	USAGERECORDTODAYCATEGORY_VOICE_INSIGHTS                                         UsageRecordTodayCategory = "voice-insights"
	USAGERECORDTODAYCATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageRecordTodayCategory = "voice-insights-client-insights-on-demand-minute"
	USAGERECORDTODAYCATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageRecordTodayCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGERECORDTODAYCATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageRecordTodayCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGERECORDTODAYCATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageRecordTodayCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGERECORDTODAYCATEGORY_WIRELESS                                               UsageRecordTodayCategory = "wireless"
	USAGERECORDTODAYCATEGORY_WIRELESS_ORDERS                                        UsageRecordTodayCategory = "wireless-orders"
	USAGERECORDTODAYCATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageRecordTodayCategory = "wireless-orders-artwork"
	USAGERECORDTODAYCATEGORY_WIRELESS_ORDERS_BULK                                   UsageRecordTodayCategory = "wireless-orders-bulk"
	USAGERECORDTODAYCATEGORY_WIRELESS_ORDERS_ESIM                                   UsageRecordTodayCategory = "wireless-orders-esim"
	USAGERECORDTODAYCATEGORY_WIRELESS_ORDERS_STARTER                                UsageRecordTodayCategory = "wireless-orders-starter"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE                                         UsageRecordTodayCategory = "wireless-usage"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS                                UsageRecordTodayCategory = "wireless-usage-commands"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageRecordTodayCategory = "wireless-usage-commands-africa"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageRecordTodayCategory = "wireless-usage-commands-asia"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageRecordTodayCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageRecordTodayCategory = "wireless-usage-commands-europe"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageRecordTodayCategory = "wireless-usage-commands-home"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageRecordTodayCategory = "wireless-usage-commands-northamerica"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageRecordTodayCategory = "wireless-usage-commands-oceania"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageRecordTodayCategory = "wireless-usage-commands-roaming"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA                                    UsageRecordTodayCategory = "wireless-usage-data"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageRecordTodayCategory = "wireless-usage-data-africa"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageRecordTodayCategory = "wireless-usage-data-asia"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageRecordTodayCategory = "wireless-usage-data-centralandsouthamerica"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageRecordTodayCategory = "wireless-usage-data-custom-additionalmb"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageRecordTodayCategory = "wireless-usage-data-custom-first5mb"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageRecordTodayCategory = "wireless-usage-data-domestic-roaming"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageRecordTodayCategory = "wireless-usage-data-europe"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageRecordTodayCategory = "wireless-usage-data-individual-additionalgb"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageRecordTodayCategory = "wireless-usage-data-individual-firstgb"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageRecordTodayCategory = "wireless-usage-data-international-roaming-canada"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageRecordTodayCategory = "wireless-usage-data-international-roaming-india"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageRecordTodayCategory = "wireless-usage-data-international-roaming-mexico"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageRecordTodayCategory = "wireless-usage-data-northamerica"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageRecordTodayCategory = "wireless-usage-data-oceania"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageRecordTodayCategory = "wireless-usage-data-pooled"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageRecordTodayCategory = "wireless-usage-data-pooled-downlink"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageRecordTodayCategory = "wireless-usage-data-pooled-uplink"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_MRC                                     UsageRecordTodayCategory = "wireless-usage-mrc"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageRecordTodayCategory = "wireless-usage-mrc-custom"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageRecordTodayCategory = "wireless-usage-mrc-individual"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageRecordTodayCategory = "wireless-usage-mrc-pooled"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageRecordTodayCategory = "wireless-usage-mrc-suspended"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_SMS                                     UsageRecordTodayCategory = "wireless-usage-sms"
	USAGERECORDTODAYCATEGORY_WIRELESS_USAGE_VOICE                                   UsageRecordTodayCategory = "wireless-usage-voice"
)