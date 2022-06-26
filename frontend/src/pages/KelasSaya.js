import React, { useState } from "react";
import "boxicons";
import "../assets/css/KelasSaya.css";
import Avatar from "../assets/images/kelas-1.jpg";
import NavSiswa from "../components/NavSiswa";
import CardKelas from "../components/CardKelas";
import { Link } from "react-router-dom"
import "../assets/css/dashboard.css";

export default function KelasSaya() {

  return (
    <div className="container-all-kelas-1">
        <div className="container-all-kelas-2">
            <NavSiswa />
            <div className="container-all-kelas">
            <h1>Kelas Saya</h1>
                <div className="container-kelas">
                    <CardKelas />
                </div>
            </div>
        </div>
    </div>
  );
}
