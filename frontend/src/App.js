import React from "react";
import { Route, Routes } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import LoginForm from "./components/LoginForm";

const App = () =>{
    return (
        <>
        <Routes>
            <Route index element={<LandingPage />} />
            <Route path="/login" element={<LoginForm />} />
        </Routes>
        </>
    )
}

export default App