/* eslint-disable react/prop-types */
/* eslint-disable react/jsx-key */
/* eslint-disable no-unused-vars */
import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { useEquipmentStore } from "../../../contextProviders/equipmentContext";
import { useGlobalStore } from "../../../contextProviders/globalContext";
import { Link, useLocation } from "react-router-dom";

const useStyles = makeStyles({
  card: {
    display: "flex",
    borderRadius: "10px",
    alignItems: "center",
  },
  cardDetails: {
    flex: 1,
    height: "300px",
    display: "flex",
    alignItems: "center",
  },
  cardMedia: {
    width: "30%",
    paddingBottom: "0px",
    paddingRight: "15px",
    borderRight: "2px solid #148FD3",
  },
});

const trim = (text) => {
  let t = "";
  let i = 0;
  while (i <= 250) {
    t = t + text[i];
    i++;
  }
  t = t + "......";
  return t;
};

function ItemList({ items }) {
  const equipmentStore = useEquipmentStore();

  const globalStore = useGlobalStore();
  const classes = useStyles();
  return (
    <div className="search-results">
      {globalStore.isLoggedIn ? (
        <div>
          {items.map((itemObj, index) => {
            const item = itemObj.Item;
            const society = itemObj.Society;
            return (
              <Link to="/itempage">
                <div
                  className="item-card"
                  key={item.id}
                  onClick={() => {
                    globalStore.equipmentFunction(item);
                  }}
                >
                  <div className="item-image">
                    <img src={"https://source.unsplash.com/random/200x200"} />
                  </div>
                  <div className="item-details">
                    <h4>{item.Name}</h4>
                    <h6>Belongs to : {item.SocietyName}</h6>
                    <p>{item.available}</p>
                    <p>
                      {item.details.length > 250 ? (
                        <span>
                          {trim(item.details)}{" "}
                          <span style={{ color: "#ccc" }}>more</span>
                        </span>
                      ) : (
                        item.details
                      )}
                    </p>
                  </div>
                </div>
              </Link>
            );
          })}
          <div className="quotes">
            <i>Looks like you have browsed all the results ....</i>
          </div>
        </div>
      ) : (
        <div className="quotes">
          <i>Login first to see the results ....</i>
        </div>
      )}
    </div>
  );
}

export default ItemList;
