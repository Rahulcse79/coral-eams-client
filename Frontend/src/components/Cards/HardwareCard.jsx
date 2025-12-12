import React from 'react';
import PropTypes from 'prop-types';
import {
  Card,
  CardHeader,
  CardContent,
  Typography,
  Grid,
  Divider,
  Box
} from '@mui/material';

const get = (data, ...keys) => {
  for (const k of keys) {
    if (data == null) break;
    if (k in data && data[k] != null) return data[k];
  }
  return undefined;
};

const fmtGB = (v) => {
  if (v == null || Number.isNaN(Number(v))) return 'N/A';
  return `${Number(v).toFixed(2)} GB`;
};

export default function HardwareCard({ data }) {
  if (!data) return null;

  const cpuModel = get(data, 'cpu_model', 'CPUModel', 'cpuModel') || 'N/A';
  const cpuCores = get(data, 'cpu_cores', 'CPUCores', 'cpuCores');
  const ramGB = get(data, 'ram_gb', 'RAMGB', 'ramGb', 'ram_gb');
  const diskGB = get(data, 'disk_total_gb', 'DiskTotalGB', 'diskTotalGb');
  const serial = get(data, 'serial_number', 'SerialNumber', 'serialNumber') || 'N/A';
  const motherboard = get(data, 'motherboard', 'Motherboard') || 'N/A';

  return (
    <Card variant="outlined" sx={{ minWidth: 300 }}>
      <CardHeader title="Hardware" subheader="System hardware summary" />
      <Divider />
      <CardContent>
        <Grid container spacing={2}>
          <Grid item xs={12} sm={6}>
            <Typography variant="subtitle2" color="text.secondary">
              CPU Model
            </Typography>
            <Typography variant="body1" sx={{ fontWeight: 600 }}>
              {cpuModel}
            </Typography>
          </Grid>

          <Grid item xs={12} sm={6}>
            <Typography variant="subtitle2" color="text.secondary">
              CPU Cores
            </Typography>
            <Typography variant="body1">{cpuCores != null ? cpuCores : 'N/A'}</Typography>
          </Grid>

          <Grid item xs={12} sm={6}>
            <Typography variant="subtitle2" color="text.secondary">
              RAM
            </Typography>
            <Typography variant="body1">{fmtGB(ramGB)}</Typography>
          </Grid>

          <Grid item xs={12} sm={6}>
            <Typography variant="subtitle2" color="text.secondary">
              Disk Total
            </Typography>
            <Typography variant="body1">{fmtGB(diskGB)}</Typography>
          </Grid>

          <Grid item xs={12} sm={6}>
            <Typography variant="subtitle2" color="text.secondary">
              Serial Number
            </Typography>
            <Typography variant="body1" sx={{ wordBreak: 'break-all' }}>{serial}</Typography>
          </Grid>

          <Grid item xs={12} sm={6}>
            <Typography variant="subtitle2" color="text.secondary">
              Motherboard
            </Typography>
            <Typography variant="body1">{motherboard}</Typography>
          </Grid>
        </Grid>

        <Box mt={2}>
          <Typography variant="caption" color="text.secondary">
            Values come from the agent's hardware collection (gopsutil and platform-specific commands).
          </Typography>
        </Box>
      </CardContent>
    </Card>
  );
}

HardwareCard.propTypes = {
  data: PropTypes.object
};
