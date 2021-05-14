/* eslint-disable react/jsx-key */
/* eslint-disable no-unused-vars */
import React, { useState, useEffect } from "react";
import { Button, Grid } from "@material-ui/core";
import Divider from "@material-ui/core/Divider";
import RequestExtension from "../modal/RequestExtension";
import Dialog from "@material-ui/core/Dialog";
import DialogActions from "@material-ui/core/DialogActions";
import { useGlobalStore } from "../../contextProviders/globalContext";
import axios from "axios";

function User() {
  // const [contact, setContact] = useState("contact");

  // Donot touch this
  // const [edit, setEdit] = useState("edit");
  // const [save, setSave] = useState("save");
  // console.log(save);
  // function editButton() {
  //   if (edit === "edit") {
  //     setEdit("save");
  //     setSave("edit");
  //   } else {
  //     setEdit("edit");
  //     setSave("save");
  //   }
  // }

  // function changeContact(e) {
  //   setContact(e.target.value);
  // }
  const handleOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };
  const maxWidth = React.useState("md");
  const [open, setOpen] = React.useState(false);
  const [email, setEmail] = React.useState("");
  const [name, setName] = React.useState("");
  const [items, setItems] = React.useState([]);

  const globalStore = useGlobalStore();

  useEffect(async () => {
    const res = await axios.get(
      `http://127.0.0.1:8080/student/profile?jwt=${globalStore.token}`,
      { "Access-Control-Allow-Origin": "http://localhost:3000" }
    );
    setEmail(res.data.user.Email);
    setName(res.data.user.Name);
    setItems(res.data.issued);
  }, []);

  return (
    <div>
      {globalStore.isLoggedIn ? (
        <div className="user-profile-page-container">
          <Grid className="user_profile_page">
            <Grid className="user_profile" xs={3}>
              <Grid>
                <img
                  src="https://source.unsplash.com/random/200x200"
                  alt="Images"
                />
              </Grid>
              <Grid className="user_information">
                <Grid>
                  <p className="user_inf">
                    Name: <span>{name}</span>
                  </p>
                  <p className="user_inf">
                    Email: <span>{email}</span>
                  </p>
                  <p className="user_inf">
                    Role: <span>Student</span>
                  </p>
                  {/* <p className="user_inf">
              Contact: <span>{contact}</span>
            </p> */}
                </Grid>
              </Grid>
            </Grid>
            <Divider orientation="vertical" flexItem className="divider" />
            <Grid item className="user_transactions" xs={8}>
              <Grid className="user_trans_heading">
                <span style={{ backgroundColor: "#148FD3", color: "#fff" }}>
                  Active Requests
                </span>
              </Grid>
              {items.map((item, index) => {
                // console.log(item)
                return (
                  <div>
                    <Grid className="user_request_content">
                      <Grid className="user_active_request">
                        <Grid className="imges">
                          <img
                            src="https://source.unsplash.com/random/150x150"
                            alt="Photos"
                          />
                        </Grid>
                        <Grid>
                          <div className="user_item_info">
                            <h2 className="heading">{item.Name}</h2>
                            <p className="item-details">
                              Belongs to: <span>{item.SocietyName}</span>
                            </p>
                            <p className="item-details">
                              Issuees on: <span>{item.IssueDate}</span>
                            </p>
                            <p className="item-details">
                              Description:
                              <span>{item.details}</span>
                            </p>
                            <p className="item-details">
                              <b>
                                Deadline: <span>{item.DueDate}</span>
                              </b>
                            </p>
                          </div>
                        </Grid>
                      </Grid>
                    </Grid>
                  </div>
                );
              })}
            </Grid>
          </Grid>
        </div>
      ) : (
        <div>
          <span>Login first :D</span>
        </div>
      )}
    </div>
  );
}

export default User;
