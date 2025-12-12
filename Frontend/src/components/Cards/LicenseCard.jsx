import React from "react";
import { Card, CardContent, Typography, Box } from "@mui/material";

const LicenseCard = ({ license }) => {
  if (!license) {
    return (
      <Card sx={{ minWidth: 275, margin: 2 }}>
        <CardContent>
          <Typography variant="h6" color="text.secondary">
            License Info
          </Typography>
          <Typography variant="body2" color="text.secondary">
            No license data available.
          </Typography>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card sx={{ minWidth: 275, margin: 2, boxShadow: 3 }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          License Information
        </Typography>

        <Box sx={{ marginTop: 1 }}>
          <Typography variant="subtitle2" color="text.secondary">
            Product Key:
          </Typography>
          <Typography variant="body1">{license.product_key}</Typography>
        </Box>

        <Box sx={{ marginTop: 2 }}>
          <Typography variant="subtitle2" color="text.secondary">
            OS Identifier:
          </Typography>
          <Typography variant="body1">{license.os_identifier}</Typography>
        </Box>
      </CardContent>
    </Card>
  );
};

export default LicenseCard;
