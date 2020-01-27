import { Chance } from 'chance'
import { testScreen } from '../support'

const c = new Chance()

testScreen('Calendar Subscriptions', testSubs)

function testSubs(screen: ScreenFormat) {
  beforeEach(() => {
    cy.resetCalendarSubscriptions()
  })

  describe('Creation', () => {
    let sched: Schedule
    beforeEach(() => {
      cy.createSchedule().then(s => {
        sched = s
      })
    })

    it('should create a subscription from a schedule', () => {
      const name = c.word({ length: 5 })
      const defaultCptn =
        'Subscribe to your shifts on this calendar from your preferred calendar app'

      cy.visit(`/schedules/${sched.id}`)

      cy.get('body').should('contain', defaultCptn)
      cy.get('body').should('not.contain', 'You have 1 active subscription')

      // fill form out and submit
      cy.get('[data-cy="subscribe-btn"]').click()
      cy.dialogTitle('Create New Calendar Subscription')
      cy.dialogForm({
        name,
        'reminderMinutes[0]': 'At time of shift',
      })
      cy.dialogClick('Submit')
      cy.dialogTitle('Success!')
      // todo: verify url generation
      cy.dialogFinish('Done')

      cy.get('body').should('not.contain', defaultCptn)
      cy.get('body').should('contain', 'You have 1 active subscription')
    })

    it('should create a subscription from their profile', () => {
      const name = c.word({ length: 5 })

      cy.visit('/profile/schedule-calendar-subscriptions')

      cy.get('[data-cy="list-empty-message"]').should(
        'contain',
        'You are not subscribed to any schedules.',
      )
      cy.get('[data-cy=calendar-subscriptions]').should('not.contain', name)

      // fill form out and submit
      cy.pageFab()
      cy.dialogTitle('Create New Calendar Subscription')
      cy.dialogForm({
        name,
        scheduleID: sched.name,
        'reminderMinutes[0]': 'At time of shift',
      })
      cy.dialogClick('Submit')
      cy.dialogTitle('Success!')
      // todo: verify url generation
      cy.dialogFinish('Done')

      cy.get('[data-cy="list-empty-message"]').should('not', 'exist')
      cy.get('[data-cy=calendar-subscriptions]').should('contain', name)
    })

    it('should add and remove additional valarms', () => {
      cy.visit('/profile/schedule-calendar-subscriptions')
      cy.pageFab()

      const check = (shouldExist: Array<boolean>) => {
        cy.get('input[name="reminderMinutes[0]"]').should(
          shouldExist[0] ? 'exist' : 'not.exist',
        )
        cy.get('input[name="reminderMinutes[1]"]').should(
          shouldExist[1] ? 'exist' : 'not.exist',
        )
        cy.get('input[name="reminderMinutes[2]"]').should(
          shouldExist[2] ? 'exist' : 'not.exist',
        )
        cy.get('input[name="reminderMinutes[3]"]').should(
          shouldExist[3] ? 'exist' : 'not.exist',
        )
        cy.get('input[name="reminderMinutes[4]"]').should(
          shouldExist[4] ? 'exist' : 'not.exist',
        )
      }

      // only 1st valarm field should exist to start
      check([true, false, false, false, false])

      cy.dialogForm({ 'reminderMinutes[0]': 'At time of shift' })
      check([true, true, false, false, false])

      cy.dialogForm({ 'reminderMinutes[1]': 'At time of shift' })
      check([true, true, true, false, false])

      cy.dialogForm({ 'reminderMinutes[2]': 'At time of shift' })
      check([true, true, true, true, false])

      cy.dialogForm({ 'reminderMinutes[3]': 'At time of shift' })
      check([true, true, true, true, true])

      // clearing the optional valarm fields should remove the redundant fields
      cy.dialogForm({ 'reminderMinutes[3]': '' })
      check([true, true, true, true, false])

      cy.dialogForm({ 'reminderMinutes[2]': '' })
      check([true, true, true, false, false])

      cy.dialogForm({ 'reminderMinutes[1]': '' })
      check([true, true, false, false, false])

      cy.dialogForm({ 'reminderMinutes[0]': '' })
      check([true, false, false, false, false])
    })
  })

  describe('Schedule', () => {
    let sched: Schedule
    beforeEach(() => {
      cy.createSchedule().then(s => {
        sched = s
        cy.visit(`/schedules/${s.id}`)
      })
    })

    it('should go to the profile subscriptions list from a schedule', () => {
      const flatListHeader =
        'Showing your current on-call subscriptions for all schedules'

      cy.get('body').should('not.contain', flatListHeader)
      cy.get('[data-cy="manage-subscriptions-link"]')
        .should('contain', 'Manage subscriptions')
        .click()
      cy.get('body').should('contain', flatListHeader)
    })

    it('should update button caption text after a subscription is created', () => {
      const defaultCptn =
        'Subscribe to your shifts on this calendar from your preferred calendar app'
      const oneSubCptn = 'You have 1 active subscription for this schedule'
      const multipleSubsCptn =
        'You have 2 active subscriptions for this schedule'

      const caption = '[data-cy="subscribe-btn-txt"]'
      const captionLoading = '[data-cy="subscribe-btn-txt-loading"]'

      cy.get(captionLoading).should('not.exist')
      cy.get(caption).should('contain', defaultCptn)
      cy.get(caption).should('not.contain', oneSubCptn)
      cy.get(caption).should('not.contain', multipleSubsCptn)

      cy.createCalendarSubscription({ scheduleID: sched.id }).then(() => {
        cy.reload()
        cy.get(captionLoading).should('not.exist')

        cy.get(caption).should('not.contain', defaultCptn)
        cy.get(caption).should('contain', oneSubCptn)
        cy.get(caption).should('not.contain', multipleSubsCptn)

        cy.createCalendarSubscription({ scheduleID: sched.id }).then(() => {
          cy.reload()
          cy.get(captionLoading).should('not.exist')

          cy.get(caption).should('not.contain', defaultCptn)
          cy.get(caption).should('not.contain', oneSubCptn)
          cy.get(caption).should('contain', multipleSubsCptn)
        })
      })
    })
  })

  describe('Profile', () => {
    let cs: CalendarSubscription
    beforeEach(() => {
      cy.createCalendarSubscription().then(sub => {
        cs = sub
        cy.visit('/profile/schedule-calendar-subscriptions')
      })
    })

    it('should navigate to and from the subscriptions list', () => {
      cy.visit('/profile')
      cy.navigateToAndFrom(
        screen,
        'Profile',
        'Profile',
        'Schedule Calendar Subscriptions',
        '/profile/schedule-calendar-subscriptions',
      )
    })

    it('should view the subscriptions list', () => {
      cy.get('body').should(
        'contain',
        'Showing your current on-call subscriptions for all schedules',
      )
      cy.get('[data-cy=calendar-subscriptions]').should('contain', cs.name)
      cy.get('[data-cy=calendar-subscriptions]').should(
        'contain',
        'Last sync: Never',
      )
    })

    it('should edit a subscription', () => {
      const name = 'SM Subscription ' + c.word({ length: 8 })

      cy.get('[data-cy=calendar-subscriptions]').should('contain', cs.name)
      cy.get('[data-cy=calendar-subscriptions]').should(
        'contain',
        'Last sync: Never',
      )

      cy.get('[data-cy=calendar-subscriptions]')
        .contains('li', cs.name)
        .find('[data-cy=other-actions]')
        .menu('Edit')

      cy.dialogTitle('Edit Calendar Subscription')
      cy.dialogForm({ name })
      cy.dialogFinish('Submit')

      cy.get('[data-cy=calendar-subscriptions]').should('contain', name)
      cy.get('[data-cy=calendar-subscriptions]').should(
        'contain',
        'Last sync: Never',
      )
    })

    it('should delete a subscription', () => {
      cy.get('[data-cy="list-empty-message"]').should('not', 'exist')
      cy.get('[data-cy=calendar-subscriptions]').should('contain', cs.name)

      cy.get('[data-cy=calendar-subscriptions]')
        .contains('li', cs.name)
        .find('[data-cy=other-actions]')
        .menu('Delete')

      cy.dialogFinish('Confirm')

      cy.get('[data-cy="list-empty-message"]').should(
        'contain',
        'You are not subscribed to any schedules.',
      )
      cy.get('[data-cy=calendar-subscriptions]').should('not.contain', cs.name)
    })

    it('should visit a schedule from the subheader link', () => {
      cy.get('[data-cy="subscribe-btn"]').should('not.exist')

      cy.get('[data-cy=calendar-subscriptions]')
        .contains('li a', cs.schedule.name)
        .click()

      cy.get('[data-cy="subscribe-btn"]').should('exist')
    })

    it('should not show route link unless on personal profile', () => {
      cy.fixture('users').then(users => {
        cy.visit(`/users/${users[0].id}`)
        cy.get('[data-cy="route-links"]').should(
          'not.contain',
          'Schedule Calendar Subscriptions',
        )
      })
    })

    it('should show an icon if a subscription is disabled, and vice-versa', () => {
      cy.createCalendarSubscription({ disabled: true }).then(disabledCs => {
        cy.reload()

        cy.get('[data-cy=calendar-subscriptions] li')
          .contains(cs.name)
          .find('[data-cy="warning-icon"]')
          .should('not.exist')

        cy.get('[data-cy=calendar-subscriptions] li')
          .contains(disabledCs.name)
          .parent()
          .parent()
          .find('[data-cy="warning-icon"]') // two divs of separation
          .should('exist')
      })
    })
  })
}
