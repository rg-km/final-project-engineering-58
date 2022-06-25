import React, { useState } from "react";
import "boxicons";
import 'bootstrap/dist/css/bootstrap.min.css';
import "../assets/css/UbahDataSiswa.css";
import NavSiswa from "./NavSiswa";
import { Link } from "react-router-dom"
import "../assets/css/dashboard.css";

export default function UbahProfil() {
    return(
        <div className="container-all-edit-1">
            <div className="container-all-edit-2">
                <NavSiswa />
                <div className="container-all-edit">
                <h1>Ubah Password</h1>
                    <div className="container-edit">
                        <div className="card-edit">
                            <div className="content-edit">
                                <div className="details-edit">
                                    <h2>Password Lama</h2>
                                </div>
                                <div class="upload">
                                    <input type="text" name="pass-old" id="pass-old" class="text_upload pass-old"/>
                                </div>
                            </div>
                            <div className="content-edit">
                                <div className="details-edit">
                                    <h2>Password Baru</h2>
                                </div>
                                <div class="upload">
                                    <input type="text" name="pass-new" id="pass-new" class="text_upload pass-new"/>
                                </div>
                            </div>
                            <div className="btn-submit text-center">
                                <Link to="/kelas-saya">
                                    <button type="submit" class="mt-2 btn btn-primary btn-lg button-submit">Simpan</button>
                                </Link>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}