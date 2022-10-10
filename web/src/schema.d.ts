// Code generated by devtools/gqltsgen DO NOT EDIT.

export interface Query {
  phoneNumberInfo?: null | PhoneNumberInfo
  debugMessages: DebugMessage[]
  user?: null | User
  users: UserConnection
  alert?: null | Alert
  alerts: AlertConnection
  service?: null | Service
  integrationKey?: null | IntegrationKey
  heartbeatMonitor?: null | HeartbeatMonitor
  services: ServiceConnection
  rotation?: null | Rotation
  rotations: RotationConnection
  calcRotationHandoffTimes: ISOTimestamp[]
  schedule?: null | Schedule
  userCalendarSubscription?: null | UserCalendarSubscription
  schedules: ScheduleConnection
  escalationPolicy?: null | EscalationPolicy
  escalationPolicies: EscalationPolicyConnection
  authSubjectsForProvider: AuthSubjectConnection
  timeZones: TimeZoneConnection
  labels: LabelConnection
  labelKeys: StringConnection
  labelValues: StringConnection
  integrationKeys: IntegrationKeyConnection
  userOverrides: UserOverrideConnection
  userOverride?: null | UserOverride
  config: ConfigValue[]
  configHints: ConfigHint[]
  systemLimits: SystemLimit[]
  debugMessageStatus: DebugMessageStatusInfo
  userContactMethod?: null | UserContactMethod
  slackChannels: SlackChannelConnection
  slackChannel?: null | SlackChannel
  generateSlackAppManifest: string
  linkAccountInfo?: null | LinkAccountInfo
}

export interface LinkAccountInfo {
  userDetails: string
  alertID?: null | number
  alertNewStatus?: null | AlertStatus
}

export interface AlertMetricsOptions {
  rInterval: ISORInterval
  filterByServiceID?: null | string[]
}

export interface AlertDataPoint {
  timestamp: ISOTimestamp
  alertCount: number
}

export interface DebugMessagesInput {
  first?: null | number
  createdBefore?: null | ISOTimestamp
  createdAfter?: null | ISOTimestamp
}

export interface DebugMessage {
  id: string
  createdAt: ISOTimestamp
  updatedAt: ISOTimestamp
  type: string
  status: string
  userID?: null | string
  userName?: null | string
  source?: null | string
  destination: string
  serviceID?: null | string
  serviceName?: null | string
  alertID?: null | number
  providerID?: null | string
}

export interface SlackChannelSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
}

export interface SlackChannel {
  id: string
  name: string
  teamID: string
}

export interface SlackChannelConnection {
  nodes: SlackChannel[]
  pageInfo: PageInfo
}

export interface SystemLimit {
  id: SystemLimitID
  description: string
  value: number
}

export interface SystemLimitInput {
  id: SystemLimitID
  value: number
}

export interface ConfigValue {
  id: string
  description: string
  value: string
  type: ConfigType
  password: boolean
  deprecated: string
}

export interface ConfigHint {
  id: string
  value: string
}

export type ConfigType = 'string' | 'stringList' | 'integer' | 'boolean'

export type SystemLimitID =
  | 'CalendarSubscriptionsPerUser'
  | 'NotificationRulesPerUser'
  | 'ContactMethodsPerUser'
  | 'EPStepsPerPolicy'
  | 'EPActionsPerStep'
  | 'ParticipantsPerRotation'
  | 'RulesPerSchedule'
  | 'IntegrationKeysPerService'
  | 'UnackedAlertsPerService'
  | 'TargetsPerSchedule'
  | 'HeartbeatMonitorsPerService'
  | 'UserOverridesPerSchedule'
  | 'MaxSMSPer15Minutes'
  | 'MaxSMSPerHour'
  | 'MaxSMSPer3Hours'
  | 'MaxVoicePer15Minutes'
  | 'MaxVoicePerHour'
  | 'MaxVoicePer3Hours'
  | 'MaxAllForAlertStatusPer3Minutes'
  | 'MaxAllForAlertStatusPer2Hours'
  | 'MaxAllForAlertStatusPer20Minutes'

export interface UserOverrideSearchOptions {
  first?: null | number
  after?: null | string
  omit?: null | string[]
  scheduleID?: null | string
  filterAddUserID?: null | string[]
  filterRemoveUserID?: null | string[]
  filterAnyUserID?: null | string[]
  start?: null | ISOTimestamp
  end?: null | ISOTimestamp
}

