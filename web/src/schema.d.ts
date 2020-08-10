// Code generated by devtools/gqltsgen DO NOT EDIT.

export interface Query {
  phoneNumberInfo?: PhoneNumberInfo
  user?: User
  users: UserConnection
  alert?: Alert
  alerts: AlertConnection
  service?: Service
  integrationKey?: IntegrationKey
  heartbeatMonitor?: HeartbeatMonitor
  services: ServiceConnection
  rotation?: Rotation
  rotations: RotationConnection
  schedule?: Schedule
  userCalendarSubscription?: UserCalendarSubscription
  schedules: ScheduleConnection
  escalationPolicy?: EscalationPolicy
  escalationPolicies: EscalationPolicyConnection
  authSubjectsForProvider: AuthSubjectConnection
  timeZones: TimeZoneConnection
  labels: LabelConnection
  labelKeys: StringConnection
  labelValues: StringConnection
  userOverrides: UserOverrideConnection
  userOverride?: UserOverride
  config: ConfigValue[]
  configHints: ConfigHint[]
  systemLimits: SystemLimit[]
  userContactMethod?: UserContactMethod
  slackChannels: SlackChannelConnection
  slackChannel?: SlackChannel
}

export interface SlackChannelSearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
}

export interface SlackChannel {
  id: string
  name: string
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

export interface UserOverrideSearchOptions {
  first?: number
  after?: string
  omit?: string[]
  scheduleID?: string
  filterAddUserID?: string[]
  filterRemoveUserID?: string[]
  filterAnyUserID?: string[]
  start?: ISOTimestamp
  end?: ISOTimestamp
}

export interface UserOverrideConnection {
  nodes: UserOverride[]
  pageInfo: PageInfo
}

export interface UserOverride {
  id: string
  start: ISOTimestamp
  end: ISOTimestamp
  addUserID: string
  removeUserID: string
  addUser?: User
  removeUser?: User
  target: Target
}

export interface LabelSearchOptions {
  first?: number
  after?: string
  search?: string
  uniqueKeys?: boolean
  omit?: string[]
}

export interface LabelKeySearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
}

export interface LabelValueSearchOptions {
  key: string
  first?: number
  after?: string
  search?: string
  omit?: string[]
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
}

export interface Mutation {
  debugCarrierInfo: DebugCarrierInfo
  debugSendSMS?: DebugSendSMSInfo
  addAuthSubject: boolean
  deleteAuthSubject: boolean
  updateUser: boolean
  testContactMethod: boolean
  updateAlerts?: Alert[]
  updateRotation: boolean
  escalateAlerts?: Alert[]
  setFavorite: boolean
  updateService: boolean
  updateEscalationPolicy: boolean
  updateEscalationPolicyStep: boolean
  deleteAll: boolean
  createAlert?: Alert
  createService?: Service
  createEscalationPolicy?: EscalationPolicy
  createEscalationPolicyStep?: EscalationPolicyStep
  createRotation?: Rotation
  createIntegrationKey?: IntegrationKey
  createHeartbeatMonitor?: HeartbeatMonitor
  setLabel: boolean
  createSchedule?: Schedule
  createUserCalendarSubscription: UserCalendarSubscription
  updateUserCalendarSubscription: boolean
  updateScheduleTarget: boolean
  createUserOverride?: UserOverride
  createUserContactMethod?: UserContactMethod
  createUserNotificationRule?: UserNotificationRule
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
  details?: string
  serviceID: string
  sanitize?: boolean
}

export interface CreateUserCalendarSubscriptionInput {
  name: string
  reminderMinutes?: number[]
  scheduleID: string
  disabled?: boolean
}

export interface UpdateUserCalendarSubscriptionInput {
  id: string
  name?: string
  reminderMinutes?: number[]
  disabled?: boolean
}

export interface UserCalendarSubscription {
  id: string
  name: string
  reminderMinutes: number[]
  scheduleID: string
  schedule?: Schedule
  lastAccess: ISOTimestamp
  disabled: boolean
  url?: string
}

export interface ConfigValueInput {
  id: string
  value: string
}

export interface UpdateUserOverrideInput {
  id: string
  start?: ISOTimestamp
  end?: ISOTimestamp
  addUserID?: string
  removeUserID?: string
}

export interface CreateUserOverrideInput {
  scheduleID?: string
  start: ISOTimestamp
  end: ISOTimestamp
  addUserID?: string
  removeUserID?: string
}

export interface CreateScheduleInput {
  name: string
  description?: string
  timeZone: string
  favorite?: boolean
  targets?: ScheduleTargetInput[]
  newUserOverrides?: CreateUserOverrideInput[]
}

export interface ScheduleTargetInput {
  scheduleID?: string
  target?: TargetInput
  newRotation?: CreateRotationInput
  rules: ScheduleRuleInput[]
}

export interface ScheduleRuleInput {
  id?: string
  start?: ClockTime
  end?: ClockTime
  weekdayFilter?: boolean[]
}

export interface SetLabelInput {
  target?: TargetInput
  key: string
  value: string
}

export interface TimeZoneSearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
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
  description?: string
  favorite?: boolean
  escalationPolicyID?: string
  newEscalationPolicy?: CreateEscalationPolicyInput
  newIntegrationKeys?: CreateIntegrationKeyInput[]
  labels?: SetLabelInput[]
  newHeartbeatMonitors?: CreateHeartbeatMonitorInput[]
}

