export const createEquipmentStore = () => {
  return {
    items: [],
    setItems(items) {
      this.items = items;
      // console.log("Items",this.items)
    },
  };
};
