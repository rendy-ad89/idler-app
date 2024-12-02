import React, { useState } from "react";
import { useNavigate } from "react-router";
import axios from "axios";

function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [usernameValidation, setUsernameValidation] = useState("");
  const [passwordValidation, setPasswordValidation] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const navigate = useNavigate();

  const usernameHandler = (username) => {
    setUsername(username);
    setUsernameValidation("");
    setErrorMessage("");
  };

  const passwordHandler = (password) => {
    setPassword(password);
    setPasswordValidation("");
    setErrorMessage("");
  };

  const registerHandler = async (e) => {
    e.preventDefault();

    if (username.length < 5 || username.length > 20) {
      setUsernameValidation("Username must be 5-20 characters.");
    }
    if (password.length < 5 || password.length > 20) {
      setPasswordValidation("Password must be 5-20 characters.");
    }
    if (!username || !password) {
      return;
    }

    await axios
      .post("http://localhost:8080/users/register", {
        username,
        password,
      })
      .then(() => {
        redirectToLogin();
      })
      .catch((error) => {
        setErrorMessage(error.response.data);
      });
  };

  const redirectToLogin = () => {
    navigate("/");
  };

  return (
    <div className="container" style={{ marginTop: "120px" }}>
      <div className="row justify-content-center">
        <div className="col-md-4">
          <div className="card border-0 rounded shadow-sm">
            <div className="card-body">
              <h4 className="fw-bold text-center">Register User</h4>
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
              {usernameValidation && (
                <div className="alert alert-danger">{usernameValidation}</div>
              )}
              <div className="mb-3">
                <label className="form-label">Password</label>
                <input
                  type="password"
                  className="form-control"
                  value={password}
                  onChange={(e) => passwordHandler(e.target.value)}
                />
              </div>
              {passwordValidation && (
                <div className="alert alert-danger">{passwordValidation}</div>
              )}
              <div className="d-grid gap-2">
                <button className="btn btn-primary" onClick={registerHandler}>
                  Create
                </button>
              </div>
              <div className="d-grid gap-5 my-2">
                <button className="btn" onClick={redirectToLogin}>
                  Login
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

export default Register;
