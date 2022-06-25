import React from "react";
import { Route, Routes } from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm";
import Dashboard from "./components/Dashboard";
import KelasSaya from "./components/KelasSaya";
import Player from "./pages/Player";
import UbahProfil from "./components/UbahProfil";
import UbahPass from "./components/UbahPass";

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
            <Route path="/ubah-profil" element={<UbahProfil />} />
            <Route path="/ubah-pass" element={<UbahPass />} />
        </Routes>
        </>
    )
}

export default App