export interface CreateEscalationPolicyInput {
  name: string
  description?: string
  repeat?: number
  steps?: CreateEscalationPolicyStepInput[]
}

export interface CreateEscalationPolicyStepInput {
  escalationPolicyID?: string
  delayMinutes: number
  targets?: TargetInput[]
  newRotation?: CreateRotationInput
  newSchedule?: CreateScheduleInput
}

export interface EscalationPolicyStep {
  id: string
  stepNumber: number
  delayMinutes: number
  targets: Target[]
  escalationPolicy?: EscalationPolicy
}

export interface UpdateScheduleInput {
  id: string
  name?: string
  description?: string
  timeZone?: string
}

export interface UpdateServiceInput {
  id: string
  name?: string
  description?: string
  escalationPolicyID?: string
}

export interface UpdateEscalationPolicyInput {
  id: string
  name?: string
  description?: string
  repeat?: number
  stepIDs?: string[]
}

export interface UpdateEscalationPolicyStepInput {
  id: string
  delayMinutes?: number
  targets?: TargetInput[]
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
  target?: ScheduleTarget
  isFavorite: boolean
}

export interface OnCallShift {
  userID: string
  user?: User
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
  weekdayFilter: boolean[]
  target: Target
}

export interface RotationConnection {
  nodes: Rotation[]
  pageInfo: PageInfo
}

export interface CreateRotationInput {
  name: string
  description?: string
  timeZone: string
  start: ISOTimestamp
  favorite?: boolean
  type: RotationType
  shiftLength?: number
  userIDs?: string[]
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
  name?: string
  description?: string
  timeZone?: string
  start?: ISOTimestamp
  type?: RotationType
  shiftLength?: number
  activeUserIndex?: number
  userIDs?: string[]
}

export interface RotationSearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
  favoritesOnly?: boolean
  favoritesFirst?: boolean
}

export interface EscalationPolicySearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
}

export interface ScheduleSearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
  favoritesOnly?: boolean
  favoritesFirst?: boolean
}

export interface ServiceSearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
  favoritesOnly?: boolean
  favoritesFirst?: boolean
}

export interface UserSearchOptions {
  first?: number
  after?: string
  search?: string
  omit?: string[]
}

export interface AlertSearchOptions {
  filterByStatus?: AlertStatus[]
  filterByServiceID?: string[]
  search?: string
  first?: number
  after?: string
  favoritesOnly?: boolean
  includeNotified?: boolean
  omit?: number[]
}

export type ISOTimestamp = string

export type ClockTime = string

export interface Alert {
  id: string
  alertID: number
  status: AlertStatus
  summary: string
  details: string
  createdAt: ISOTimestamp
  serviceID: string
  service?: Service
  state?: AlertState
  recentEvents: AlertLogEntryConnection
}

export interface AlertRecentEventsOptions {
  limit?: number
  after?: string
}

export interface AlertLogEntryConnection {
  nodes: AlertLogEntry[]
  pageInfo: PageInfo
}

export interface AlertLogEntry {
  id: number
  timestamp: ISOTimestamp
  message: string
  state?: AlertLogEntryState
}

export interface AlertLogEntryState {
  details: string
  status?: AlertLogStatus
}

export type AlertLogStatus = 'OK' | 'WARN' | 'ERROR'

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
  escalationPolicy?: EscalationPolicy
  isFavorite: boolean
  onCallUsers: ServiceOnCallUser[]
  integrationKeys: IntegrationKey[]
  labels: Label[]
  heartbeatMonitors: HeartbeatMonitor[]
}

export interface CreateIntegrationKeyInput {
  serviceID?: string
  type: IntegrationKeyType
  name: string
}

export interface CreateHeartbeatMonitorInput {
  serviceID: string
  name: string
  timeoutMinutes: number
}

export interface UpdateHeartbeatMonitorInput {
  id: string
  name?: string
  timeoutMinutes?: number
}

export type HeartbeatMonitorState = 'inactive' | 'healthy' | 'unhealthy'

export interface HeartbeatMonitor {
  id: string
  serviceID: string
  name: string
  timeoutMinutes: number
  lastState: HeartbeatMonitorState
  lastHeartbeat?: ISOTimestamp
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
  assignedTo: Target[]
  steps: EscalationPolicyStep[]
}

export type AlertStatus =
  | 'StatusAcknowledged'
  | 'StatusClosed'
  | 'StatusUnacknowledged'

export interface Target {
  id: string
  type: TargetType
  name?: string
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
  endCursor?: string
  hasNextPage: boolean
}

export interface UpdateUserInput {
  id: string
  name?: string
  email?: string
  role?: UserRole
  statusUpdateContactMethodID?: string
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
  onCallSteps: EscalationPolicyStep[]
}

export interface UserNotificationRule {
  id: string
  delayMinutes: number
  contactMethodID: string
  contactMethod?: UserContactMethod
}

export type ContactMethodType = 'SMS' | 'VOICE'

export interface UserContactMethod {
  id: string
  type?: ContactMethodType
  name: string
  value: string
  formattedValue: string
  disabled: boolean
}

export interface CreateUserContactMethodInput {
  userID: string
  type: ContactMethodType
  name: string
  value: string
  newUserNotificationRule?: CreateUserNotificationRuleInput
}

export interface CreateUserNotificationRuleInput {
  userID?: string
  contactMethodID?: string
  delayMinutes: number
}

export interface UpdateUserContactMethodInput {
  id: string
  name?: string
  value?: string
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
