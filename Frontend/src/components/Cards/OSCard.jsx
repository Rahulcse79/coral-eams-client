import React from "react";
import { Card, CardContent, Typography, Grid } from "@mui/material";

const OSCard = ({ osInfo }) => {
  if (!osInfo) return null;
  
  return (
    <Card sx={{ minWidth: 275, mb: 2, boxShadow: 3 }}>
      <CardContent>
        <Typography variant="h6" gutterBottom>
          Operating System Info
        </Typography>

        <Grid container spacing={1}>
          {/* OS Name */}
          <Grid item xs={6}>
            <Typography variant="body2" color="textSecondary">
              Name:
            </Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography variant="body1">{osInfo.name}</Typography>
          </Grid>

          {/* OS Version */}
          <Grid item xs={6}>
            <Typography variant="body2" color="textSecondary">
              Version:
            </Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography variant="body1">{osInfo.version}</Typography>
          </Grid>

          {/* Kernel Version */}
          <Grid item xs={6}>
            <Typography variant="body2" color="textSecondary">
              Kernel:
            </Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography variant="body1">{osInfo.kernel}</Typography>
          </Grid>

          {/* Architecture */}
          <Grid item xs={6}>
            <Typography variant="body2" color="textSecondary">
              Architecture:
            </Typography>
          </Grid>
          <Grid item xs={6}>
            <Typography variant="body1">{osInfo.architecture}</Typography>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
};

export default OSCard;
