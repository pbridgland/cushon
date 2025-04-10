import React, { useState, useEffect } from "react";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [funds, setFunds] = useState([]);
  const [selectedFund, setSelectedFund] = useState("");
  const [amount, setAmount] = useState("");
  const [isLoggedIn, setLoggedIn] = useState(false);

  // Fetch funds and handle auth visibility
  const fetchFunds = async () => {
    try {
      const res = await fetch("/funds");
      if (res.status === 200) {
        const data = await res.json();
        setFunds(data);
        if (data.length > 0) setSelectedFund(data[0].id);
        setLoggedIn(true);
      } else if (res.status === 401) {
        setLoggedIn(false);
      }
    } catch (err) {
      console.error("Error fetching funds", err);
      setLoggedIn(false);
    }
  };

  useEffect(() => {
    fetchFunds();
  }, []);

  const handleLogin = async () => {
    try {
      const res = await fetch("/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
      });

      if (res.status === 200) {
        alert("Login successful!");
        fetchFunds(); // Try to show investment section
      } else {
        alert("Login failed.");
      }
    } catch (err) {
      console.error("Login error:", err);
    }
  };

  const handleInvestmentSubmit = async () => {
    const parsedAmount = parseFloat(amount);
    if (isNaN(parsedAmount) || !/^\d+\.\d{2}$/.test(amount)) {
      alert("Please enter a valid amount with two decimal places.");
      return;
    }

    const investment = {
      fundID: parseInt(selectedFund),
      amount: Math.round(parsedAmount * 100)
    };

    try {
      const res = await fetch("/investments/newinvestment", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(investment)
      });

      if (res.status === 200) {
        alert("Investment submitted!");
      } else if (res.status === 401) {
        alert("Session expired or unauthorized.");
        setLoggedIn(false);
        fetchFunds(); // Try to re-authenticate
      } else {
        alert("Submission failed.");
      }
    } catch (err) {
      console.error("Submission error:", err);
    }
  };

  return (
    <div>
      {!isLoggedIn && (
        <>
          <h2>Login</h2>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          /><br /><br />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          /><br /><br />
          <button onClick={handleLogin}>Login</button>
        </>
      )}
      <hr/>
      {isLoggedIn && (
        <>
          <h2>New Investment</h2>
          <label>Fund:</label><br />
          <select
            value={selectedFund}
            onChange={(e) => setSelectedFund(e.target.value)}
          >
            {funds.map((fund) => (
              <option key={fund.id} value={fund.id}>
                {fund.name}
              </option>
            ))}
          </select><br /><br />

          <label>Amount:</label><br />
          <input
            type="text"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            placeholder="Amount in GBP"
          /><br /><br />
          <button onClick={handleInvestmentSubmit}>Submit Investment</button>
        </>
      )}
    </div>
  );
}

export default App;
