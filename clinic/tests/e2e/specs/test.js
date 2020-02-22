// https://docs.cypress.io/api/introduction/api.html

describe('Login Page Test', () => {
  it('Visits the app root url', () => {
    cy.visit('/')
    cy.contains('div#login-title', 'เข้าสู่ระบบ')
  })

  it('Login success', () => {
    cy.visit('/')
    cy.get('#username-input')
      .type('1234')

    cy.get('#password-input')
      .type('12345')

    cy.get('#login-button')
      .click()

    cy.contains('div#menu-title', 'MENU')
  })
})