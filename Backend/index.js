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
app.post("/receive", (req, res) => {
  try {
    console.log("📥 Received Data from Go Client:");
    // console.log(JSON.stringify(req.body, null, 2));

    storedSystemInfo = req.body; 

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
app.get("/system-info", (req, res) => {
  return res.json(storedSystemInfo || {});  // FIXED: prevent 500
});

const PORT = 7000;
app.listen(PORT, () => {
  console.log(`🚀 Node server running on port ${PORT}`);
});
