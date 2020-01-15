declare namespace Cypress {
  interface Chainable {
    login: typeof login
    adminLogin: typeof adminLogin
  }
}

function login(
  username?: string,
  password?: string,
  tokenOnly: boolean = false,
): Cypress.Chainable<string> {
  if (!username) {
    return cy
      .fixture('profile')
      .then(p => login(p.username, p.password, tokenOnly))
  }
  if (!password) {
    return cy
      .fixture('profile')
      .then(p => login(username, p.password, tokenOnly))
  }

  if (tokenOnly) {
    return cy
      .request({
        url: '/api/v2/identity/providers/basic?noRedirect=1',
        method: 'POST',
        form: true, // indicates the body should be form urlencoded and sets Content-Type: application/x-www-form-urlencoded headers
        body: {
          username,
          password,
        },
        followRedirect: false,
        headers: {
          referer: Cypress.config('baseUrl'),
          Cookie: '',
        },
      })
      .then(res => {
        return res.body
      })
  }

  return cy
    .request({
      url: '/api/v2/identity/providers/basic',
      method: 'POST',
      form: true, // indicates the body should be form urlencoded and sets Content-Type: application/x-www-form-urlencoded headers
      body: {
        username,
        password,
      },
      followRedirect: false,
      headers: {
        referer: Cypress.config('baseUrl'),
        Cookie: '',
      },
    })
    .then(res => {
      expect(res.redirectedToUrl, 'response redirect').to.eq(
        Cypress.config('baseUrl') + '/',
      )
      return ''
    }) as Cypress.Chainable<string>
}

function adminLogin(tokenOnly: boolean = false): Cypress.Chainable<string> {
  return cy
    .fixture('profileAdmin')
    .then(p => login(p.username, p.password, tokenOnly))
}

Cypress.Commands.add('login', login)
Cypress.Commands.add('adminLogin', adminLogin)
