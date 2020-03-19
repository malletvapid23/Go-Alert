// eslint-disable-next-line @typescript-eslint/no-unused-vars
namespace Cypress {
  interface Chainable {
    createAlert: typeof createAlert
    createManyAlerts: typeof createManyAlerts
    closeAlert: typeof closeAlert
    createAlertLogs: typeof createAlertLogs
  }
}

interface Alert {
  number: number
  id: number
  summary: string
  details: string
  serviceID: string
  service: Service
}

interface AlertOptions {
  summary?: string
  details?: string
  serviceID?: string

  service?: ServiceOptions
}

interface AlertLogOptions {
  count?: number
  alertID?: number
  alert?: AlertOptions
}

interface AlertLogs {
  alert: Alert
  logs: Array<AlertLog>
}
interface AlertLog {
  timestamp: string
  message: string
}
