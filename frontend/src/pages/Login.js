import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router";

import axios from "axios";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const navigate = useNavigate();

  useEffect(() => {
    if (localStorage.getItem("accessToken")) {
      navigate("/dashboard");
    }
  }, [navigate]);

  const usernameHandler = (username) => {
    setUsername(username);
    setErrorMessage("");
  };

  const passwordHandler = (password) => {
    setPassword(password);
    setErrorMessage("");
  };

  const loginHandler = async (e) => {
    e.preventDefault();

    if (!username || !password) {
      setErrorMessage("Username and password can't be empty.");
      return;
    }

    await axios
      .post("http://localhost:8080/auth", {
        username,
        password,
      })
      .then((response) => {
        localStorage.setItem("accessToken", response.data.accessToken);
        localStorage.setItem("username", response.data.username);
        localStorage.setItem("cash", response.data.cash);

        navigate("/dashboard");
      })
      .catch((error) => {
        setErrorMessage(error.response.data);
      });
  };

  const redirectToRegister = () => {
    navigate("/register");
  };

  return (
    <div className="container" style={{ marginTop: "120px" }}>
      <div className="row justify-content-center">
        <div className="col-md-4">
          <div className="card border-0 rounded shadow-sm">
            <div className="card-body">
              <h4 className="fw-bold text-center">Login</h4>
              <hr />
              <div className="mb-3">
                <label className="form-label">Username</label>
                <input
                  type="text"
                  className="form-control"
                  value={username}
                  onChange={(e) => usernameHandler(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Password</label>
                <input
                  type="password"
                  className="form-control"
                  value={password}
                  onChange={(e) => passwordHandler(e.target.value)}
                />
              </div>
              <div className="d-grid gap-2">
                <button className="btn btn-primary" onClick={loginHandler}>
                  Submit
                </button>
              </div>
              <div className="d-grid gap-5 my-2">
                <button className="btn" onClick={redirectToRegister}>
                  Create account
                </button>
              </div>
              {errorMessage && (
                <div className="alert alert-danger">{errorMessage}</div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Login;
