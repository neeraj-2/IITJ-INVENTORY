/* eslint-disable no-unused-vars */
import React, { Component } from "react";
import { Grid, Button } from "@material-ui/core";
import { useHistory } from "react-router-dom";
// import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import GoogleLogin from "react-google-login";
import axios from "axios";
import { makeStyles, useTheme } from "@material-ui/core/styles";
import { useGlobalStore } from "../../contextProviders/globalContext";

const googleLogin = async (accesstoken) => {
  const res = await axios.post(
    `http://127.0.0.1:8080/student/login?token=${accesstoken}`,
    { "Access-Control-Allow-Origin": "http://localhost:3000" }
  );
  console.log("response is", res);
  return res.data.jwt;
};

function GoogleSocialAuth() {
  const history = useHistory();

  const globalStore = useGlobalStore();

  const responseGoogle = async (response) => {
    // console.log(response);
    const token = await googleLogin(response.accessToken, response);
    // Save the token somewhere
    localStorage.setItem("token", token);
    globalStore.modifyToken(token);
    globalStore.logging(true);
    history.push("/profile");
  };
  return (
    <Grid style={{ width: "10%" }}>
      <GoogleLogin
        clientId="486959585513-krclq6b6csu77ds91kcedg0e5e2llnmn.apps.googleusercontent.com"
        render={(renderProps) => (
          <Button
            variant="contained"
            onClick={renderProps.onClick}
            style={{
              backgroundColor: "#148FD3",
              color: "#fff",
              width: "100%",
            }}
          >
            LogIn
          </Button>
        )}
        buttonText="LOGIN"
        onSuccess={responseGoogle}
        onFailure={responseGoogle}
      />
    </Grid>
  );
}

export default GoogleSocialAuth;
