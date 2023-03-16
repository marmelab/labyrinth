import '@testing-library/cypress/add-commands';

describe('Admin Login', () => {
    it('Should display and error if invalid credentials are provided', () => {
        cy.visit('https://localhost:9443/admin/');

        cy.findByLabelText(/Username/i, {timeout: 7000})
            .should('be.visible')
            .type("testuser@example.org");

        cy.findByLabelText(/Password/i, {timeout: 7000})
            .should('be.visible')
            .type("invalid");

            cy.findByRole("button", {name: "Sign in", timeout: 7000})
            .should('be.visible')
            .click();


        cy.findByText("Invalid credentials.")
            .should('be.visible');
    });

    it('Should display the admin if creadentials are valid', () => {
        cy.visit('https://localhost:9443/admin/');

        cy.findByLabelText(/Username/i, {timeout: 7000})
            .should('be.visible')
            .type(Cypress.env("ADMIN_USER"));

        cy.findByLabelText(/Password/i, {timeout: 7000})
            .should('be.visible')
            .type(Cypress.env("ADMIN_PASSWORD"));

        cy.findByRole("button", {name: "Sign in", timeout: 7000})
            .should('be.visible')
            .click();

        cy.findByRole("menuitem", {name: /Boards/, timeout: 7000})
            .should('be.visible');

        cy.findByRole("menuitem", {name: /Users/, timeout: 7000})
            .should('be.visible');
    });
});
