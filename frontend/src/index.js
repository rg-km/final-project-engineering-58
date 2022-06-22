import React, { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App";
import LoginForm from "./components/login";
import "./index.css";

const rootElement = document.getElementById("root");
const root = createRoot(rootElement);

root.render(
  <StrictMode>
    <App />
  </StrictMode>
);

function App() {
  const adminUser = {
    email: "agistayuana19@gmail.com",
    password: "12345678",
  };

  const [user, setUser] = useState({ name: "", email: "" });
  const [error, setError] = useState("");

  const Login = (details) => {
    console.log(details);

    if (details.email === adminUser.email && details.password === adminUser.password) {
      console.log(details);
      setUser({
        name: details.name,
        email: details.email,
      });
    } else if (details.email == "" && details.name == "") {
      setError("Please fill the bellow!");
    } else {
      console.log("Welcome to HaloBelajar", details.name);
      setUser({
        name: details.name,
        email: details.email,
      });
    }
  };

  const Logout = () => {
    setUser({ name: "", email: "" });
  };
  return (
    <div className="App">
      {user.email !== "" ? (
        <div className="welcome">
          <h2>
            Welcome to HaloBelajar, <span>{user.name}</span>
          </h2>
          <button onClick={Logout}>Logout</button>
        </div>
      ) : (
        <LoginForm LoginForm={Login} error={error} />
      )}
    </div>
  );
}

export default App;
