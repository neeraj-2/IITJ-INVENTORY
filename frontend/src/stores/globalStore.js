export const createGlobalStore = () => {
  return {
    rules: `none`,
    searchText: null,
    searchfunction(text) {
      this.searchText = text;
      console.log(this.searchText);
    },
    equipmentDetails: null,
    equipmentFunction(item) {
      this.equipmentDetails = item;
      console.log("Equipment Details", this.equipmentDetails);
    },
    loginType: "admin",
    token: null,
    modifyToken(value) {
      this.token = value;
      console.log("Token Value", value);
    },
    isLoggedIn: false,
    logging(flag) {
      this.isLoggedIn = flag;
      console.log("Logging flag value", flag);
    },
  };
};
