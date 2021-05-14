/* eslint-disable no-undef */
/// <reference types="cypress" />

describe("Check Footer Content", () => {
  describe("Check visibility of different static components", () => {
    it("check image", () => {
      cy.visit("/");
      cy.get("#logo-image").should("be.visible");
      cy.get(
        "#root > div > div.footer > div > div > div.MuiGrid-root.column-2.MuiGrid-item.MuiGrid-grid-xs-12.MuiGrid-grid-sm-4.MuiGrid-grid-md-4 > div > div > p"
      )
        .should("be.visible")
        .should("have.text", "We are CSE Geeks(Tarun n Neeraj)");
    });

    it("Check Column one content", () => {
      cy.get(".footer-list").should("be.visible");
    });

    it("Check IITJ tag in footer", () => {
      cy.get(
        "#root > div > div.footer > div > div > div.MuiGrid-root.column-3.MuiGrid-item.MuiGrid-grid-xs-12.MuiGrid-grid-sm-4.MuiGrid-grid-md-4 > p"
      ).should("be.visible");
    });
  });

  describe("Check Copyright Tags", () => {
    it("Check Visiblility of Copyright tag", () => {
      cy.get(
        "#root > div > div.footer > div > div > div.MuiGrid-root.secondary-footer.MuiGrid-container.MuiGrid-item.MuiGrid-grid-xs-12.MuiGrid-grid-sm-12.MuiGrid-grid-md-12"
      ).should("be.visible");
    });
  });
});
