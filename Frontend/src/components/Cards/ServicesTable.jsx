import React from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Typography,
} from "@mui/material";

const ServicesTable = ({ services }) => {
  if (!services || services.length === 0) {
    return <Typography>No services data available</Typography>;
  }

  return (
    <TableContainer component={Paper} sx={{ maxHeight: 600 }}>
      <Table stickyHeader>
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell>Status</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {services.map((service, index) => (
            <TableRow key={index}>
              <TableCell>{service.name}</TableCell>
              <TableCell
                sx={{
                  color: service.status.toLowerCase() === "running" ? "green" : "red",
                  fontWeight: "bold",
                }}
              >
                {service.status}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default ServicesTable;
