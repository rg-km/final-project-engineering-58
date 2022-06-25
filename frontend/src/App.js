import React from "react";
import { Route, Routes } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import Dashboard from "./components/Dashboard";
import KelasSaya from "./components/KelasSaya";
import WatchVideo from "./pages/WatchVideo";
import Player from "./pages/Player";

const App = () =>{
    return (
        <>
        <Routes>
            <Route index element={<LandingPage />} />
            <Route path="/login" element={<LoginForm />} />
            <Route path="/sign-up" element={<RegisterForm />} />
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/kelas-saya" element={<KelasSaya />} />
            <Route path="/video" element={<Player />} />
            <Route path="/:activeVideo" element={<Player />} />
        </Routes>
        </>
    )
}

export default App