export interface UserOverrideConnection {
  nodes: UserOverride[]
  pageInfo: PageInfo
}

export interface IntegrationKeyConnection {
  nodes: IntegrationKey[]
  pageInfo: PageInfo
}

export interface UserOverride {
  id: string
  start: ISOTimestamp
  end: ISOTimestamp
  addUserID: string
  removeUserID: string
  addUser?: null | User
  removeUser?: null | User
  target: Target
}

export interface LabelSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  uniqueKeys?: null | boolean
  omit?: null | string[]
}

export interface LabelKeySearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
}

export interface LabelValueSearchOptions {
  key: string
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
}

export interface IntegrationKeySearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
}

export interface LabelConnection {
  nodes: Label[]
  pageInfo: PageInfo
}

export interface StringConnection {
  nodes: string[]
  pageInfo: PageInfo
}

export interface PhoneNumberInfo {
  id: string
  countryCode: string
  regionCode: string
  formatted: string
  valid: boolean
  error: string
}

export interface DebugCarrierInfo {
  name: string
  type: string
  mobileNetworkCode: string
  mobileCountryCode: string
}

export interface DebugCarrierInfoInput {
  number: string
}

export interface DebugSendSMSInput {
  from: string
  to: string
  body: string
}

export interface DebugSendSMSInfo {
  id: string
  providerURL: string
  fromNumber: string
}

export interface DebugMessageStatusInput {
  providerMessageID: string
}

export interface DebugMessageStatusInfo {
  state: NotificationState
}

export interface TemporarySchedule {
  start: ISOTimestamp
  end: ISOTimestamp
  shifts: OnCallShift[]
}

export interface ClearTemporarySchedulesInput {
  scheduleID: string
  start: ISOTimestamp
  end: ISOTimestamp
}

export interface SetTemporaryScheduleInput {
  scheduleID: string
  clearStart?: null | ISOTimestamp
  clearEnd?: null | ISOTimestamp
  start: ISOTimestamp
  end: ISOTimestamp
  shifts: SetScheduleShiftInput[]
}

export interface SetScheduleShiftInput {
  userID: string
  start: ISOTimestamp
  end: ISOTimestamp
}

export interface Mutation {
  linkAccount: boolean
  setTemporarySchedule: boolean
  clearTemporarySchedules: boolean
  setScheduleOnCallNotificationRules: boolean
  debugCarrierInfo: DebugCarrierInfo
  debugSendSMS?: null | DebugSendSMSInfo
  addAuthSubject: boolean
  deleteAuthSubject: boolean
  endAllAuthSessionsByCurrentUser: boolean
  updateUser: boolean
  testContactMethod: boolean
  updateAlerts?: null | Alert[]
  updateRotation: boolean
  escalateAlerts?: null | Alert[]
  setFavorite: boolean
  updateService: boolean
  updateEscalationPolicy: boolean
  updateEscalationPolicyStep: boolean
  deleteAll: boolean
  createAlert?: null | Alert
  createService?: null | Service
  createEscalationPolicy?: null | EscalationPolicy
  createEscalationPolicyStep?: null | EscalationPolicyStep
  createRotation?: null | Rotation
  createIntegrationKey?: null | IntegrationKey
  createHeartbeatMonitor?: null | HeartbeatMonitor
  setLabel: boolean
  createSchedule?: null | Schedule
  createUser?: null | User
  createUserCalendarSubscription: UserCalendarSubscription
  updateUserCalendarSubscription: boolean
  updateScheduleTarget: boolean
  createUserOverride?: null | UserOverride
  createUserContactMethod?: null | UserContactMethod
  createUserNotificationRule?: null | UserNotificationRule
  updateUserContactMethod: boolean
  sendContactMethodVerification: boolean
  verifyContactMethod: boolean
  updateSchedule: boolean
  updateUserOverride: boolean
  updateHeartbeatMonitor: boolean
  updateAlertsByService: boolean
  setConfig: boolean
  setSystemLimits: boolean
}

export interface UpdateAlertsByServiceInput {
  serviceID: string
  newStatus: AlertStatus
}

export interface CreateAlertInput {
  summary: string
  details?: null | string
  serviceID: string
  sanitize?: null | boolean
}

