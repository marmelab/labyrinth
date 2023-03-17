import { AdminSignIn } from '../../pages/AdminSignIn';

describe('Admin Login', () => {
    const adminSignInPage = new AdminSignIn();
    
    it('Should display and error if invalid credentials are provided', () => {
        adminSignInPage.signIn("testuser@example.org", "invalid");

        cy.findByText("Invalid credentials.")
            .should('be.visible');
    });

    it('Should display the admin if creadentials are valid', () => {
        adminSignInPage.signIn(Cypress.env("ADMIN_USER"), Cypress.env("ADMIN_PASSWORD"));

        [/Boards/, /Users/].forEach((name) => {
            cy.findByRole("menuitem", {name})
                .should('be.visible');
        });
    });
});
