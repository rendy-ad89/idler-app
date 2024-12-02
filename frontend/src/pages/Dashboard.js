import React, { useState, useEffect } from "react";
import Modal from "react-modal";
import { useNavigate } from "react-router";
import axios from "axios";

function Dashboard() {
  const [cash, setCash] = useState(0);
  const [usersMachines, setUsersMachines] = useState([]);
  const [profit, setProfit] = useState(0);
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();

  const accessToken = localStorage.getItem("accessToken");
  const username = localStorage.getItem("username");

  useEffect(() => {
    if (!accessToken) {
      navigate("/");
      return;
    }

    axios.defaults.headers.common["Authorization"] = accessToken;
    axios
      .post("http://localhost:8080/users/calc-profits")
      .then((response) => {
        setCash(response.data);
      })
      .catch(() => {});
  }, [accessToken, navigate]);

  useEffect(() => {
    if (!accessToken) {
      navigate("/");
      return;
    }

    axios.defaults.headers.common["Authorization"] = accessToken;
    axios
      .post("http://localhost:8080/users/machines")
      .then((response) => {
        setUsersMachines(response.data);
      })
      .catch();
  }, [accessToken, navigate]);

  useEffect(() => {
    let newProfit = 0.0;
    let amplification = 100.0;
    const amplifiers = usersMachines.filter((um) => um.Type === "amplifier");
    const generators = usersMachines.filter((um) => um.Type === "generator");
    amplifiers.forEach((m) => {
      amplification += m.Level * m.Generation;
    });
    generators.forEach((m) => {
      newProfit += (m.Level * m.Generation * amplification) / 100;
    });
    console.log("Run");
    setProfit(newProfit);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [JSON.stringify(usersMachines)]);

  useEffect(() => {
    const interval = setInterval(
      () => setCash((oldCash) => oldCash + profit),
      1000
    );
    return () => {
      clearInterval(interval);
    };
  });

  const resolvePrice = (m) => {
    const increment = m.Increment / 100;
    return (m.Price * Math.pow(increment, m.Level)).toFixed(2);
  };

  const resolveDisabled = (m) => {
    return resolvePrice(m) > cash;
  };

  const upgradeHandler = (m) => {
    setCash(cash - resolvePrice(m));
    m.Level += 1;
    setUsersMachines(usersMachines);
  };

  const saveHandler = async () => {
    axios.defaults.headers.common["Authorization"] = accessToken;

    await axios
      .post("http://localhost:8080/users/save-progress", {
        cash: Number(cash),
      })
      .catch(() => {});

    const updatedUsersMachines = [];
    usersMachines.forEach((m) => {
      updatedUsersMachines.push({
        userId: m.UserID,
        machineId: m.MachineID,
        level: m.Level,
      });
    });
    await axios
      .post("http://localhost:8080/users/machines/update", updatedUsersMachines)
      .then(handleOpen)
      .catch(() => {});
  };

  const logoutHandler = () => {
    saveHandler();
    localStorage.removeItem("accessToken");
    localStorage.removeItem("username");
    localStorage.removeItem("cash");
    navigate("/");
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleOpen = () => {
    setOpen(true);
  };

  const modalStyles = {
    content: {
      top: "50%",
      left: "50%",
      right: "auto",
      bottom: "auto",
      marginRight: "-50%",
      transform: "translate(-50%, -50%)",
    },
  };

  return (
    <div className="container" style={{ marginTop: "50px" }}>
      <div className="row justify-content-center">
        <div className="col-md-12 ">
          <div className="card border-0 rounded shadow-sm">
            <div className="card-body">
              <div className="row align-items-center">
                <h1 className="card-body col-md-3 text-center">{username}</h1>
                <h1 className="card-body col-md-3 text-center">
                  Cash: {cash.toFixed(2)}
                </h1>
                <h1 className="card-body col-md-3 text-center">
                  Profit: {profit.toFixed(2)}/s
                </h1>
              </div>
              <hr />
              <div className="row align-items-center">
                {usersMachines.map((m) => (
                  <div
                    className="card-body border rounded col-md-3 m-4"
                    key={m.MachineID}
                  >
                    <h4 className="text-center">{m.Name}</h4>
                    <div className="text-center pb-2">{m.Description}</div>
                    <h5 className="text-center">Price: {resolvePrice(m)}</h5>
                    <h5 className="text-center">Level: {m.Level}</h5>
                    <button
                      onClick={() => upgradeHandler(m)}
                      disabled={resolveDisabled(m)}
                      className="btn btn-primary col-md-10 mx-4"
                    >
                      Upgrade
                    </button>
                  </div>
                ))}
              </div>
              <hr />
              <div class="row justify-content-end">
                <button
                  onClick={saveHandler}
                  className="btn btn-md btn-primary me-2 col-1"
                >
                  Save
                </button>
                <button
                  onClick={logoutHandler}
                  className="btn btn-md btn-danger me-2 col-1"
                >
                  Logout
                </button>
                <Modal style={modalStyles} isOpen={open} onClose={handleClose}>
                  <div className="row justify-content-center">
                    <h3 className="text-center">Progress saved!</h3>
                    <button
                      className="btn btn-primary col-8 my-2"
                      onClick={handleClose}
                    >
                      Close
                    </button>
                  </div>
                </Modal>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
