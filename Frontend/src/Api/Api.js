export async function fetchSystemInfo() {
  try {
    const res = await fetch("http://localhost:7000/system-info");

    if (!res.ok) {
      console.error("Failed to fetch system info, status:", res.status);
      return null;
    }

    const data = await res.json();
    console.log("Fetched real system info:", data);

    return {
      hardware: {
        CPUCores: data?.hardware?.cpu_cores || 0,
        CPUModel: data?.hardware?.cpu_model || "",
        DiskTotalGB: data?.hardware?.disk_total_gb || 0,
        Motherboard: data?.hardware?.motherboard || "",
        RAMGB: data?.hardware?.ram_gb || 0,
        SerialNumber: data?.hardware?.serial_number || "",
      },

      os: {
        name: data?.os?.name || "",
        version: data?.os?.version || "",
        kernel: data?.os?.kernel || "",
        architecture: data?.os?.architecture || "",
      },

      network: data?.network || [],
      software: data?.software || [],
      license: data?.license || null,
      processes: data?.processes || [],
      services: data?.services || [],
    };
  } catch (error) {
    console.error("Error fetching system data:", error);
    return null;
  }
}
