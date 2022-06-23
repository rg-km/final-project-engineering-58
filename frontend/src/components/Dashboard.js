import React from "react";
import Container from "react-bootstrap/Container";
import { Link } from "react-router-dom";
import "../assets/css/dashboard.css";

function Dashboard() {
  return (
    <section className="dashboard bg-white" style={{ height: "100vh" }}>
      <Container fluid className="h-100">
        <div className="row h-100">
          {/* sidebar */}
          <div className="col-md-3 bg-dark pt-4 pb-5">
            <h2 className="mb-5 text-white">HaloBelajar</h2>
            <h6 className="mb-3">
              <Link to="/upload-video" className="text-white text-decoration-none">
                Upload Video
              </Link>
            </h6>
            <h6 className="mb-5">
              <Link to="/user" className="text-white text-decoration-none">
                Welcome, Admin
              </Link>
            </h6>

            <Link to="/" className="btn btn-sm btn-danger">
              Logout
            </Link>
          </div>

          {/* content dashboard */}
          <div className="col-md-9 bg-light pt-4 pb-5 ps-5">
            <h1>Dashboard Admin</h1>
            <div>
              <div className="uploadv">
                <span class="input-group-text"></span>
              </div>
              <div class="col align-self-end">
                <input type="file" class="form-control" id="inputGroupFile02" />
              </div>
            </div>
          </div>
        </div>
      </Container>
    </section>
  );
}

export default Dashboard;
