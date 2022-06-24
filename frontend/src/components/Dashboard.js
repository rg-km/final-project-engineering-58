import React from "react";
import Container from "react-bootstrap/Container";
import { Link } from "react-router-dom";
import "../assets/css/dashboard.css";
// import class from "../assets/images/icons8-class-50.png";
import user from "../assets/images/icons8-user-64.png";
import kelas from "../assets/images/icons8-class-50.png";
import edit from "../assets/images/icons8-edit-30.png";
import pass from "../assets/images/icons8-show-password-64.png";
import out from "../assets/images/icons8-logout-24.png";

function Dashboard() {
  return (
    // <section className="dashboard bg-white" style={{ height: "100vh" }}>
    //   <Container fluid className="h-100">
    //     <div className="row h-100">
    //       {/* sidebar */}
    //       <div className="col-md-3 bg-dark pt-4 pb-5">
    //         <h2 className="mb-5 text-white">HaloBelajar</h2>
    //         <h6 className="mb-3">
    //           <Link to="/upload-video" className="text-white text-decoration-none">
    //             Upload Video
    //           </Link>
    //         </h6>
    //         <h6 className="mb-5">
    //           <Link to="/user" className="text-white text-decoration-none">
    //             Welcome, Admin
    //           </Link>
    //         </h6>

    //         <Link to="/" className="btn btn-sm btn-danger">
    //           Logout
    //         </Link>
    //       </div>

    //       {/* content dashboard */}
    //       <div className="col-md-9 bg-light pt-4 pb-5 ps-5">
    //         <h1>Dashboard Admin</h1>
    //         <div>
    //           <div className="uploadv">
    //             <span class="input-group-text"></span>
    //             <div>
    //               {/* <div class="ratio ratio-16x9">
    //                 <iframe src="https://www.youtube.com/embed/zpOULjyy-n8?rel=0" title="YouTube video" allowfullscreen></iframe>
    //               </div> */}
    //             </div>
    //           </div>
    //           <div className="container"></div>
    //         </div>
    //       </div>
    //     </div>
    //   </Container>
    // </section>

    <div class="list-group">
      <a href="#" class="list-group-item list-group-item-action list-group-item-dark">
        <div class="image">
          <img alt="" src={user} width="70" height="70" />
        </div>
        <br></br>
        <br></br>
        <p>
          Jane Doe
          <br></br>
          Mahasiswa
          <br></br>
        </p>
      </a>
      <a href="#" class="list-group-item list-group-item-action list-group-item-dark">
        <div class="image" />
        <img alt="" src={kelas} width="20" height="20" /> Kelas Saya
      </a>
      <a href="#" class="list-group-item list-group-item-action list-group-item-dark">
        <div class="image" />
        <img alt="" src={edit} width="20" height="20" />
        Ubah Profil
      </a>
      <a href="#" class="list-group-item list-group-item-action list-group-item-dark">
        <div class="image" />
        <img alt="" src={pass} width="20" height="20" />
        Ubah Password
      </a>
      <a href="#" class="list-group-item list-group-item-action list-group-item-dark">
        <div class="image" />
        <img alt="" src={out} width="20" height="20" />
        Keluar
      </a>
    </div>
  );
}

export default Dashboard;
