import React, { useState } from "react";
import { useNavigate } from "react-router";
import axios from "axios";

function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [validation, setValidation] = useState([]);

  const navigate = useNavigate();

  const registerHandler = async (e) => {
    e.preventDefault();

    await axios
      .post("http://localhost:8080/users/register", {
        username,
        password,
      })
      .then(() => {
        navigate("/");
      })
      .catch((error) => {
        setValidation(error.response);
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
              <form onSubmit={registerHandler}>
                <div className="mb-3">
                  <label className="form-label">Username</label>
                  <input
                    type="text"
                    className="form-control"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                  />
                </div>
                {validation.name && (
                  <div className="alert alert-danger">{validation.name[0]}</div>
                )}
                <div className="mb-3">
                  <label className="form-label">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
                {validation.password && (
                  <div className="alert alert-danger">
                    {validation.password[0]}
                  </div>
                )}
                <div className="d-grid gap-2">
                  <button type="submit" className="btn btn-primary">
                    Create
                  </button>
                </div>
              </form>
              <div className="d-grid gap-5">
                <button className="btn" onClick={redirectToLogin}>
                  Login
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Register;