export interface CreateUserInput {
  username: string
  password: string
  name?: null | string
  email?: null | string
  role?: null | UserRole
  favorite?: null | boolean
}

export interface CreateUserCalendarSubscriptionInput {
  name: string
  reminderMinutes?: null | number[]
  scheduleID: string
  disabled?: null | boolean
}

export interface UpdateUserCalendarSubscriptionInput {
  id: string
  name?: null | string
  reminderMinutes?: null | number[]
  disabled?: null | boolean
}

export interface UserCalendarSubscription {
  id: string
  name: string
  reminderMinutes: number[]
  scheduleID: string
  schedule?: null | Schedule
  lastAccess: ISOTimestamp
  disabled: boolean
  url?: null | string
}

export interface ConfigValueInput {
  id: string
  value: string
}

export interface UpdateUserOverrideInput {
  id: string
  start?: null | ISOTimestamp
  end?: null | ISOTimestamp
  addUserID?: null | string
  removeUserID?: null | string
}

export interface CreateUserOverrideInput {
  scheduleID?: null | string
  start: ISOTimestamp
  end: ISOTimestamp
  addUserID?: null | string
  removeUserID?: null | string
}

export interface CreateScheduleInput {
  name: string
  description?: null | string
  timeZone: string
  favorite?: null | boolean
  targets?: null | ScheduleTargetInput[]
  newUserOverrides?: null | CreateUserOverrideInput[]
}

export interface ScheduleTargetInput {
  scheduleID?: null | string
  target?: null | TargetInput
  newRotation?: null | CreateRotationInput
  rules: ScheduleRuleInput[]
}

export interface ScheduleRuleInput {
  id?: null | string
  start?: null | ClockTime
  end?: null | ClockTime
  weekdayFilter?: null | WeekdayFilter
}

export interface SetLabelInput {
  target?: null | TargetInput
  key: string
  value: string
}

export interface TimeZoneSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
}

export interface TimeZoneConnection {
  nodes: TimeZone[]
  pageInfo: PageInfo
}

export interface TimeZone {
  id: string
}

export interface CreateServiceInput {
  name: string
  description?: null | string
  favorite?: null | boolean
  escalationPolicyID?: null | string
  newEscalationPolicy?: null | CreateEscalationPolicyInput
  newIntegrationKeys?: null | CreateIntegrationKeyInput[]
  labels?: null | SetLabelInput[]
  newHeartbeatMonitors?: null | CreateHeartbeatMonitorInput[]
}

export interface CreateEscalationPolicyInput {
  name: string
  description?: null | string
  repeat?: null | number
  favorite?: null | boolean
  steps?: null | CreateEscalationPolicyStepInput[]
}

export interface CreateEscalationPolicyStepInput {
  escalationPolicyID?: null | string
  delayMinutes: number
  targets?: null | TargetInput[]
  newRotation?: null | CreateRotationInput
  newSchedule?: null | CreateScheduleInput
}

export interface EscalationPolicyStep {
  id: string
  stepNumber: number
  delayMinutes: number
  targets: Target[]
  escalationPolicy?: null | EscalationPolicy
}

export interface UpdateScheduleInput {
  id: string
  name?: null | string
  description?: null | string
  timeZone?: null | string
}

export interface UpdateServiceInput {
  id: string
  name?: null | string
  description?: null | string
  escalationPolicyID?: null | string
  maintenanceExpiresAt?: null | ISOTimestamp
}

export interface UpdateEscalationPolicyInput {
  id: string
  name?: null | string
  description?: null | string
  repeat?: null | number
  stepIDs?: null | string[]
}

export interface UpdateEscalationPolicyStepInput {
  id: string
  delayMinutes?: null | number
  targets?: null | TargetInput[]
}

export interface SetFavoriteInput {
  target: TargetInput
  favorite: boolean
}

export interface EscalationPolicyConnection {
  nodes: EscalationPolicy[]
  pageInfo: PageInfo
}

export interface AlertConnection {
  nodes: Alert[]
  pageInfo: PageInfo
}

export interface ScheduleConnection {
  nodes: Schedule[]
  pageInfo: PageInfo
}

