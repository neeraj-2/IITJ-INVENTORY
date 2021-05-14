/* eslint-disable no-undef */
/// <reference types="cypress" />

describe("Checking Cards on Home page", () => {
  describe("Check Robotics Club card content", () => {
    it("Check Content of CARD", () => {
      cy.visit("/");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(1) > button > span.MuiButton-label > div > button > img"
      ).should("be.visible");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(1) > button > span.MuiButton-label > div > button > div"
      )
        .should("be.visible")
        .click();
    });

    it("Check Visited Search page", () => {
      cy.url().should("be.equal", "http://localhost:3000/searchPage");
    });

    it("Go back to Home page", () => {
      cy.visit("/");
    });
  });

  describe("Check PCLUB card content", () => {
    it("Check PCLUB Club CARD", () => {
      cy.visit("/");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(2) > button > span.MuiButton-label > div > button > img"
      ).should("be.visible");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(2) > button > span.MuiButton-label > div > button > div"
      )
        .should("be.visible")
        .click();
    });

    it("Check Visited Search page", () => {
      cy.url().should("be.equal", "http://localhost:3000/searchPage");
    });

    it("Go back to Home page", () => {
      cy.visit("/");
    });
  });

  describe("Check Drama Club card content", () => {
    it("Check Drama Club CARD", () => {
      cy.visit("/");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(3) > button > span.MuiButton-label > div > button > img"
      ).should("be.visible");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(3) > button > span.MuiButton-label > div > button > div"
      )
        .should("be.visible")
        .click();
    });

    it("Check Visited Search page", () => {
      cy.url().should("be.equal", "http://localhost:3000/searchPage");
    });

    it("Go back to Home page", () => {
      cy.visit("/");
    });
  });
  describe("Check Music Clubs card content", () => {
    it("Check Music Club CARD", () => {
      cy.visit("/");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(4) > button > span.MuiButton-label > div > button > img"
      ).should("be.visible");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(4) > button > span.MuiButton-label > div > button > div"
      )
        .should("be.visible")
        .click();
    });

    it("Check Visited Search page", () => {
      cy.url().should("be.equal", "http://localhost:3000/searchPage");
    });

    it("Go back to Home page", () => {
      cy.visit("/");
    });
  });
  describe("Check Dance Club card content", () => {
    it("Check Dance Club CARD", () => {
      cy.visit("/");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(5) > button > span.MuiButton-label > div > button > img"
      ).should("be.visible");
      cy.get(
        "#root > div > div.router-div > div > div.Grid-container.society-cards-container > div > div:nth-child(5) > button > span.MuiButton-label > div > button > div"
      )
        .should("be.visible")
        .click();
    });

    it("Check Visited Search page", () => {
      cy.url().should("be.equal", "http://localhost:3000/searchPage");
    });

    it("Go back to Home page", () => {
      cy.visit("/");
    });
  });
});
