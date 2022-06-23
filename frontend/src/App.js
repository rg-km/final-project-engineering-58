import React from "react";
import { Route, Routes } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import Dashboard from "./components/Dashboard";

const App = () =>{
    return (
        <>
        <Routes>
            <Route index element={<LandingPage />} />
            <Route path="/login" element={<LoginForm />} />
            <Route path="/sign-up" element={<RegisterForm />} />
            <Route path="/dashboard" element={<Dashboard />} />
        </Routes>
        </>
    )
}

export default App