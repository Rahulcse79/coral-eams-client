const express = require("express");
const cors = require("cors");
const helmet = require("helmet");
const morgan = require("morgan");

const app = express();

// Store latest Go system info (IMPORTANT!)
let storedSystemInfo = {};  // FIXED

// Middlewares
app.use(cors());
app.use(helmet());
app.use(express.json({ limit: "10mb" }));
app.use(morgan("dev"));

// Receive data from Go client
app.post("/receive/:macAddress", (req, res) => {
  try {
    const { macAddress } = req.params;
    console.log("📥 Received Data from Go Client:");

    storedSystemInfo[macAddress] = req.body; 

    return res.status(200).json({
      success: true,
      message: "Data received successfully",
    });
  } catch (error) {
    console.error("❌ Error:", error);
    return res.status(500).json({
      success: false,
      message: "Internal server error",
    });
  }
});

// Send system info to frontend
app.get("/system-info/:macAddress", (req, res) => {
  const { macAddress } = req.params;
  return res.json(storedSystemInfo[macAddress] || {});
});

const PORT = 7000;
app.listen(PORT, () => {
  console.log(`🚀 Node server running on port ${PORT}`);
});
