import '@testing-library/cypress/add-commands';

export class AdminSignIn {
    constructor() {
      this.url = "/admin/"
    }

    visit() {
        cy.visit(this.url)
    }

    getUsernameField() {
        return cy.findByLabelText(/Username/i).should('be.visible');
    }

    getPasswordField() {
        return cy.findByLabelText(/Password/i).should('be.visible');
    }

    getSignInButton() {
        return cy.findByRole("button", {name: "Sign in"}).should('be.visible');
    }

    signIn(username, password)  {
        this.visit();
        this.getUsernameField().type(username);
        this.getPasswordField().type(password);
        this.getSignInButton().click();
    }
}