/* eslint-disable react/prop-types */
/* eslint-disable no-unused-vars */
/* eslint-disable react/jsx-key */
/* eslint-disable no-unused-vars */
import React, { createContext, useContext } from "react";

import { createGlobalStore } from "../stores/globalStore";

import { useLocalObservable } from "mobx-react";

const GlobalContext = createContext(null);

export const GlobalProvider = ({ children }) => {
  const globalStore = useLocalObservable(createGlobalStore);

  return (
    <GlobalContext.Provider value={globalStore}>
      {children}
    </GlobalContext.Provider>
  );
};
export const useGlobalStore = () => useContext(GlobalContext);