export interface Schedule {
  id: string
  name: string
  description: string
  timeZone: string
  assignedTo: Target[]
  shifts: OnCallShift[]
  targets: ScheduleTarget[]
  target?: null | ScheduleTarget
  isFavorite: boolean
  temporarySchedules: TemporarySchedule[]
  onCallNotificationRules: OnCallNotificationRule[]
}

export interface SetScheduleOnCallNotificationRulesInput {
  scheduleID: string
  rules: OnCallNotificationRuleInput[]
}

export interface OnCallNotificationRuleInput {
  id?: null | string
  target: TargetInput
  time?: null | ClockTime
  weekdayFilter?: null | WeekdayFilter
}

export interface OnCallNotificationRule {
  id: string
  target: Target
  time?: null | ClockTime
  weekdayFilter?: null | WeekdayFilter
}

export interface OnCallShift {
  userID: string
  user?: null | User
  start: ISOTimestamp
  end: ISOTimestamp
  truncated: boolean
}

export interface ScheduleTarget {
  scheduleID: string
  target: Target
  rules: ScheduleRule[]
}

export interface ScheduleRule {
  id: string
  scheduleID: string
  start: ClockTime
  end: ClockTime
  weekdayFilter: WeekdayFilter
  target: Target
}

export interface RotationConnection {
  nodes: Rotation[]
  pageInfo: PageInfo
}

export interface CreateRotationInput {
  name: string
  description?: null | string
  timeZone: string
  start: ISOTimestamp
  favorite?: null | boolean
  type: RotationType
  shiftLength?: null | number
  userIDs?: null | string[]
}

export interface Rotation {
  id: string
  name: string
  description: string
  isFavorite: boolean
  start: ISOTimestamp
  timeZone: string
  type: RotationType
  shiftLength: number
  activeUserIndex: number
  userIDs: string[]
  users: User[]
  nextHandoffTimes: ISOTimestamp[]
}

export type RotationType = 'weekly' | 'daily' | 'hourly'

export interface UpdateAlertsInput {
  alertIDs: number[]
  newStatus: AlertStatus
}

export interface UpdateRotationInput {
  id: string
  name?: null | string
  description?: null | string
  timeZone?: null | string
  start?: null | ISOTimestamp
  type?: null | RotationType
  shiftLength?: null | number
  activeUserIndex?: null | number
  userIDs?: null | string[]
}

export interface RotationSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
  favoritesOnly?: null | boolean
  favoritesFirst?: null | boolean
}

export interface CalcRotationHandoffTimesInput {
  handoff: ISOTimestamp
  from?: null | ISOTimestamp
  timeZone: string
  shiftLengthHours: number
  count: number
}

export interface EscalationPolicySearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
  favoritesOnly?: null | boolean
  favoritesFirst?: null | boolean
}

export interface ScheduleSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
  favoritesOnly?: null | boolean
  favoritesFirst?: null | boolean
}

export interface ServiceSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
  favoritesOnly?: null | boolean
  favoritesFirst?: null | boolean
}

export interface UserSearchOptions {
  first?: null | number
  after?: null | string
  search?: null | string
  omit?: null | string[]
  CMValue?: null | string
  CMType?: null | ContactMethodType
  favoritesOnly?: null | boolean
  favoritesFirst?: null | boolean
}

export interface AlertSearchOptions {
  filterByStatus?: null | AlertStatus[]
  filterByServiceID?: null | string[]
  search?: null | string
  first?: null | number
  after?: null | string
  favoritesOnly?: null | boolean
  includeNotified?: null | boolean
  omit?: null | number[]
  sort?: null | AlertSearchSort
  createdBefore?: null | ISOTimestamp
  notCreatedBefore?: null | ISOTimestamp
  closedBefore?: null | ISOTimestamp
  notClosedBefore?: null | ISOTimestamp
}

export type AlertSearchSort = 'statusID' | 'dateID' | 'dateIDReverse'

export type ISODuration = string

export type ISORInterval = string

export type ISOTimestamp = string

export type ClockTime = string

export type WeekdayFilter = [
  boolean,
  boolean,
  boolean,
  boolean,
  boolean,
  boolean,
  boolean,
]

export interface Alert {
  id: string
  alertID: number
  status: AlertStatus
  summary: string
  details: string
  createdAt: ISOTimestamp
  serviceID: string
  service?: null | Service
  state?: null | AlertState
  recentEvents: AlertLogEntryConnection
  pendingNotifications: AlertPendingNotification[]
  metrics?: null | AlertMetric
}

