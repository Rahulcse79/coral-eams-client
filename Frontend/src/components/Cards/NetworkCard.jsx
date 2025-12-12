import React from "react";
import { Card, CardContent, Typography, Grid, Chip } from "@mui/material";

const NetworkCard = ({ networkInfo }) => {
  return (
    <Grid container spacing={2}>
      {networkInfo.map((iface, index) => (
        <Grid item xs={12} md={6} lg={4} key={index}>
          <Card variant="outlined" sx={{ borderRadius: 2, boxShadow: 1 }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                {iface.interface_name}
              </Typography>

              <Typography variant="body2">
                <strong>MAC:</strong> {iface.mac || "N/A"}
              </Typography>
              <Typography variant="body2">
                <strong>IPv4:</strong> {iface.ipv4 || "N/A"}
              </Typography>
              <Typography variant="body2">
                <strong>IPv6:</strong> {iface.ipv6 || "N/A"}
              </Typography>

              <Chip
                label={iface.is_up ? "Up" : "Down"}
                color={iface.is_up ? "success" : "error"}
                size="small"
                sx={{ mt: 1 }}
              />
            </CardContent>
          </Card>
        </Grid>
      ))}
    </Grid>
  );
};

export default NetworkCard;
