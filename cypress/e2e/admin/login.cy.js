import '@testing-library/cypress/add-commands';

const adminSignIn = function(username, password) {
    cy.visit('https://localhost:9443/admin/');
    
    cy.findByLabelText(/Username/i)
        .should('be.visible')
        .type(username);

    cy.findByLabelText(/Password/i)
        .should('be.visible')
        .type(password);

    cy.findByRole("button", {name: "Sign in"})
        .should('be.visible')
        .click();
}

describe('Admin Login', () => {
    it('Should display and error if invalid credentials are provided', () => {
        adminSignIn("testuser@example.org", "invalid")

        cy.findByText("Invalid credentials.")
            .should('be.visible');
    });

    it('Should display the admin if creadentials are valid', () => {
        adminSignIn(Cypress.env("ADMIN_USER"), Cypress.env("ADMIN_PASSWORD"));

        [/Boards/, /Users/].forEach((name) => {
            cy.findByRole("menuitem", {name})
                .should('be.visible');
        });
    });
});
