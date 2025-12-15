import React, { useEffect, useState } from "react";
import HardwareCard from "../Cards/HardwareCard";
import LicenseCard from "../Cards/LicenseCard";
import NetworkCard from "../Cards/NetworkCard";
import OSCard from "../Cards/OSCard";
import SoftwareList from "../Cards/SoftwareList";
import ServicesTable from "../Cards/ServicesTable";
import ProcessesTable from "../Cards/ProcessesTable";
import { fetchSystemInfo } from "../../Api/Api";
import { useParams } from "react-router-dom";
import "./Dashboard.css";

export default function Dashboard() {
  
  const { macAddress } = useParams();
  const [networkInfo, setNetworkInfo] = useState([]);
  const [licenseInfo, setLicenseInfo] = useState(null);
  const [osInfo, setOSInfo] = useState(null);
  const [processes, setProcesses] = useState([]);
  const [services, setServices] = useState([]);
  const [software, setSoftware] = useState([]);
  const [hardwareData, setHardwareData] = useState([]);

  useEffect(() => {

    if (!macAddress) return;
    async function loadData() {
      const data = await fetchSystemInfo(macAddress);
      if (!data) return;

      setHardwareData(data.hardware);
      setOSInfo(data.os);
      setNetworkInfo(data.network);
      setSoftware(data.software);
      setLicenseInfo(data.license);
      setProcesses(data.processes);
      setServices(data.services);
    }

    loadData();
  }, []);

  return (
    <div className="eams-home-page">
      <section>
        <h2>Network Interfaces</h2>
        <NetworkCard networkInfo={networkInfo} />
      </section>

      <section>
        <h2>Installed Software</h2>
        <SoftwareList software={software} />
      </section>

      <section>
        <h2>License Information</h2>
        <LicenseCard license={licenseInfo} />
      </section>

      <section>
        <h2>Services</h2>
        <ServicesTable services={services} />
      </section>

      <section>
        <h2>System Information</h2>
        <OSCard osInfo={osInfo} />
      </section>

      <section>
        <h2>Processes</h2>
        <ProcessesTable processes={processes} />
      </section>

      <section className="home-container">
        <h2>Hardware Information</h2>
        <HardwareCard data={hardwareData} />
      </section>
    </div>
  );
}