export interface AlertMetric {
  escalated: boolean
  closedAt: ISOTimestamp
  timeToAck: ISODuration
  timeToClose: ISODuration
}

export interface AlertPendingNotification {
  destination: string
}

export interface AlertRecentEventsOptions {
  limit?: null | number
  after?: null | string
}

export interface AlertLogEntryConnection {
  nodes: AlertLogEntry[]
  pageInfo: PageInfo
}

export interface AlertLogEntry {
  id: number
  timestamp: ISOTimestamp
  message: string
  state?: null | NotificationState
}

export interface NotificationState {
  details: string
  status?: null | NotificationStatus
  formattedSrcValue: string
}

export type NotificationStatus = 'OK' | 'WARN' | 'ERROR'

export interface AlertState {
  lastEscalation: ISOTimestamp
  stepNumber: number
  repeatCount: number
}

export interface Service {
  id: string
  name: string
  description: string
  escalationPolicyID: string
  escalationPolicy?: null | EscalationPolicy
  isFavorite: boolean
  maintenanceExpiresAt?: null | ISOTimestamp
  onCallUsers: ServiceOnCallUser[]
  integrationKeys: IntegrationKey[]
  labels: Label[]
  heartbeatMonitors: HeartbeatMonitor[]
}

export interface CreateIntegrationKeyInput {
  serviceID?: null | string
  type: IntegrationKeyType
  name: string
}

export interface CreateHeartbeatMonitorInput {
  serviceID?: null | string
  name: string
  timeoutMinutes: number
}

export interface UpdateHeartbeatMonitorInput {
  id: string
  name?: null | string
  timeoutMinutes?: null | number
}

export type HeartbeatMonitorState = 'inactive' | 'healthy' | 'unhealthy'

export interface HeartbeatMonitor {
  id: string
  serviceID: string
  name: string
  timeoutMinutes: number
  lastState: HeartbeatMonitorState
  lastHeartbeat?: null | ISOTimestamp
  href: string
}

export interface Label {
  key: string
  value: string
}

export interface IntegrationKey {
  id: string
  serviceID: string
  type: IntegrationKeyType
  name: string
  href: string
}

export type IntegrationKeyType =
  | 'generic'
  | 'grafana'
  | 'site24x7'
  | 'prometheusAlertmanager'
  | 'email'

export interface ServiceOnCallUser {
  userID: string
  userName: string
  stepNumber: number
}

export interface EscalationPolicy {
  id: string
  name: string
  description: string
  repeat: number
  isFavorite: boolean
  assignedTo: Target[]
  steps: EscalationPolicyStep[]
  notices: Notice[]
}

export type AlertStatus =
  | 'StatusAcknowledged'
  | 'StatusClosed'
  | 'StatusUnacknowledged'

export interface Target {
  id: string
  type: TargetType
  name: string
}

export interface TargetInput {
  id: string
  type: TargetType
}

export type TargetType =
  | 'escalationPolicy'
  | 'notificationChannel'
  | 'slackChannel'
  | 'notificationPolicy'
  | 'rotation'
  | 'service'
  | 'schedule'
  | 'user'
  | 'integrationKey'
  | 'userOverride'
  | 'notificationRule'
  | 'contactMethod'
  | 'heartbeatMonitor'
  | 'calendarSubscription'
  | 'userSession'

export interface ServiceConnection {
  nodes: Service[]
  pageInfo: PageInfo
}

export interface UserConnection {
  nodes: User[]
  pageInfo: PageInfo
}

export interface AuthSubjectConnection {
  nodes: AuthSubject[]
  pageInfo: PageInfo
}

export interface PageInfo {
  endCursor?: null | string
  hasNextPage: boolean
}

export interface UpdateUserInput {
  id: string
  name?: null | string
  email?: null | string
  role?: null | UserRole
  statusUpdateContactMethodID?: null | string
}

export interface AuthSubjectInput {
  userID: string
  providerID: string
  subjectID: string
}

export type UserRole = 'unknown' | 'user' | 'admin'

export interface User {
  id: string
  role: UserRole
  name: string
  email: string
  contactMethods: UserContactMethod[]
  notificationRules: UserNotificationRule[]
  calendarSubscriptions: UserCalendarSubscription[]
  statusUpdateContactMethodID: string
  authSubjects: AuthSubject[]
  sessions: UserSession[]
  onCallSteps: EscalationPolicyStep[]
  isFavorite: boolean
}

