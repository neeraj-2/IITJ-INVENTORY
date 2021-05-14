/// <reference types="cypress" />
/* eslint-disable no-undef */

describe("Checking Home page", () => {
  describe("Check content", () => {
    it("Checking navbar content of 'Home' Page", () => {
      cy.visit("/");
      cy.get(
        "#root > div > div.header > div > header > div > div.navbar-brand > div > h6 > span"
      ).should("be.visible");
      cy.get("#root > div > div.header > div > header > div").should(
        "be.visible"
      );
      cy.get(
        "#root > div > div.header > div > header > div > div.MuiGrid-root > button"
      ).should("be.visible");
      cy.get(
        "#root > div > div.header > div > header > div > div.MuiGrid-root > button > span.MuiButton-label"
      )
        .should("be.visible")
        .should("have.text", "LogIn")
        .click({ force: true });
    });
  });

  describe("Check functional components", () => {
    it("Checking Search Bar", () => {
      cy.visit("/");
      cy.get("#heading").should("be.visible");
    });
    it("Checking content for search bar", () => {
      cy.get("#root > div > div.header > div > div > div > p").should(
        "be.visible"
      );
    });

    it("Let type something in serach bar and see results", () => {
      cy.get("#search-input").should("be.visible").type("Robotics");
      cy.get(
        "#root > div > div.header > div > div > form > div > a > span.MuiButton-label"
      ).click();
    });

    it("Check results before loggied in", () => {
      cy.get(
        "#root > div > div.router-div > div > div:nth-child(2) > div > div > div > div > i"
      )
        .should("be.visible")
        .should("have.text", "Login first to see the results ....");
    });

    it("Check visited URL", () => {
      cy.url().should("be.equal", "http://localhost:3000/SearchPage");
    });
  });
});