export interface UserSession {
  id: string
  current: boolean
  userAgent: string
  createdAt: ISOTimestamp
  lastAccessAt: ISOTimestamp
}

export interface UserNotificationRule {
  id: string
  delayMinutes: number
  contactMethodID: string
  contactMethod?: null | UserContactMethod
}

export type ContactMethodType = 'SMS' | 'VOICE' | 'EMAIL' | 'WEBHOOK'

export interface UserContactMethod {
  id: string
  type?: null | ContactMethodType
  name: string
  value: string
  formattedValue: string
  disabled: boolean
  lastTestVerifyAt?: null | ISOTimestamp
  lastTestMessageState?: null | NotificationState
  lastVerifyMessageState?: null | NotificationState
}

export interface CreateUserContactMethodInput {
  userID: string
  type: ContactMethodType
  name: string
  value: string
  newUserNotificationRule?: null | CreateUserNotificationRuleInput
}

export interface CreateUserNotificationRuleInput {
  userID?: null | string
  contactMethodID?: null | string
  delayMinutes: number
}

export interface UpdateUserContactMethodInput {
  id: string
  name?: null | string
  value?: null | string
}

export interface SendContactMethodVerificationInput {
  contactMethodID: string
}

export interface VerifyContactMethodInput {
  contactMethodID: string
  code: number
}

export interface AuthSubject {
  providerID: string
  subjectID: string
  userID: string
}

export interface Notice {
  type: NoticeType
  message: string
  details: string
}

export type NoticeType = 'WARNING' | 'ERROR' | 'INFO'

type ConfigID =
  | 'General.ApplicationName'
  | 'General.PublicURL'
  | 'General.GoogleAnalyticsID'
  | 'General.NotificationDisclaimer'
  | 'General.DisableMessageBundles'
  | 'General.ShortURL'
  | 'General.DisableSMSLinks'
  | 'General.DisableLabelCreation'
  | 'General.DisableCalendarSubscriptions'
  | 'Maintenance.AlertCleanupDays'
  | 'Maintenance.APIKeyExpireDays'
  | 'Maintenance.ScheduleCleanupDays'
  | 'Auth.RefererURLs'
  | 'Auth.DisableBasic'
  | 'GitHub.Enable'
  | 'GitHub.NewUsers'
  | 'GitHub.ClientID'
  | 'GitHub.ClientSecret'
  | 'GitHub.AllowedUsers'
  | 'GitHub.AllowedOrgs'
  | 'GitHub.EnterpriseURL'
  | 'OIDC.Enable'
  | 'OIDC.NewUsers'
  | 'OIDC.OverrideName'
  | 'OIDC.IssuerURL'
  | 'OIDC.ClientID'
  | 'OIDC.ClientSecret'
  | 'OIDC.Scopes'
  | 'OIDC.UserInfoEmailPath'
  | 'OIDC.UserInfoEmailVerifiedPath'
  | 'OIDC.UserInfoNamePath'
  | 'Mailgun.Enable'
  | 'Mailgun.APIKey'
  | 'Mailgun.EmailDomain'
  | 'Slack.Enable'
  | 'Slack.ClientID'
  | 'Slack.ClientSecret'
  | 'Slack.AccessToken'
  | 'Slack.SigningSecret'
  | 'Slack.InteractiveMessages'
  | 'Twilio.Enable'
  | 'Twilio.AccountSID'
  | 'Twilio.AuthToken'
  | 'Twilio.AlternateAuthToken'
  | 'Twilio.FromNumber'
  | 'Twilio.MessagingServiceSID'
  | 'Twilio.DisableTwoWaySMS'
  | 'Twilio.SMSCarrierLookup'
  | 'Twilio.SMSFromNumberOverride'
  | 'SMTP.Enable'
  | 'SMTP.From'
  | 'SMTP.Address'
  | 'SMTP.DisableTLS'
  | 'SMTP.SkipVerify'
  | 'SMTP.Username'
  | 'SMTP.Password'
  | 'Webhook.Enable'
  | 'Webhook.AllowedURLs'
  | 'Feedback.Enable'
  | 'Feedback.OverrideURL'
